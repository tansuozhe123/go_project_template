package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

type LogKafka struct {
	Producer *kafka.Producer
	Topic    string
}

func (lk *LogKafka) Write(p []byte) (n int, err error) {
	if !lk.SendKafkaMessage(lk.Topic, string(p)) {
		fmt.Println("write kafka log error")
		return 0, errors.New("write kafka log error")
	}
	return len(p), nil
}

/*
*
发送kafka日志
*/
func (lk *LogKafka) SendKafkaMessage(topic string, message string) bool {
	success := false
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	err := lk.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	e := <-deliveryChan
	defer func() {
		close(deliveryChan)
	}()
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil && err != nil {
		fmt.Println("send kafka log error", zap.Any("topic", topic), zap.Any("message", message), zap.Error(err))
	} else {
		fmt.Println("send kafka log done", zap.Any("topic", *m.TopicPartition.Topic), zap.Any("partition", m.TopicPartition.Partition), zap.Any("offset", m.TopicPartition.Offset))
		success = true
	}
	return success
}

func InitLogger(mode string, enableKafka bool, kafkaAddress string, topic string) error {
	//测试环境不输出日志到kfaka
	if mode == "test" {
		Logger = zap.NewNop()
		return nil
	}
	// 打印错误级别的日志
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	var allCore []zapcore.Core

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// for human operators.
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	// kafka
	if len(kafkaAddress) > 0 && enableKafka {
		var (
			kl  LogKafka
			err error
		)
		var conf = kafka.ConfigMap{}
		conf.SetKey("bootstrap.servers", kafkaAddress)
		conf.SetKey("acks", "all")
		conf.SetKey("max.in.flight.requests.per.connection", 1)
		conf.SetKey("message.send.max.retries", 3)
		kl.Producer, err = kafka.NewProducer(&conf)
		kl.Topic = topic
		if err != nil {
			fmt.Printf("Failed to create producer: %s", err)
			return err
		}
		topicErrors := zapcore.AddSync(&kl)
		// 打印在kafka
		kafkaEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		var kafkaCore zapcore.Core
		if mode == "debug" {
			kafkaCore = zapcore.NewCore(kafkaEncoder, topicErrors, lowPriority)
		} else {
			kafkaCore = zapcore.NewCore(kafkaEncoder, topicErrors, highPriority)

		}
		allCore = append(allCore, kafkaCore)

	}
	if mode == "debug" {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	}

	core := zapcore.NewTee(allCore...)

	// From a zapcore.Core, it's easy to construct a Logger.
	Logger = zap.New(core).WithOptions(zap.AddCaller())
	return nil
}
