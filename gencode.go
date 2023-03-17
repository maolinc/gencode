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

	jsonConfig := getDefaultConfig()
	err = parseFlag(jsonConfig)
	if err != nil {
		log.Fatal(err)
	}

	dataset := gencode.From(jsonConfig.DBConfig, jsonConfig.GlobalConfig)

	//jsonConfig.ModelConfig.DBConfig = jsonConfig.DBConfig
	modelSchema := gencode.NewModelSchema(dataset.Session(jsonConfig.ModelConfig.SessionConfig), jsonConfig.ModelConfig.ModelConfig)

	jsonConfig.ApiConfig.ModelPath = jsonConfig.ModelConfig.OutPath
	jsonConfig.ApiConfig.IsCache = modelSchema.IsCache
	apiSchema := gencode.NewApiSchema(dataset.Session(jsonConfig.ApiConfig.SessionConfig), jsonConfig.ApiConfig.ApiConfig)

	jsonConfig.ProtoConfig.ModelPath = jsonConfig.ModelConfig.OutPath
	jsonConfig.ProtoConfig.IsCache = modelSchema.IsCache
	protoSchema := gencode.NewProtoSchema(dataset.Session(jsonConfig.ProtoConfig.SessionConfig), jsonConfig.ProtoConfig.ProtoConfig)

	gencode.Generates(apiSchema, protoSchema, modelSchema)

}
