package main

import (
	"context"
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router/app"
	"net"
	"net/http"
	"time"
)

func runHttp(ctx context.Context) {
	svc, afterCloseFns := httpServer()
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
	if err := svc.Shutdown(ctx); err != nil {
		if err != http.ErrServerClosed {
			logger.Errorw("http server close fail,",
				"err", err,
			)
		}
	}
	for _, fn := range afterCloseFns {
		fn()
	}

	logger.Info("bye!")
}

func httpServer() (h *http.Server, afterCLoseFns []func()) {
	afterCLoseFns = append(afterCLoseFns, func() {
		fmt.Println("close")
	})

	svcRouter := router.New()
	routers := svcRouter.Register(&app.Login{

	}).HTTPRouters()

	return &http.Server{
		Addr: defaultConfigs.httpListen,
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			// 健康检查
			if request.URL.Path == "/healthz" {
				writer.WriteHeader(http.StatusOK)
				_, _ = writer.Write([]byte("ok"))
				return
			}
			ml := middleware.New()
			ml.Add(middleware.CsrfToken())
			ml.Handle(routers.HTTPHandler()).ServeHTTP(writer, request)
		}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}, afterCLoseFns
}
