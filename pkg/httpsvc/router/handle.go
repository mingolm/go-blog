package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/utils/errutil"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
)

type Handler struct {
	routerMiddlewaresSet map[string][]middleware.Middleware
	routerSet            map[string]Router
	routerMethodSet      map[string]struct{}
	routerRegResult      []routerRegDetail
}

type routerRegDetail struct {
	reg         *regexp.Regexp
	method      string
	keys        []string
	middlewares []middleware.Middleware
	router      Router
}

func (h *Handler) HTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				zap.S().Errorw("panic for http handle",
					"err", err,
				)
			}
		}()
		resp, err := h.handle(r)
		var httpStatusCode int
		if err != nil {
			switch err {
			case errutil.ErrPageNotFound:
				httpStatusCode = http.StatusNotFound
			case errutil.ErrMethodNotAllowed:
				httpStatusCode = http.StatusMethodNotAllowed
			default:
				httpStatusCode = http.StatusInternalServerError
			}
			resp = response.Error(err)
		} else {
			httpStatusCode = http.StatusOK
		}

		w.Header().Set("Cache-Control", "no-cache, private")
		if redirectResponse, ok := resp.(interface {
			Redirect() (url string, code int)
		}); ok {
			url, code := redirectResponse.Redirect()
			http.Redirect(w, r, url, code)
			return
		}

		bs, err := resp.Bytes()
		if err != nil {
			httpStatusCode = http.StatusInternalServerError
			bs, _ = response.Error(err).Bytes()
		}

		w.WriteHeader(httpStatusCode)
		for key := range resp.Headers() {
			w.Header().Set(key, resp.GetHeader(key))
		}
		_, _ = w.Write(bs)
		_ = r.Body.Close()
	}
}

func (h *Handler) handle(r *http.Request) (resp response.Response, err error) {
	requestPath := strings.TrimRight(r.URL.Path, "/")
	routerIndex := fmt.Sprintf("%s#%s", requestPath, r.Method)
	router, ok := h.routerSet[routerIndex]
	middlewares, _ := h.routerMiddlewaresSet[routerIndex]
	if !ok {
		if _, methodNotAllowed := h.routerMethodSet[requestPath]; methodNotAllowed {
			return nil, errutil.ErrMethodNotAllowed
		}
		detailRouter, err := h.parseRegRouter(r, requestPath)
		if err != nil {
			return nil, err
		}
		router = detailRouter.router
		middlewares = detailRouter.middlewares
	}
	hl := router.Handler
	if middlewares != nil {
		for i := range middlewares {
			hl = middlewares[len(middlewares)-i-1](hl)
		}
	}

	return hl(r)
}

func (h *Handler) parseRegRouter(r *http.Request, path string) (router *routerRegDetail, err error) {
	for _, detail := range h.routerRegResult {
		regResults := detail.reg.FindAllStringSubmatch(path, -1)
		if len(regResults) == 0 {
			continue
		}
		if detail.method != r.Method {
			return nil, errutil.ErrMethodNotAllowed
		}
		for _, result := range regResults {
			if len(result[1:]) != len(detail.keys) {
				continue
			}
			for index, value := range result[1:] {
				if r.Form == nil {
					r.Form = make(map[string][]string, 0)
				}
				r.Form.Set(detail.keys[index], value)
			}
		}
		return &detail, nil
	}

	return nil, errutil.ErrPageNotFound
}
