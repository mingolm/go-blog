package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net/http"
)

func New() *RI {
	return &RI{
		handlers: make([]RouterHandler, 0),
	}
}

type RI struct {
	handlers []RouterHandler
}

func (r *RI) Register(hls ...RouterHandler) *RI {
	r.handlers = append(r.handlers, hls...)
	return r
}

func (r *RI) HTTPRouters() *Handler {
	routerSet := make(map[string]*Router, 0)
	routerMethodSet := make(map[string]struct{}, 0)
	for _, hl := range r.handlers {
		for _, router := range hl.Routers() {
			routerSet[fmt.Sprintf("%s#%s", router.Path, router.Method)] = &router
			routerMethodSet[router.Path] = struct{}{}
		}
	}
	return &Handler{
		routerSet:       routerSet,
		routerMethodSet: routerMethodSet,
	}
}

type Router struct {
	Path    string
	Method  string
	Handler func(req *http.Request) (resp response.Response, err error)
}

type Routers []Router

type RouterHandler interface {
	Routers() Routers
}
