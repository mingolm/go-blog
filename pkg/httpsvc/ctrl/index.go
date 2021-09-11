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

func NewIndex() *Index {
	return &Index{
		Service:      core.Instance(),
		articleCache: cache.NewArticleCache(),
		Limit:        10,
	}
}

type Index struct {
	*core.Service
	articleCache cache.ArticleCache
	Limit        int
}

func (s *Index) Routers() router.Routers {
	return []router.Router{
		{ // 首页
			Path:    "/index",
			Handler: s.index,
			Method:  "GET",
		},
	}
}

func (s *Index) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *Index) index(req *http.Request) (resp response.Response, err error) {
	currentPage, _ := strconv.Atoi(req.FormValue("page"))
	offset := currentPage * s.Limit
	rows, err := s.ArticleRepo.GetList(req.Context(), int(offset), int(s.Limit))
	if err != nil {
		return nil, err
	}

	total, err := s.articleCache.GetTotal(req.Context())
	if err != nil {
		return nil, err
	}

	return response.Html("index", struct {
		Rows  []*model.Article   `json:"rows"`
		Pages *pagingutil.Paging `json:"pages"`
	}{
		Rows:  rows,
		Pages: pagingutil.Paginator(currentPage, s.Limit, int(total)),
	}), nil
}
