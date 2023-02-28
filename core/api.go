package gencode

import (
	"bytes"
	"fmt"
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
		log.Println("goZeroOutput:" + out.String())
		if err != nil {
			return err
		}
	}
	log.Println("api success!")

	return nil
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
