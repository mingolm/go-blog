package main

import (
	"context"
	"github.com/mingolm/go-recharge/utils/argutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 配置初始化
	argutil.Parse(&defaultConfigs)

	// 日志初始化
	initLogger()

	// 程序退出信号监听
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel context.CancelFunc) {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		s := <-signalCh
		logger.Infof("signal received: %v", s)
		time.Sleep(time.Second)
		cancel()
	}(cancel)

	switch defaultConfigs.run {
	case "http":
		runHttp(ctx)
	case "tmpjob":
	default:
		logger.Fatalw("unknown run flag",
			"value", defaultConfigs.run,
		)
	}
	os.Exit(0)
}
