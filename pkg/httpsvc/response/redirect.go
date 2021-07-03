package response

import (
	"encoding/json"
	"net/http"
)

type redirectResponse struct {
	Url     string
	Code    int
	Cookies []*http.Cookie
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func (h *redirectResponse) Headers() (headers map[string]string) {
	return
}
func (h *redirectResponse) AddHeader(key, value string) {
	return
}
func (h *redirectResponse) GetHeader(key string) (value string) {
	return
}
func (h *redirectResponse) WithCookie(cookie *http.Cookie) (ins Response) {
	h.Cookies = append(h.Cookies, cookie)
	return h
}
func (h *redirectResponse) GetCookie() (cookies []*http.Cookie) {
	return h.Cookies
}
func (h *redirectResponse) Bytes() (bs []byte, err error) {
	return json.Marshal(h)
}
func (h *redirectResponse) Redirect() (url string, code int) {
	return h.Url, h.Code
}
