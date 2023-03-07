package gencode

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

const (
	modelModel  = "/model_model.tpl"
	modelSchema = "/model_schema.tpl"
	modelTyps   = "/model_typs.tpl"
)

type ModelSchema struct {
	*Dataset
	*ModelConfig
}

type ModelConfig struct {
	*DBConfig
	Tables      []string
	Switch      string
	GoZeroStyle string
	IsCache     bool
}

func NewModelSchema(dataset *Dataset, config *ModelConfig) *ModelSchema {
	model := &ModelSchema{
		Dataset:     dataset,
		ModelConfig: config,
	}
	if model.OutPath == "" {
		model.OutPath = "model"
	}
	if model.GoZeroStyle == "" {
		model.GoZeroStyle = "goZero"
	}
	model.filterTable()

	return model
}

func (m *ModelSchema) Generate() (err error) {
	if m.Switch != switch_file && m.Switch != switch_file_cmd {
		return nil
	}

	schemaPath := m.TemplateFilePath + modelSchema
	typesPath := m.TemplateFilePath + modelTyps
	modelPath := m.TemplateFilePath + modelModel
	t := WithTemplate(schemaPath, typesPath, modelPath)

	err = m.genSchema(t)
	if err != nil {
		return err
	}

	err = m.genTypes(t)
	if err != nil {
		return err
	}

	err = m.genModel(t)
	if err != nil {
		return err
	}

	log.Println("model success!")

	return nil
}

func (m *ModelSchema) filterTable() {
	if len(m.Tables) == 0 {
		return
	}
	ts := make([]*Table, 0)
	for _, table := range m.Dataset.TableSet {
		if inSlice(m.Tables, table.Name) {
			ts = append(ts, table)
		}
	}
	m.Dataset.TableSet = ts
}

func (m *ModelSchema) genSchema(template *template.Template) error {
	buf := new(bytes.Buffer)

	err := template.ExecuteTemplate(buf, strings.TrimLeft(modelSchema, "/"), *m)
	if err != nil {
		return err
	}

	return CreateAndWriteFile(m.OutPath, "schema.go", buf.String())
}

func (m *ModelSchema) genTypes(template *template.Template) error {
	buf := new(bytes.Buffer)

	err := template.ExecuteTemplate(buf, strings.TrimLeft(modelTyps, "/"), *m)
	if err != nil {
		return err
	}

	return CreateAndWriteFile(m.OutPath, "typs.go", buf.String())
}

func (m *ModelSchema) genModel(template *template.Template) error {
	buf := new(bytes.Buffer)

	type tp struct {
		IsCache bool
		Primary *Field
		*Table
	}

	for _, t := range m.Dataset.TableSet {
		tp := &tp{
			IsCache: m.IsCache,
			Table:   t,
		}
		for _, f := range t.Fields {
			if f.IsPrimary {
				tp.Primary = f
			}
			if isDateType(f.OriginDataType) {
				f.DataType = "*time.Time"
			}
		}
		if tp.Primary == nil {
			continue
		}

		err := template.ExecuteTemplate(buf, strings.TrimLeft(modelModel, "/"), *tp)
		if err != nil {
			return err
		}

		fileName := t.StyleName + ".go"
		err = CreateAndWriteFile(m.OutPath, fileName, buf.String())
		if err != nil {
			return err
		}
		buf.Reset()
	}

	return nil
}
