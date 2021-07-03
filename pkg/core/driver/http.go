package driver

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpClient struct {
	client *http.Client
	header map[string]string
}

type httpConfig struct {
	proxy              string
	insecureSkipVerify bool
	timeout            time.Duration
}

// 默认 5 秒超时
var defaultTimeout = time.Second * 5

func dhc() *httpClient {
	return hc(&httpConfig{
		timeout: defaultTimeout,
	})
}

func hc(conf *httpConfig) *httpClient {
	cli := &http.Client{
		Timeout: conf.timeout,
	}
	if conf.proxy != "" {
		proxy, _ := url.Parse(conf.proxy)
		cli.Transport = &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.insecureSkipVerify},
		}
	}
	return &httpClient{
		client: cli,
		header: make(map[string]string, 0),
	}
}

func (h *httpClient) Header(header map[string]string) *httpClient {
	for key, value := range header {
		h.header[key] = value
	}
	return h
}

func (h *httpClient) Get(path string, queries map[string]string) (responseBs []byte, err error) {
	form := url.Values{}
	for pk, pv := range queries {
		form.Add(pk, pv)
	}

	if encodeQuery := form.Encode(); encodeQuery != "" {
		if strings.Contains(path, "?") {
			path = fmt.Sprintf("%s&%s", path, form.Encode())
		} else {
			path = fmt.Sprintf("%s?%s", path, form.Encode())
		}
	}

	req, err := h.request("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (h *httpClient) Post(path string, body map[string]interface{}) (responseBs []byte, err error) {
	bodyBs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := h.request("POST", path, strings.NewReader(string(bodyBs)))
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (h *httpClient) request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	for key, value := range h.header {
		req.Header.Add(key, value)
	}

	return req, nil
}
