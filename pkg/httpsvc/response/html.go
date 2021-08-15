package response

import (
	"bytes"
	"fmt"
	"github.com/mingolm/go-recharge/configs"
	"github.com/mingolm/go-recharge/utils/errutil"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	templateSet        = map[string]*template.Template{}
	templateLoadLocker sync.Mutex
)

type htmlResponse struct {
	Filename string
	Data     interface{}
	Cookies  []*http.Cookie
}

func (h *htmlResponse) Headers() (headers map[string]string) {
	return
}
func (h *htmlResponse) AddHeader(key, value string) {
	return
}
func (h *htmlResponse) GetHeader(key string) (value string) {
	return
}
func (h *htmlResponse) WithCookie(cookie *http.Cookie) (ins Response) {
	h.Cookies = append(h.Cookies, cookie)
	return h
}
func (h *htmlResponse) GetCookie() (cookies []*http.Cookie) {
	return h.Cookies
}
func (h *htmlResponse) Bytes() (bs []byte, err error) {
	tmpl, err := h.getTemplate()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, h.Data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (h *htmlResponse) getTemplate() (*template.Template, error) {
	filename := h.Filename
	tmpl, ok := templateSet[filename]
	if ok && tmpl == nil {
		return nil, errutil.ErrNotFound
	}
	if !ok {
		templateLoadLocker.Lock()
		defer func() {
			templateLoadLocker.Unlock()
		}()

		filePath := filename
		if !strings.HasPrefix(filePath, configs.SystemConfig.TemplateHtmlPrefix) {
			if !strings.HasSuffix(filePath, configs.SystemConfig.TemplateBladeType) {
				filePath += fmt.Sprintf(".%s", configs.SystemConfig.TemplateBladeType)
			}
			filePath = filepath.Join(configs.SystemConfig.TemplateHtmlPrefix, filePath)
		}
		bs, err := ioutil.ReadFile(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				templateSet[filename] = nil
				return nil, errutil.ErrUnimplemented
			}
			return nil, fmt.Errorf("read template failed: %w", err)
		}

		tmpl, err = template.New(filename).Parse(string(bs))
		if err != nil {
			return nil, err
		}

		templateSet[filename] = tmpl
	}

	return tmpl, nil
}
