package ctrl

import (
	"github.com/mingolm/go-recharge/pkg/cache"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/pagingutil"
	"net/http"
	"strconv"
)

func NewLeave() *Leave {
	return &Leave{
		Service:    core.Instance(),
		leaveCache: cache.NewLeaveCache(),
		limit:      10,
	}
}

type Leave struct {
	*core.Service
	leaveCache cache.LeaveCache
	limit      int
}

func (s *Leave) Routers() router.Routers {
	return []router.Router{
		{ // 归档
			Path:    "/leave",
			Handler: s.leave,
			Method:  "GET",
		},
	}
}

func (s *Leave) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *Leave) leave(req *http.Request) (resp response.Response, err error) {
	currentPage, _ := strconv.Atoi(req.FormValue("page"))
	offset := currentPage * s.limit
	rows, err := s.ArticleRepo.GetList(req.Context(), offset, s.limit)
	if err != nil {
		return nil, err
	}

	total, err := s.leaveCache.GetTotal(req.Context())
	if err != nil {
		return nil, err
	}

	return response.Html("leave", struct {
		Rows  []*model.Article   `json:"rows"`
		Pages *pagingutil.Paging `json:"pages"`
	}{
		Rows:  rows,
		Pages: pagingutil.Paginator(currentPage, s.limit, int(total)),
	}), nil
}
