package gencode

import (
	"bytes"
	"fmt"
	"github.com/maolinc/gencode/tools/astx"
	"github.com/maolinc/gencode/tools/filex"
	"github.com/maolinc/gencode/tools/stringx"
	"log"
	"os/exec"
	"strings"
	"text/template"
)

const (
	apiApi    = "/api_api.tpl"
	apiCommon = "/api_common.tpl"
	apiTypes  = "/api_types.tpl"

	syntaxx = "v1"
)

var _ Generate = (*ApiSchema)(nil)

type ApiSchema struct {
	*Dataset
	*ApiConfig
}

type ApiConfig struct {
	Switch      string
	Syntax      string
	Prefix      string
	GoZeroStyle string
	Author      string
	Email       string
	Version     string
}

type ApiOption func(schema *ApiSchema)

func WithApiSyntax(syntax string) ApiOption {
	return func(schema *ApiSchema) {
		schema.Syntax = syntax
	}
}

func WithGoZeroStyleSyntax(goZeroStyle string) ApiOption {
	return func(schema *ApiSchema) {
		schema.GoZeroStyle = goZeroStyle
	}
}

func WithApiPrefix(prefix string) ApiOption {
	return func(schema *ApiSchema) {
		schema.Prefix = prefix
	}
}

func NewApiSchema(dataset *Dataset, config *ApiConfig, opts ...ApiOption) *ApiSchema {
	s := &ApiSchema{
		Dataset:   dataset,
		ApiConfig: config,
	}
	if config == nil {
		s.ApiConfig = &ApiConfig{}
	}
	if s.ApiConfig.Syntax == "" {
		s.ApiConfig.Syntax = syntaxx
	}
	if s.ApiConfig.Prefix == "" {
		s.ApiConfig.Prefix = dataset.ServiceName
	}
	if s.ApiConfig.GoZeroStyle == "" {
		s.ApiConfig.GoZeroStyle = "goZero"
	}
	if s.OutPath == "" {
		s.OutPath = "api"
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *ApiSchema) Generate() error {
	if s.Switch != switch_file && s.Switch != switch_file_cmd {
		return nil
	}

	var err error
	commonPath := s.TemplateFilePath + apiCommon
	apiPath := s.TemplateFilePath + apiApi
	typesPath := s.TemplateFilePath + apiTypes
	t := WithTemplate(commonPath, apiPath, typesPath)

	err = s.genCommon(t)
	if err != nil {
		return err
	}

	err = s.genApi(t)
	if err != nil {
		return err
	}

	err = s.genTypes(t)
	if err != nil {
		return err
	}

	if s.Switch == switch_file_cmd {
		apiPath := fmt.Sprintf("%s/desc/%s.api", s.OutPath, s.ServiceName)
		rpcCmd := fmt.Sprintf("api go -api %s -dir %s -style %s", apiPath, s.OutPath, s.GoZeroStyle)
		args := strings.Split(rpcCmd, " ")
		cmd := exec.Command("goctl", args...)
		out := bytes.Buffer{}
		cmd.Stdout = &out
		err = cmd.Run()
		//log.Println("goZeroOutput:" + out.String())
		if err != nil {
			return err
		}
	}
	log.Println("api success!")
	exec.Command("gofmt", "-s", "-w", s.OutPath).Run()

	return nil
}

type Tp struct {
	IsCache     bool
	Package     string
	ModelPkg    string
	SourcePath  string
	PrimaryFmtV string
	*Table
	Dataset *Dataset
}

const (
	createTpl = "/api_create.tpl"
	updateTpl = "/api_update.tpl"
	detailTpl = "/api_detail.tpl"
	deleteTpl = "/api_delete.tpl"
	pageTpl   = "/api_page.tpl"

	svcTpl = "/api_svc.tpl"
)

func (s *ApiSchema) GenerateCrud(modelPath string) error {
	if s.Switch != switch_file_cmd {
		return nil
	}
	module, path := filex.GetModule(modelPath)
	modelPkg := module + "/" + strings.TrimPrefix(path, "/")
	modelPkg = "\"" + modelPkg + "\""

	createPath := s.TemplateFilePath + createTpl
	updatePath := s.TemplateFilePath + updateTpl
	detailPath := s.TemplateFilePath + detailTpl
	deleteTPath := s.TemplateFilePath + deleteTpl
	pagePath := s.TemplateFilePath + pageTpl

	svcPath := s.TemplateFilePath + svcTpl

	template := WithTemplate(createPath, updatePath, detailPath, deleteTPath, pagePath, svcPath)

	err := s.genSvc(template, modelPkg)
	if err != nil {
		return err
	}

	pfv := func(t Table) string {
		var v string
		for _, field := range t.Fields {
			if field.IsPrimary {
				v = v + " , req." + field.CamelName
			}
		}
		return strings.TrimLeft(v, " ,")
	}

	for _, table := range s.TableSet {
		t := Tp{
			IsCache:     s.IsCache,
			Package:     strings.ToLower(table.CamelName),
			ModelPkg:    modelPkg,
			PrimaryFmtV: pfv(*table),
			Table:       table,
			SourcePath:  s.OutPath + "/internal/logic/" + strings.ToLower(table.CamelName),
		}
		// not primary don‘t gen
		if t.PrimaryFmtV == "" {
			continue
		}
		err := doGenerateCrud(template, &t, s.GoZeroStyle)
		if err != nil {
			return err
		}
	}

	exec.Command("gofmt", "-s", "-w", s.OutPath).Run()

	return nil
}

func (s *ApiSchema) genSvc(template *template.Template, modelPkg string) error {
	t := Tp{
		IsCache:    s.IsCache,
		Package:    "svc",
		ModelPkg:   modelPkg,
		SourcePath: s.OutPath + "/internal/svc",
		Dataset:    s.Dataset,
	}
	buf := new(bytes.Buffer)
	err := template.ExecuteTemplate(buf, strings.TrimLeft(svcTpl, "/"), t)
	if err != nil {
		return err
	}

	// Append to the original file
	path := t.SourcePath + "/" + getRealNameByStyle("ServiceContext.go", s.GoZeroStyle)
	err = filex.AppendToFile(path, buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func doGenerateCrud(template *template.Template, tp *Tp, style string) error {
	file := tp.CamelName
	mp := map[string]string{"Create" + file: createTpl, "Update" + file: updateTpl, "Delete" + file: deleteTpl,
		file + "Detail": detailTpl, file + "Page": pageTpl}

	for f, t := range mp {
		err := crud(template, tp, getRealNameByStyle(f+"Logic.go", style), t)
		if err != nil {
			return err
		}
	}
	return nil
}

func crud(template *template.Template, tp *Tp, file string, tpl string) error {
	buf := new(bytes.Buffer)
	err := template.ExecuteTemplate(buf, strings.TrimLeft(tpl, "/"), *tp)
	if err != nil {
		return err
	}

	fullFile := tp.SourcePath + "/" + file
	err = astx.MergeSource(fullFile, buf.String(), tp.Package)
	if err != nil {
		return err
	}

	return nil
}

func getRealNameByStyle(name, style string) string {
	switch style {
	case "gozero":
		name = stringx.From(name).Lower()
	case "go_zero":
		name = stringx.From(name).ToSnake()
	default:
		name = stringx.From(name).ToCamelWithStartLower()
	}
	return name
}

func (s *ApiSchema) genCommon(template *template.Template) error {
	buf := new(bytes.Buffer)

	err := template.ExecuteTemplate(buf, strings.TrimLeft(apiCommon, "/"), *s)
	if err != nil {
		return err
	}

	fileP := fmt.Sprintf("%s/desc/%s", s.OutPath, "types")
	return CreateAndWriteFile(fileP, "common.api", buf.String())
}

func (s *ApiSchema) genApi(template *template.Template) error {
	buf := new(bytes.Buffer)

	err := template.ExecuteTemplate(buf, strings.TrimLeft(apiApi, "/"), *s)
	if err != nil {
		return err
	}

	return CreateAndWriteFile(s.OutPath+"/desc", s.ServiceName+".api", buf.String())
}

func (s *ApiSchema) genTypes(template *template.Template) error {
	buf := new(bytes.Buffer)

	type tp struct {
		ServiceName string
		Syntax      string
		Prefix      string
		Author      string
		Email       string
		Version     string
		*Table
	}

	for _, m := range s.Dataset.TableSet {
		tp := &tp{
			ServiceName: s.ServiceName,
			Syntax:      syntaxx,
			Prefix:      s.Prefix,
			Author:      s.Author,
			Email:       s.Email,
			Version:     s.Version,
			Table:       m,
		}

		err := template.ExecuteTemplate(buf, strings.TrimLeft(apiTypes, "/"), *tp)
		if err != nil {
			return err
		}

		fileP := fmt.Sprintf("%s/desc/%s", s.OutPath, "types")
		typeFileName := strings.ToLower(m.CamelName) + ".api"
		err = CreateAndWriteFile(fileP, typeFileName, buf.String())
		if err != nil {
			return err
		}
		buf.Reset()
	}

	return nil
}
