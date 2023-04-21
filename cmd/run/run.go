package run

import (
	"fmt"
	"go_project_template/api/router"
	"go_project_template/internal/conf"
	"go_project_template/internal/data/persistence"
	"os"

	"go_project_template/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitToRun() *gin.Engine {
	r := gin.Default()
	router.InitRouter(r)                                                                                                                       // 注入路由
	err := logger.InitLogger(os.Getenv("RUN_MODE"), conf.GetEnv().EnableKafka, conf.GetEnv().KafkaConfig.KafkaAddress, conf.GetEnv().LogTopic) // 初始化日志
	if err != nil {
		fmt.Println("init logger err:", zap.Error(err))
		os.Exit(1)
		return nil
	}

	db, err := persistence.NewDB(conf.GetEnv().Data, &gorm.Config{}) // 初始化数据库
	if err != nil {
		logger.Logger.Error("init db err:", zap.Error(err))
		os.Exit(1)
		return nil
	}
	persistence.NewProductRepo(db) // 注入数据库
	logger.Logger.Info("启动")
	return r

}
