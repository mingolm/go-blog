package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/utils/errutil"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	routerMiddlewaresSet map[string][]middleware.Middleware
	routerSet            map[string]Router
	routerMethodSet      map[string]struct{}
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
	routerIndex := fmt.Sprintf("%s#%s", r.URL.Path, r.Method)
	router, ok := h.routerSet[routerIndex]
	if !ok {
		if _, methodNotAllowed := h.routerMethodSet[r.URL.Path]; methodNotAllowed {
			return nil, errutil.ErrMethodNotAllowed
		}
		return nil, errutil.ErrPageNotFound
	}
	hl := router.Handler
	if middlewares, ok := h.routerMiddlewaresSet[routerIndex]; ok {
		for i := range middlewares {
			hl = middlewares[len(middlewares)-i-1](hl)
		}
	}

	return hl(r)
}
