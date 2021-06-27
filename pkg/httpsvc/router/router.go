package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net/http"
)

func New() *RI {
	return &RI{
		middlewareHandlers: make([]middleware.Middleware, 0),
		handlers:           make([]RouterHandler, 0),
	}
}

type RI struct {
	middlewareHandlers []middleware.Middleware
	handlers           []RouterHandler
}

func (r *RI) RegisterMiddleware(ml ...middleware.Middleware) *RI {
	r.middlewareHandlers = append(r.middlewareHandlers, ml...)
	return r
}

func (r *RI) Register(hls ...RouterHandler) *RI {
	r.handlers = append(r.handlers, hls...)
	return r
}

func (r *RI) HTTPRouters() *Handler {
	routerMiddlewaresSet := make(map[string][]middleware.Middleware, 0)
	routerSet := make(map[string]Router, 0)
	routerMethodSet := make(map[string]struct{}, 0)
	for _, hl := range r.handlers {
		var middlewares []middleware.Middleware
		middlewares = append(middlewares, r.middlewareHandlers...)
		middlewares = append(middlewares, hl.Middlewares()...)
		for _, router := range hl.Routers() {
			routerIndex := fmt.Sprintf("%s#%s", router.Path, router.Method)
			routerMiddlewaresSet[routerIndex] = append(middlewares, router.Middlewares...)
			routerSet[routerIndex] = router
			routerMethodSet[router.Path] = struct{}{}
		}
	}
	return &Handler{
		routerMiddlewaresSet: routerMiddlewaresSet,
		routerSet:            routerSet,
		routerMethodSet:      routerMethodSet,
	}
}

type Router struct {
	Path        string
	Method      string
	Middlewares []middleware.Middleware
	Handler     func(req *http.Request) (resp response.Response, err error)
}

type Routers []Router

type RouterHandler interface {
	Routers() Routers
	Middlewares() []middleware.Middleware
}
