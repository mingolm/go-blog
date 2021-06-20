package middleware

import "net/http"

func New() *Ml {
	return &Ml{
		Middlewares: make([]Middleware, 0),
	}
}

type Middleware func(next http.Handler) http.Handler

type Ml struct {
	Middlewares []Middleware
}

func (ml *Ml) Add(m Middleware) {
	ml.Middlewares = append(ml.Middlewares, m)
}

func (ml *Ml) Handle(hl http.Handler) http.Handler {
	for _, m := range ml.Middlewares {
		hl = m(hl)
	}
	return hl
}
