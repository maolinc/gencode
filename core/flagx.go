package gencode

import (
	"encoding/json"
	"flag"
	"os"
)

type JsonConfig struct {
	DBConfig     *DBConfig
	GlobalConfig *Config
	ApiConfig    *ApiSchema
	ProtoConfig  *ProtoSchema
}

func GetDefaultConfig() *JsonConfig {
	jsonConfig := &JsonConfig{
		DBConfig:     &DBConfig{},
		GlobalConfig: &Config{},
		ApiConfig:    &ApiSchema{Dataset: &Dataset{SessionConfig: &SessionConfig{}}, ApiConfig: &ApiConfig{}},
		ProtoConfig:  &ProtoSchema{Dataset: &Dataset{SessionConfig: &SessionConfig{}}, ProtoConfig: &ProtoConfig{}},
	}
	return jsonConfig
}

func ParseJson(jsonConfig *JsonConfig) (err error) {
	configFile := flag.String("f", "genConfig.json", "")
	flag.Parse()

	fileContent, _ := os.ReadFile(*configFile)
	err = json.Unmarshal(fileContent, jsonConfig)
	if err != nil {
		return err
	}
	return nil
}
