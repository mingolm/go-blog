package main

import (
	"context"
	"fmt"
	"github.com/mingolm/go-recharge/configs"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router/app"
	"net"
	"net/http"
	"time"
)

func runHttp(ctx context.Context) {
	svc, shutdownCallback := httpServer()
	listener, err := net.Listen("tcp", svc.Addr)
	if err != nil {
		logger.Fatalw("could not listen on port",
			"listen", svc.Addr,
			"err", err,
		)
	}

	go func() {
		if err := svc.Serve(listener); err != nil {
			if err != http.ErrServerClosed {
				logger.Errorw("http server fail",
					"err", err,
				)
			}
		}
	}()

	logger.Infow("http server is running",
		"listen", svc.Addr,
	)

	// 平滑退出
	<-ctx.Done()

	// 缓冲时间
	time.Sleep(time.Second)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	if err := svc.Shutdown(shutdownCtx); err != nil {
		if err != http.ErrServerClosed {
			logger.Errorw("http server close failed",
				"err", err,
			)
		}
	}
	cancel()

	shutdownCallback()

	logger.Info("bye!")
}

func httpServer() (h *http.Server, shutdownCallback func()) {
	svcRouter := router.New()
	routers := svcRouter.RegisterMiddleware(
		middleware.Context,
		middleware.ReteLimit,
	).Register(app.NewLogin(), app.NewApp()).HTTPRouters()

	shutdownCallback = func() {
		fmt.Println("shutdown")
	}

	return &http.Server{
		Addr: configs.DefaultConfigs.HttpListen,
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			// 健康检查
			if request.URL.Path == "/health" {
				writer.WriteHeader(http.StatusOK)
				_, _ = writer.Write([]byte("ok"))
				return
			}
			routers.HTTPHandler().ServeHTTP(writer, request)
		}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}, shutdownCallback
}
