package ctrl

import (
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"net/http"
)

func NewAbout() *About {
	return &About{
		Service: core.Instance(),
	}
}

type About struct {
	*core.Service
}

func (s *About) Routers() router.Routers {
	return []router.Router{
		{ // 关于我
			Path:    "/about",
			Handler: s.about,
			Method:  "GET",
		},
	}
}

func (s *About) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *About) about(req *http.Request) (resp response.Response, err error) {
	return response.Html("about", nil), nil
}
