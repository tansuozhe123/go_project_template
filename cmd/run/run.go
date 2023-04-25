package run

import (
	"fmt"
	"go_project/api/router"
	"go_project/internal/conf"
	"go_project/internal/data/persistence"
	"go_project/pkg/logger"
	mongodb "go_project/pkg/mongo"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitToRun() *gin.Engine {
	r := gin.Default()
	//注入路由
	router.InitRouter(r)
	// 初始化日志
	err := logger.InitLogger(os.Getenv("RUN_MODE"), conf.GetEnv().EnableKafka, conf.GetEnv().KafkaConfig.KafkaAddress, conf.GetEnv().LogTopic)
	if err != nil {
		fmt.Println("init logger err:", zap.Error(err))
		os.Exit(1)
		return nil
	}
	// 初始化mysql
	mysqlDb, err := persistence.NewDB(conf.GetEnv().Data, &gorm.Config{})
	conf.GetEnv().MySQLCli = *mysqlDb
	if err != nil {
		logger.Logger.Error("init db err:", zap.Error(err))
		os.Exit(1)
		return nil
	}
	// 初始化mongodb
	mongoDb, err := mongodb.InitConnect(&conf.GetEnv().MongoCfg)
	conf.GetEnv().MongoCli = *mongoDb
	if err != nil {
		logger.Logger.Error("init mongo err:", zap.Error(err))
		os.Exit(1)
		return nil
	}
	// persistence.NewProductRepo(conf.GetEnv().MySQLCli, conf.GetEnv().MongoCli) // 注入数据库
	logger.Logger.Sugar().Info(conf.GetEnv().MongoCli)
	return r

}
