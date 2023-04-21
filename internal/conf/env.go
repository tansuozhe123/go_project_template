package conf

import (
	"fmt"
	"os"
)

var (
	env = Env{
		MySQLConfig: &MySQLConfig{
			UserName:     "root",
			Password:     "123456",
			Host:         "127.0.0.1",
			Port:         "3306",
			DatabaseName: "test",
			Charset:      "utf8mb4",
			ParseTime:    "True",
			Loc:          "Local",
		},
		KafkaConfig: &KafkaConfig{
			KafkaAddress: "127.0.0.1:9092",
		},
		EnableKafka: true,
		LogTopic:    "product",
		Data: &Data{
			Database: &Data_Database{
				Driver: "mysql", // 数据库类型
				Source: "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
			},
		},
	}
)

// 初始化环境配置
func init() {
	if osenv := os.Getenv("RUN_MODE"); osenv != "debug" {
		env = Env{
			MySQLConfig: &MySQLConfig{
				UserName:     os.Getenv("MYSQL_USERNAME"),
				Password:     os.Getenv("MYSQL_PASSWORD"),
				Host:         os.Getenv("MYSQL_HOST"),
				Port:         os.Getenv("MYSQL_PORT"),
				DatabaseName: os.Getenv("MYSQL_DATABASENAME"),
				Charset:      os.Getenv("MYSQL_CHARSET"),
				ParseTime:    os.Getenv("MYSQL_PARSETIME"),
				Loc:          os.Getenv("MYSQL_LOC"),
			},
			KafkaConfig: &KafkaConfig{
				KafkaAddress: os.Getenv("KAFKA_ADDRESS"),
			},
			LogTopic:    os.Getenv("LOG_TOPIC"),
			EnableKafka: os.Getenv("ENABLE_KAFKA") == "true",
			Data: &Data{
				Database: &Data_Database{
					Driver: "mysql", // 数据库类型
					Source: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
						os.Getenv("MYSQL_DATABASENAME"), os.Getenv("MYSQL_CHARSET"), os.Getenv("MYSQL_PARSETIME"), os.Getenv("MYSQL_LOC")),
				},
			},
		}
	}
}

func GetEnv() *Env {
	return &env
}
