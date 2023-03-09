package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/maolinc/gencode/core"
	"github.com/maolinc/gencode/tools/filex"
	"os"
	"strings"
)

type JsonConfig struct {
	DBConfig     *gencode.DBConfig
	GlobalConfig *gencode.Config
	ApiConfig    *gencode.ApiSchema
	ProtoConfig  *gencode.ProtoSchema
	ModelConfig  *gencode.ModelSchema
}

func getDefaultConfig() *JsonConfig {
	jsonConfig := &JsonConfig{
		DBConfig:     &gencode.DBConfig{},
		GlobalConfig: &gencode.Config{},
		ApiConfig:    &gencode.ApiSchema{Dataset: &gencode.Dataset{SessionConfig: &gencode.SessionConfig{}}, ApiConfig: &gencode.ApiConfig{}},
		ProtoConfig:  &gencode.ProtoSchema{Dataset: &gencode.Dataset{SessionConfig: &gencode.SessionConfig{}}, ProtoConfig: &gencode.ProtoConfig{}},
	}
	return jsonConfig
}

func parseFlag(jsonConfig *JsonConfig) (err error) {
	templateFlag := flag.String("t", "", "input 'init', init template file")
	configFlag := flag.String("f", "genConfig.json", "json config file")
	flag.Parse()

	if *templateFlag == "init" {
		if err := initTemplate(); err != nil {
			return err
		}
	}

	if !strings.Contains(*configFlag, ".json") {
		fmt.Println(" - please input the json config file ")
		return nil
	}

	err = checkTemplate()
	if err != nil {
		return err
	}

	fileContent, _ := os.ReadFile(*configFlag)
	err = json.Unmarshal(fileContent, jsonConfig)
	if err != nil {
		return err
	}

	return nil
}

func checkTemplate() error {
	templateDir := filex.GetHomeDir() + "/" + gencode.TemplatePath
	exists, err := filex.PathExists(templateDir)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	return initTemplate()
}

func initTemplate() (err error) {
	templateDir := filex.GetHomeDir() + "/" + gencode.TemplatePath
	err = filex.CopyDirEm(dir, templateDir, "template")

	return err
}
