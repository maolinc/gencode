package gencode

import (
	"bytes"
	"fmt"
	"github.com/maolinc/gencode/tools/filex"
	"log"
	"os/exec"
	"strings"
	"text/template"
)

var _ Generate = (*ProtoSchema)(nil)

const (
	// proto3 is a describing the proto3 syntax type.
	proto3    = "proto3"
	goPackage = "./pb"
	protoT    = "/proto.tpl"

	switch_no       = "N"
	switch_file     = "A"
	switch_file_cmd = "B"
)

type ProtoSchema struct {
	*Dataset
	*ProtoConfig
}

type ProtoConfig struct {
	Switch      string
	Syntax      string
	GoPackage   string
	Package     string
	GoZeroStyle string
	ModelPath   string
}

type ProtoOption func(schema *ProtoSchema)

func WithProtoPath(path string) ProtoOption {
	return func(schema *ProtoSchema) {
		schema.OutPath = path
	}
}

func WithProtoGoZeroStyle(goZeroStyle string) ProtoOption {
	return func(schema *ProtoSchema) {
		schema.GoZeroStyle = goZeroStyle
	}
}

func WithProtoSyntax(syntax string) ProtoOption {
	return func(schema *ProtoSchema) {
		schema.Syntax = syntax
	}
}
func WithProtoGoPackage(goPkg string) ProtoOption {
	return func(schema *ProtoSchema) {
		schema.GoPackage = goPkg
	}
}

func WithProtoPackage(pkg string) ProtoOption {
	return func(schema *ProtoSchema) {
		schema.Package = pkg
	}
}

func (s *ProtoSchema) Generate() error {
	if s.Switch != switch_file && s.Switch != switch_file_cmd {
		return nil
	}

	filePath := s.TemplateFilePath + protoT
	buf := new(bytes.Buffer)

	err := PareTemplate(protoT, filePath, *s, buf)
	if err != nil {
		return err
	}

	err = CreateAndWriteFile(s.OutPath, s.ServiceName+".proto", buf.String())
	if err != nil {
		return err
	}

	if s.Switch == switch_file_cmd {
		protoPath := fmt.Sprintf("%s/%s.proto", s.OutPath, s.ServiceName)
		rpcCmd := fmt.Sprintf("rpc protoc %s --go_out=%s/pb --go-grpc_out=%s/pb --zrpc_out=%s -m=true --style=%s", protoPath, s.OutPath, s.OutPath, s.OutPath, s.GoZeroStyle)
		args := strings.Split(rpcCmd, " ")
		cmd := exec.Command("goctl", args...)
		out := bytes.Buffer{}
		cmd.Stdout = &out
		err = cmd.Run()
		log.Println("goZeroOutput:" + out.String())
		if err != nil {
			return err
		}

		err = s.GenerateCrud()
	}
	log.Println("proto success!")
	exec.Command("gofmt", "-s", "-w", s.OutPath).Run()

	return nil
}

func NewProtoSchema(dataset *Dataset, config *ProtoConfig, opts ...ProtoOption) *ProtoSchema {
	s := &ProtoSchema{
		Dataset:     dataset,
		ProtoConfig: config,
	}
	if config == nil {
		s.ProtoConfig = &ProtoConfig{}
	}
	if s.ProtoConfig.Syntax == "" {
		s.ProtoConfig.Syntax = proto3
	}
	if s.ProtoConfig.GoPackage == "" {
		s.ProtoConfig.GoPackage = goPackage
	}
	if s.ProtoConfig.Package == "" {
		s.ProtoConfig.Package = dataset.ServiceName
	}
	if s.ProtoConfig.GoZeroStyle == "" {
		s.ProtoConfig.GoZeroStyle = "goZero"
	}
	if s.OutPath == "" {
		s.OutPath = "rpc"
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

const (
	createTplP = "/proto_create.tpl"
	updateTplP = "/proto_update.tpl"
	detailTplP = "/proto_detail.tpl"
	deleteTplP = "/proto_delete.tpl"
	pageTplP   = "/proto_page.tpl"

	svcTplP = "/api_svc.tpl"
)

func (s *ProtoSchema) GenerateCrud() error {
	module, path := filex.GetModule(s.ModelPath)
	modelPkg := module + "/" + strings.TrimPrefix(path, "/")
	modelPkg = "\"" + modelPkg + "\""

	createPath := s.TemplateFilePath + createTplP
	updatePath := s.TemplateFilePath + updateTplP
	detailPath := s.TemplateFilePath + detailTplP
	deleteTPath := s.TemplateFilePath + deleteTplP
	pagePath := s.TemplateFilePath + pageTplP

	svcPath := s.TemplateFilePath + svcTplP

	template := WithTemplate(createPath, updatePath, detailPath, deleteTPath, pagePath, svcPath)

	err := s.genSvc(template, modelPkg)
	if err != nil {
		return err
	}

	pfv := func(t Table) string {
		var v string
		for _, field := range t.Fields {
			if field.IsPrimary {
				v = v + " , in." + field.CamelName
			}
		}
		return strings.TrimLeft(v, " ,")
	}

	for _, table := range s.TableSet {
		t := Tp{
			IsCache:     s.IsCache,
			Package:     strings.ToLower(table.CamelName) + "logic",
			ModelPkg:    modelPkg,
			PrimaryFmtV: pfv(*table),
			Table:       table,
			SourcePath:  s.OutPath + "/internal/logic/" + strings.ToLower(table.CamelName),
		}
		// not primary don‘t gen
		if t.PrimaryFmtV == "" {
			continue
		}
		err := doGenerateCrudProto(template, &t, s.GoZeroStyle)
		if err != nil {
			return err
		}
	}

	exec.Command("gofmt", "-s", "-w", s.OutPath).Run()

	return nil
}

func doGenerateCrudProto(template *template.Template, tp *Tp, style string) error {
	file := tp.CamelName
	mp := map[string]string{"Create" + file: createTplP, "Update" + file: updateTplP, "Delete" + file: deleteTplP,
		"Detail" + file: detailTplP, "Page" + file: pageTplP}

	for f, t := range mp {
		err := crud(template, tp, getRealNameByStyle(f+"Logic.go", style), t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ProtoSchema) genSvc(template *template.Template, modelPkg string) error {
	t := Tp{
		IsCache:    s.IsCache,
		Package:    "svc",
		ModelPkg:   modelPkg,
		SourcePath: s.OutPath + "/internal/svc",
		Dataset:    s.Dataset,
	}
	buf := new(bytes.Buffer)
	err := template.ExecuteTemplate(buf, strings.TrimLeft(svcTplP, "/"), t)
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
