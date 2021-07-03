package main

import (
	"context"
	"github.com/mingolm/go-recharge/configs"
	"github.com/mingolm/go-recharge/utils/argutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 配置初始化
	argutil.Parse(&configs.DefaultConfigs)

	// 日志初始化
	initLogger()

	// 程序退出信号监听
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel context.CancelFunc) {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		s := <-signalCh
		logger.Infof("signal received: %v", s)
		cancel()
	}(cancel)

	switch configs.DefaultConfigs.Run {
	case "http":
		runHttp(ctx)
	case "tmpjob":
		runJob(ctx, newOrderQueryJob())
	default:
		logger.Fatalw("unknown run flag",
			"value", configs.DefaultConfigs.Run,
		)
	}
	os.Exit(0)
}
