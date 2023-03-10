{{ $Cache := .IsCache -}}
/*
1. The following are the new dependencies and packages.
2. You must copy and paste them into the above ServiceContext.
3. The comment is to avoid overwriting the previous dependencies, Delete the comment section after completion 2.

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
    sc := &ServiceContext{
        Config:         c,
        {{range  .Dataset.TableSet}}{{.CamelName}}Model:  model.New{{.CamelName}}Model(db {{- if $Cache}}, c.Cache{{end}}),
        {{end}}}
	return sc
}
*/