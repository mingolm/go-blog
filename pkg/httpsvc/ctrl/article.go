package ctrl

import (
	"github.com/mingolm/go-recharge/pkg/cache"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/utils/errutil"
	"net/http"
	"strconv"
)

func NewArticle() *Article {
	return &Article{
		Service:      core.Instance(),
		articleCache: cache.NewArticleCache(),
	}
}

type Article struct {
	*core.Service
	articleCache cache.ArticleCache
}

func (s *Article) Routers() router.Routers {
	return []router.Router{
		{ // 归档
			Path:    "/article",
			Handler: s.article,
			Method:  "GET",
		},
		{ // 文章详情
			Path:    "/article/{id}",
			Handler: s.articleDetail,
			Method:  "GET",
		},
	}
}

func (s *Article) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *Article) article(req *http.Request) (resp response.Response, err error) {
	rows, err := s.articleCache.GetList(req.Context())
	if err != nil {
		return nil, err
	}
	return response.Html("article", rows), nil
}

func (s *Article) articleDetail(req *http.Request) (resp response.Response, err error) {
	id, err := strconv.ParseUint(req.FormValue("id"), 10, 64)
	if err != nil {
		return nil, errutil.ErrInvalidArguments.Msg("id")
	}
	row, err := s.articleCache.Get(req.Context(), id)
	if err != nil {
		return nil, err
	}
	return response.Html("article_detail", row), nil
}
