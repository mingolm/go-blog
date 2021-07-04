package httputil

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/mingolm/go-recharge/utils/errutil"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	defaultTimeout = time.Second * 5
)

func NewHTTPClient(config *HTTPClientConfig) *HTTPClient {
	if config.Timeout != 0 {
		defaultTimeout = config.Timeout
	}
	cli := &http.Client{
		Timeout: defaultTimeout,
	}
	if config.Proxy != "" {
		proxyUrl, _ := url.Parse(config.Proxy)
		cli.Transport = &http.Transport{
			Proxy:           http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.InsecureSkipVerify},
		}
	}

	return &HTTPClient{
		client:  cli,
		headers: make(map[string]string, 0),
	}
}

type HTTPClient struct {
	client  *http.Client
	headers map[string]string
}

type HTTPClientConfig struct {
	Proxy              string
	InsecureSkipVerify bool
	Timeout            time.Duration
}

func (h *HTTPClient) Proxy(proxy string) *HTTPClient {
	proxyUrl, _ := url.Parse(proxy)
	h.client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	return h
}

func (h *HTTPClient) AddHeader(headers map[string]string) *HTTPClient {
	for key, value := range headers {
		h.headers[key] = value
	}
	return h
}

func (h *HTTPClient) Get(path string, queries map[string]string) (bs []byte, err error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	q := u.Query()
	for key, value := range queries {
		q.Set(key, value)
	}
	request, err := h.buildRequest("GET", fmt.Sprintf("%s?%s", path, q.Encode()), nil)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	resp, err := h.client.Do(request)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	defer resp.Body.Close()
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	return bs, nil
}

func (h *HTTPClient) PostForm(fullPath string, body map[string]string) (bs []byte, err error) {
	h.AddHeader(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})

	form := url.Values{}
	for key, value := range body {
		form.Set(key, value)
	}
	request, err := h.buildRequest("POST", fullPath, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	resp, err := h.client.Do(request)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	defer resp.Body.Close()
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	return bs, nil
}

func (h *HTTPClient) PostJson(fullPath string, body map[string]interface{}) (bs []byte, err error) {
	h.AddHeader(map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	})
	bodyBs, err := json.Marshal(body)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	request, err := h.buildRequest("POST", fullPath, strings.NewReader(string(bodyBs)))
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	resp, err := h.client.Do(request)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	defer resp.Body.Close()
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	return bs, nil
}

func (h *HTTPClient) buildRequest(method, fullPath string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}
	for key, value := range h.headers {
		request.Header.Add(key, value)
	}

	return request, nil
}
