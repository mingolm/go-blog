package router

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/utils/errutil"
	"net/http"
)

type Handler struct {
	routerSet       map[string]*Router
	routerMethodSet map[string]struct{}
}

func (h *Handler) HTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		bs, err := resp.Bytes()
		if err != nil {
			httpStatusCode = http.StatusInternalServerError
			bs = response.ErrInternalBytes
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
	router, ok := h.routerSet[fmt.Sprintf("%s#%s", r.URL.Path, r.Method)]
	if !ok {
		if _, methodNotAllowed := h.routerMethodSet[r.URL.Path]; methodNotAllowed {
			return nil, errutil.ErrMethodNotAllowed
		}
		return nil, errutil.ErrPageNotFound
	}
	return router.Handler(r)
}
