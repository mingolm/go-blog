package response

import "encoding/json"

type httpResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
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
func (h *httpResponse) Bytes() (bs []byte, err error) {
	return json.Marshal(h)
}
