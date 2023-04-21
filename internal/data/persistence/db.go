package persistence

import (
	"go_project_template/internal/conf"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *conf.Data, gormConf *gorm.Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	return db, nil
}
