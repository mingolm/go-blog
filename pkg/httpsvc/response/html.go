package response

import (
	"bytes"
	"fmt"
	"github.com/mingolm/go-recharge/utils/errutil"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	templateBasePath = "resources/html/"
)

var (
	templateSet        = map[string]*template.Template{}
	templateLoadLocker sync.Mutex
)

type htmlResponse struct {
	Filename string
	Data     interface{} `json:"data"`
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
	tmpl, ok := templateSet[h.Filename]
	if ok && tmpl == nil {
		return nil, errutil.ErrNotFound
	}
	if !ok {
		templateLoadLocker.Lock()
		defer func() {
			templateLoadLocker.Unlock()
		}()

		bs, err := ioutil.ReadFile(filepath.Join(templateBasePath, h.Filename+".html"))
		if err != nil {
			if os.IsNotExist(err) {
				templateSet[h.Filename] = nil
				return nil, errutil.ErrUnimplemented
			}
			return nil, fmt.Errorf("read template failed: %w", err)
		}

		tmpl, err = template.New(h.Filename).Parse(string(bs))
		if err != nil {
			return nil, err
		}

		templateSet[h.Filename] = tmpl
	}

	return tmpl, nil
}
