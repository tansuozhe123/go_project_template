package persistence

import (
	"go_project_template/internal/biz/entity"
	"go_project_template/internal/conf"

	"go_project_template/pkg/logger"
	mongodb "go_project_template/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../../../mocks/product_persitence_mock.go -package=mocks go_project_template/internal/data/persistence ProductPersitenceInf

type ProductPersitenceInf interface {
	GetOneProduct(productkey string) (*entity.Product, error)
}
type ProductPersitence struct {
	MySQlDb *gorm.DB
	MongoDb *mongodb.MongoClient
}

var ProductPersi = &ProductPersitence{
	MySQlDb: &conf.GetEnv().MySQLCli,
	MongoDb: &conf.GetEnv().MongoCli,
}

func (model *ProductPersitence) GetOneProduct(productkey string) (*entity.Product, error) {
	var val entity.Product
	ret, err := model.MongoDb.FindOne("product", bson.M{"productkey": productkey})
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		logger.Logger.Sugar().Error("find product err:", err)
		return nil, err
	}
	if err := ret.Decode(&val); err != nil {
		logger.Logger.Sugar().Error("decode product err:", err)
	}
	return &val, nil
}
