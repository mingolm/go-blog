package main

import (
	"go.uber.org/zap"
)

var rawLogger *zap.Logger
var logger *zap.SugaredLogger

func initLogger() {
	var options []zap.Option
	if defaultConfigs.mode == "prod" {
		rawLogger, _ = zap.NewProduction(options...)
	} else {
		rawLogger, _ = zap.NewDevelopment(options...)
	}
	logger = rawLogger.Sugar()
	zap.ReplaceGlobals(rawLogger)
}
