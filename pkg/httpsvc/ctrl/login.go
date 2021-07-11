package ctrl

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"net/http"
)

func NewLogin() *Login {
	return &Login{
		core.Instance(),
	}
}

type Login struct {
	*core.Service
}

func (s *Login) Routers() router.Routers {
	return []router.Router{
		{
			Path:    "/login",
			Handler: s.LoginTemplate,
			Method:  "GET",
			Middlewares: []middleware.Middleware{
				middleware.Authentication,
			},
		},
		{
			Path:    "/login",
			Handler: s.Login,
			Method:  "POST",
		},
		{
			Path:    "/register",
			Handler: s.Register,
			Method:  "POST",
		},
	}
}

func (s *Login) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *Login) LoginTemplate(req *http.Request) (resp response.Response, err error) {
	return response.Html("login", nil), nil
}

func (s *Login) Login(req *http.Request) (resp response.Response, err error) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		return response.Error(fmt.Errorf("login: username or password is empty")), nil
	}

	userRow, err := s.UserRepo.GetForLogin(req.Context(), username, password)
	if err != nil {
		return nil, err
	}
	if userRow.Status != model.UserStatusNormal {
		return nil, errutil.ErrFailedPrecondition.Msg("user disabled")
	}

	s.Logger.Infow("login success",
		"id", userRow.ID,
	)

	return response.Redirect("index", 302), nil
}

func (s *Login) Register(req *http.Request) (resp response.Response, err error) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		return response.Error(fmt.Errorf("login: username or password is empty")), nil
	}

	if err := s.UserRepo.Create(req.Context(), &model.User{
		Username: username,
		Password: password,
		Status:   model.UserStatusNormal,
		IP:       model.GetIPv4(req.Context().Value("ip").(string)),
	}); err != nil {
		return nil, err
	}

	return response.Redirect("login", 302), nil
}
