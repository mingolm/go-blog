package response

import (
	"encoding/json"
	"net/http"
)

var ErrInternalBytes []byte

func init() {
	ErrInternalBytes, _ = json.Marshal(&Response{
		Message: "",
		Success: false,
	})
}

type Response struct {
	Message string      `json:"message"`
	Header  http.Header `json:"-"`
	Success bool        `json:"success"`
}

func (r *Response) SetHeader(key, value string) {
	r.Header.Set(key, value)
}

func (r *Response) Headers() http.Header {
	return r.Header
}

func (r *Response) Bytes() (bs []byte, err error) {
	bs, err = json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
