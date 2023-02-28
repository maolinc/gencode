package main

import (
	gencode "gencode/core"
	"log"
)

func main() {

	var err error

	jsonConfig := gencode.GetDefaultConfig()
	err = gencode.ParseJson(jsonConfig)
	if err != nil {
		log.Println(err)
	}

	dataset := gencode.From(jsonConfig.DBConfig, jsonConfig.GlobalConfig)
	//
	//// IgnoreFieldValue desc:
	//// value and ignore mapping rule: 1(create), 2(update),4(select),8(delete), 1+2=3(create,update),1+2+4=7(create,select,update)
	//// eg: {create_time:1, delete_flag:7, id:1}
	apiSchema := gencode.NewApiSchema(dataset.Session(jsonConfig.ApiConfig.SessionConfig), jsonConfig.ApiConfig.ApiConfig)

	protoSchema := gencode.NewProtoSchema(dataset.Session(jsonConfig.ProtoConfig.SessionConfig), jsonConfig.ProtoConfig.ProtoConfig)

	gencode.Generates(apiSchema, protoSchema)
}
