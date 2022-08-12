package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kafka_mate_go/pkg/api"
	"kafka_mate_go/pkg/config"
	"kafka_mate_go/pkg/kafka"
	"kafka_mate_go/pkg/util"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	if !config.RemoteMode {
		err := kafka.Config()
		if err != nil {
			util.Logger().Error("generate config failed ", zap.Error(err))
			os.Exit(1)
		}
		kafka.Start()
	}
	engine := gin.Default()
	engine.GET("/readiness", api.Readiness)
	engine.Run(":31008")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
