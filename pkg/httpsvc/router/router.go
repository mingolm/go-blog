package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net/http"
	"regexp"
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
	// 可能存在参数，`{}` 包裹表示可替换参数
	routerReg := regexp.MustCompile("(?U){(.+)}")

	routerMiddlewaresSet := make(map[string][]middleware.Middleware, 0)
	routerSet := make(map[string]Router, 0)
	routerRegResults := make([]routerRegDetail, 0)
	routerMethodSet := make(map[string]struct{}, 0)
	for _, hl := range r.handlers {
		var middlewares []middleware.Middleware
		middlewares = append(middlewares, r.middlewareHandlers...)
		middlewares = append(middlewares, hl.Middlewares()...)
		for _, router := range hl.Routers() {
			keySubs := routerReg.FindAllStringSubmatch(router.Path, -1)
			if len(keySubs) != 0 {
				keys := make([]string, 0)
				for _, sub := range keySubs {
					keys = append(keys, sub[1:]...)
				}
				routerPathResult := routerReg.ReplaceAllString(router.Path, "([0-9a-zA-Z\\-_]+)")
				routerPathResultReg := regexp.MustCompile(fmt.Sprintf("(?U)^%s$", routerPathResult))
				routerRegResults = append(routerRegResults, routerRegDetail{
					reg:         routerPathResultReg,
					method:      router.Method,
					keys:        keys,
					middlewares: middlewares,
					router:      router,
				})
			} else {
				routerIndex := fmt.Sprintf("%s#%s", router.Path, router.Method)
				routerMiddlewaresSet[routerIndex] = append(middlewares, router.Middlewares...)
				routerSet[routerIndex] = router
				routerMethodSet[router.Path] = struct{}{}
			}
		}
	}
	return &Handler{
		routerMiddlewaresSet: routerMiddlewaresSet,
		routerSet:            routerSet,
		routerMethodSet:      routerMethodSet,
		routerRegResult:      routerRegResults,
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
