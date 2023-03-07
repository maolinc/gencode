package main

import (
	"embed"
	gencode "github.com/maolinc/gencode/core"
	"log"
)

//go:embed template
var dir embed.FS

func main() {
	var err error

	err = checkTemplate()
	if err != nil {
		log.Fatal(err)
	}

	jsonConfig := getDefaultConfig()
	err = parseFlag(jsonConfig)
	if err != nil {
		log.Fatal(err)
	}

	dataset := gencode.From(jsonConfig.DBConfig, jsonConfig.GlobalConfig)

	apiSchema := gencode.NewApiSchema(dataset.Session(jsonConfig.ApiConfig.SessionConfig), jsonConfig.ApiConfig.ApiConfig)

	protoSchema := gencode.NewProtoSchema(dataset.Session(jsonConfig.ProtoConfig.SessionConfig), jsonConfig.ProtoConfig.ProtoConfig)

	//jsonConfig.ModelConfig.DBConfig = jsonConfig.DBConfig
	modelSchema := gencode.NewModelSchema(dataset.Session(jsonConfig.ModelConfig.SessionConfig), jsonConfig.ModelConfig.ModelConfig)

	gencode.Generates(apiSchema, protoSchema, modelSchema)
}
