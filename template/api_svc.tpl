{{ $Cache := .IsCache -}}
package svc

import (
	{{.ModelPkg}}
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config         config.Config
	{{range  .Dataset.TableSet -}}
	{{.CamelName}}Model  model.{{.CamelName}}Model
	{{end -}}
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	return &ServiceContext{
		Config:         c,
        {{range  .Dataset.TableSet -}}
        {{.CamelName}}Model:  model.New{{.CamelName}}Model(db {{- if $Cache}}, c.Cache{{end}}),
        {{end -}}
	}
}
