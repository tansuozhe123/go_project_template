package persistence

import (
	"go_project_template/internal/conf"
	"go_project_template/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *conf.Data, gormConf *gorm.Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("failed opening connection to mysql: %v", zap.Error(err))
		return nil, err
	}
	return db, nil
}
