package response

import "encoding/json"

type redirectResponse struct {
	Url     string
	Code    int
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
func (h *redirectResponse) Bytes() (bs []byte, err error) {
	return json.Marshal(h)
}
func (h *redirectResponse) Redirect() (url string, code int) {
	return h.Url, h.Code
}
