package main

import (
	"github.com/mingolm/go-recharge/configs"
	"go.uber.org/zap"
)

var rawLogger *zap.Logger
var logger *zap.SugaredLogger

func initLogger() {
	var options []zap.Option
	if configs.DefaultConfigs.Mode == "prod" {
		rawLogger, _ = zap.NewProduction(options...)
	} else {
		rawLogger, _ = zap.NewDevelopment(options...)
	}
	logger = rawLogger.Sugar()
	zap.ReplaceGlobals(rawLogger)
}
