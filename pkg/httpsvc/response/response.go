package response

type Response interface {
	Headers() (headers map[string]string)
	AddHeader(key, value string)
	GetHeader(key string) (value string)
	Bytes() (bs []byte, err error)
}

var ErrInternalBytes []byte
var Success Response

func init() {
	ErrInternalBytes, _ = (&httpResponse{
		Data: "",
	}).Bytes()
	Success = &httpResponse{
		Data: "ok",
	}
}

func Data(v interface{}) Response {
	return &httpResponse{
		Data:    v,
		Success: true,
	}
}

func Html(filename string, v interface{}) Response {
	return &htmlResponse{
		Filename: filename,
		Data:     v,
	}
}

func Error(err error) Response {
	return &httpResponse{
		Data:    err.Error(),
		Success: false,
	}
}

type Common struct {
	Message string `json:"message"`
}
