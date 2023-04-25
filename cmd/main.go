package main

import (
	"go_project_template/cmd/run"
	"go_project_template/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	r := run.InitToRun()
	err := r.Run(":8080")
	if err != nil {
		logger.Logger.Panic("%v", zap.Error(err))
	}
}
