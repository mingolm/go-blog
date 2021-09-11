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

var templateCommonLoadOnce sync.Once
var templateCommonPaths []string

func (h *htmlResponse) getTemplate() (tmpl *template.Template, err error) {
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

		templateCommonLoadOnce.Do(func() {
			commonDir := configs.SystemConfig.TemplateHtmlCommon
			if commonDir != "" {
				files, err := ioutil.ReadDir(commonDir)
				if err != nil {
					panic(fmt.Sprintf("template common dir load failed. [dir=%s, err=%s]", commonDir, err.Error()))
				}
				for _, f := range files {
					templateCommonPaths = append(templateCommonPaths, fmt.Sprintf("%s%s", commonDir, f.Name()))
				}
			}
		})

		filePath := filename
		if !strings.HasPrefix(filePath, configs.SystemConfig.TemplateHtmlPrefix) {
			if !strings.HasSuffix(filePath, configs.SystemConfig.TemplateBladeType) {
				filePath += fmt.Sprintf(".%s", configs.SystemConfig.TemplateBladeType)
			}
			filePath = filepath.Join(configs.SystemConfig.TemplateHtmlPrefix, filePath)
		}
		loadPaths := []string{filePath}
		if templateCommonPaths != nil {
			loadPaths = append(loadPaths, templateCommonPaths...)
		}
		tmpl, err = template.ParseFiles(loadPaths...)
		if err != nil {
			if os.IsNotExist(err) {
				templateSet[filename] = nil
				return nil, errutil.ErrNotFound
			}
			return nil, fmt.Errorf("template parse file failed. [err=%s]", err.Error())
		}
		templateSet[filename] = tmpl
	}

	return tmpl, nil
}
