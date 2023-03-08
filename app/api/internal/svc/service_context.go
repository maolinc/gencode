package svc

import (
	"github.com/maolinc/gencode/app/api/internal/config"
	"github.com/maolinc/gencode/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config         config.Config
	MArtifactModel model.MArtifactModel
	MEnvModel      model.MEnvModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	return &ServiceContext{
		Config:         c,
		MArtifactModel: model.NewMArtifactMode(db, c.Cache),
		MEnvModel:      model.NewMEnvMode(db, c.Cache),
	}
}
