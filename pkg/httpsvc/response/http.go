package response

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Data    interface{}    `json:"data"`
	Cookies []*http.Cookie `json:"-"`
	Success bool           `json:"success"`
}

func (h *httpResponse) Headers() (headers map[string]string) {
	return
}
func (h *httpResponse) AddHeader(key, value string) {
	return
}
func (h *httpResponse) GetHeader(key string) (value string) {
	return
}
func (h *httpResponse) WithCookie(cookie *http.Cookie) (ins Response) {
	h.Cookies = append(h.Cookies, cookie)
	return h
}
func (h *httpResponse) GetCookie() (cookies []*http.Cookie) {
	return h.Cookies
}
func (h *httpResponse) Bytes() (bs []byte, err error) {
	return json.Marshal(h)
}
