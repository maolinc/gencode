package gencode

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	}
	log.Println("proto success!")

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
