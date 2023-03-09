package gencode

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maolinc/gencode/tools"
	"github.com/maolinc/gencode/tools/stringx"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"
)

const (
	TemplatePath = "gencode"
)

var (
	DefaultIgnoreFileValue = map[string]int64{"create_time": 3, "create_at": 3, "update_time": 3, "update_by": 3, "delete_flag": 7, "del_flag": 7, "create_by": 3}
)

type Generate interface {
	Generate() error
}

func Generates(gens ...Generate) {
	wait := sync.WaitGroup{}
	wait.Add(len(gens))
	for _, gen := range gens {
		go func(gen Generate) {
			err := gen.Generate()
			if err != nil {
				log.Println(err)
			}
			wait.Done()
		}(gen)
	}
	wait.Wait()
}

type Table struct {
	Name      string
	CamelName string
	StyleName string
	Comment   string
	Fields    []*Field
}

type Field struct {
	Name           string
	CamelName      string
	StyleName      string
	OriginDataType string
	DataType       string
	HasDefault     bool
	Default        string
	Comment        string
	Key            string
	IsPrimary      bool
	IsNullable     bool
	MaxLength      int64
	Sort           int64
	// used to identify the ignored value of the field, and the binary sum, binary system:1111111, decimal system 127
	// Create:0000001-1 Delete:0000010-2 Select(View):0000100-4 Update:0001000-8
	//  1 | 110 111    11  0
	IgnoreValue int64
}

type styleType string

const (
	Style_Camel styleType = "mLc"
	Style_Snake styleType = "m_lc"
)

type Dataset struct {
	*SessionConfig
	TableSet []*Table
}

type SessionConfig struct {
	// Specify the path to generate the code file; default .
	OutPath          string
	TemplateFilePath string
	ServiceName      string
	// value and ignore mapping rule: 1(create), 2(update),4(select),8(delete), 1+2=3(create,update),1+2+4=7(create,select,update)
	// eg: {create_time:1, delete_flag:7, id:1}
	IgnoreFieldValue map[string]int64
	// two style: string|number, default number
	DateStyle string
}

// Global configuration of dataset
type Config struct {
	FieldStyle  styleType
	ServiceName string
	// specify select table
	Tables []string
	// ignore select table
	IgnoreTables []string
	// global ignore select field
	IgnoreFields []string
	// ignore the specified fields of the specified table
	IgnoreMap map[string][]string
}

type DBConfig struct {
	DbType   string
	DBName   string
	Host     string
	User     string
	Password string
	Port     int
}

// Generate a new dataset without affecting the original dataset
func (d *Dataset) Session(config *SessionConfig) *Dataset {
	var (
		conf       = *d.SessionConfig
		newDataset = &Dataset{
			SessionConfig: &conf,
		}
	)

	if config == nil {
		config = &SessionConfig{}
	}

	if config.IgnoreFieldValue == nil {
		config.IgnoreFieldValue = DefaultIgnoreFileValue
	}
	if config.ServiceName != "" {
		newDataset.SessionConfig.ServiceName = config.ServiceName
	}
	if config.OutPath != "" {
		newDataset.SessionConfig.OutPath = config.OutPath
	}
	if config.TemplateFilePath != "" {
		newDataset.SessionConfig.TemplateFilePath = config.TemplateFilePath
	} else {
		newDataset.SessionConfig.TemplateFilePath = tools.GetHomeDir() + "/" + TemplatePath
	}

	// copy tableSet and filter field
	tableSet := make([]*Table, 0)
	for _, table := range d.TableSet {
		fields := make([]*Field, 0)
		for _, field := range table.Fields {
			nField := *field
			nField.IgnoreValue = config.IgnoreFieldValue[nField.Name]
			if config.DateStyle == "string" && isDateType(field.OriginDataType) {
				nField.DataType = "string"
			}
			fields = append(fields, &nField)
		}
		nTable := *table
		nTable.Fields = fields
		tableSet = append(tableSet, &nTable)
	}
	newDataset.TableSet = tableSet

	return newDataset
}

func From(dbConfig *DBConfig, config *Config) (dataset *Dataset) {

	if config == nil {
		config = &Config{}
	}

	if config.ServiceName == "" {
		config.ServiceName = dbConfig.DBName
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := sql.Open(dbConfig.DbType, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbName, err := existDB(db)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := dbColumn(db, dbName, config)
	if err != nil {
		log.Fatal(err)
	}

	tables, err := transformTable(columns, config)
	if err != nil {
		log.Fatal(err)
	}

	dataset = &Dataset{
		SessionConfig: &SessionConfig{
			ServiceName:      config.ServiceName,
			TemplateFilePath: TemplatePath,
			OutPath:          "",
		},
		TableSet: tables,
	}

	return
}

func transformTable(columns []*column, config *Config) ([]*Table, error) {
	ignoreMap := config.IgnoreMap
	if ignoreMap == nil {
		ignoreMap = map[string][]string{}
	}

	tableMap := make(map[string][]*Field)
	tables := make([]*Table, 0)

	for _, c := range columns {
		if fields, ok := ignoreMap[c.TableName]; ok && inSlice(fields, c.ColumnName) {
			continue
		}
		fields, ok := tableMap[c.TableName]
		if !ok {
			fields = make([]*Field, 0)
			tables = append(tables, &Table{
				Name:      c.TableName,
				CamelName: stringx.From(c.TableName).ToCamel(),
				StyleName: getStyleName(config.FieldStyle, c.TableName),
				Comment:   c.TableComment,
				Fields:    nil,
			})
		}
		field := &Field{
			Name:           c.ColumnName,
			CamelName:      stringx.From(c.ColumnName).ToCamel(),
			StyleName:      getStyleName(config.FieldStyle, c.ColumnName),
			Default:        c.ColumnDefault.String,
			HasDefault:     c.ColumnDefault.String != "",
			Comment:        c.ColumnComment,
			OriginDataType: c.DataType,
			DataType:       transformDataType(c.DataType),
			Key:            c.ColumnKey.String,
			IsPrimary:      isPrimaryKey(c.ColumnKey.String),
			IsNullable:     isNullable(c.IsNullable),
			MaxLength:      c.CharacterMaximumLength.Int64,
			Sort:           int64(len(fields) + 1),
		}

		fields = append(fields, field)
		tableMap[c.TableName] = fields
	}

	for i := range tables {
		fields := tableMap[tables[i].Name]
		tables[i].Fields = fields
	}

	return tables, nil
}

type column struct {
	Style                  string
	TableName              string
	TableComment           string
	ColumnName             string
	IsNullable             string
	DataType               string
	ColumnKey              sql.NullString
	ColumnDefault          sql.NullString
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	ColumnComment          string
}

func dbColumn(db *sql.DB, dbName string, config *Config) ([]*column, error) {

	tables := config.Tables
	ignoreTables := config.IgnoreTables
	ignoreFields := config.IgnoreFields

	q := "SELECT c.TABLE_NAME, c.COLUMN_NAME, c.IS_NULLABLE, c.DATA_TYPE,c.COLUMN_KEY,c.COLUMN_DEFAULT,  " +
		"c.CHARACTER_MAXIMUM_LENGTH, c.NUMERIC_PRECISION, c.NUMERIC_SCALE, c.COLUMN_TYPE ,c.COLUMN_COMMENT,t.TABLE_COMMENT " +
		"FROM INFORMATION_SCHEMA.COLUMNS as c  LEFT JOIN  INFORMATION_SCHEMA.TABLES as t  on c.TABLE_NAME = t.TABLE_NAME and  c.TABLE_SCHEMA = t.TABLE_SCHEMA" +
		" WHERE c.TABLE_SCHEMA = ?"

	if tables != nil && len(tables) > 0 {
		q += fmt.Sprintf(" AND c.TABLE_NAME IN ('%s')", strings.Join(tables, "','"))
	}
	if ignoreTables != nil && len(ignoreTables) > 0 {
		q += fmt.Sprintf(" AND c.TABLE_NAME NOT IN ('%s')", strings.Join(ignoreTables, "','"))
	}
	if ignoreFields != nil && len(ignoreFields) > 0 {
		q += fmt.Sprintf(" AND c.COLUMN_NAME NOT IN ('%s')", strings.Join(ignoreFields, "','"))
	}

	q += " ORDER BY c.TABLE_NAME, c.ORDINAL_POSITION"
	rows, err := db.Query(q, dbName)
	defer rows.Close()
	if nil != err {
		return nil, err
	}

	cols := make([]*column, 0)

	for rows.Next() {
		col := &column{}
		err := rows.Scan(&col.TableName, &col.ColumnName, &col.IsNullable, &col.DataType, &col.ColumnKey, &col.ColumnDefault,
			&col.CharacterMaximumLength, &col.NumericPrecision, &col.NumericScale, &col.ColumnType, &col.ColumnComment, &col.TableComment)
		if err != nil {
			log.Fatal(err)
		}

		if col.TableComment == "" {
			col.TableComment = stringx.From(col.TableName).ToCamelWithStartLower()
		}

		cols = append(cols, col)
	}
	if err := rows.Err(); nil != err {
		return nil, err
	}

	return cols, nil
}

func existDB(db *sql.DB) (string, error) {
	var schema string
	err := db.QueryRow("SELECT SCHEMA()").Scan(&schema)
	return schema, err
}

func isPrimaryKey(key string) bool {
	return key == "PRI"
}

func isNullable(key string) bool {
	return key == "YES"
}

func transformDataType(typ string) string {
	dataType := "string"
	switch typ {
	case "char", "varchar", "text", "longtext", "mediumtext", "tinytext":
		dataType = "string"
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		dataType = "bytes"
	case "date", "time", "datetime", "timestamp":
		//s.AppendImport("google/protobuf/timestamp.proto")
		dataType = "int64"
	case "bool", "bit":
		dataType = "bool"
	case "tinyint", "smallint", "int", "mediumint", "bigint":
		dataType = "int64"
	case "float", "decimal", "double":
		dataType = "float64"
	case "json":
		dataType = "string"
	}
	return dataType
}

func isDateType(typ string) bool {
	return inSlice([]string{"date", "time", "datetime", "timestamp"}, typ)
}

func inSlice(slice []string, dest string) bool {
	for i := range slice {
		if slice[i] == dest {
			return true
		}
	}
	return false
}

func PareTemplate(fileName string, filePath string, data any, buffer *bytes.Buffer) error {
	funcMap := template.FuncMap{"add": add, "toCamelWithStartLower": toCamelWithStartLower, "toLower": toLower, "isIgnore": isIgnore}
	parseFiles, err := template.New(strings.TrimLeft(fileName, "/")).Funcs(funcMap).ParseFiles(filePath)
	if err != nil {
		return err
	}
	err = parseFiles.Execute(buffer, data)
	return err
}

func WithTemplate(filePaths ...string) *template.Template {
	funcMap := template.FuncMap{"add": add, "toCamelWithStartLower": toCamelWithStartLower, "toLower": toLower, "isIgnore": isIgnore}
	mt := template.Must(
		template.New("").Funcs(funcMap).ParseFiles(filePaths...),
	)
	return mt
}

func CreateAndWriteFile(fileP, fileN, context string) error {
	err := os.MkdirAll(fileP, 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(fileP + "/" + fileN)
	_, err = file.WriteString(context)
	if err != nil {
		return err
	}
	return nil
}

func getStyleName(style styleType, name string) string {
	if Style_Snake == style {
		return stringx.From(name).ToSnake()
	}
	return stringx.From(name).ToCamelWithStartLower()
}

func add(args ...int) int64 {
	var sum int64
	for _, i := range args {
		sum = sum + int64(i)
	}
	return sum
}

func toCamelWithStartLower(s string) string {
	return stringx.From(s).ToCamelWithStartLower()
}

func toLower(s string) string {
	return strings.ToLower(s)
}

// false->ignore  true->show  eg:2 & 7 = 010 & 111 = 010= 2==0=false ignore
func isIgnore(checkValue, ignoreValue int64) bool {
	a := checkValue & ignoreValue
	return a == 0
}
