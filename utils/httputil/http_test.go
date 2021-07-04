package httputil

import "testing"

func TestSend(t *testing.T) {
	cli := NewHTTPClient(&HTTPClientConfig{})
	_, err := cli.AddHeader(nil).Get("http://baidu.com?aa=11&vv=22", map[string]string{
		"a": "1",
		"b": "2",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	cli2 := NewHTTPClient(&HTTPClientConfig{})
	_, err = cli2.PostForm("http://baidu.com", map[string]string{
		"a": "1",
		"b": "2",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	cli3 := NewHTTPClient(&HTTPClientConfig{})
	_, err = cli3.PostJson("http://baidu.com", map[string]interface{}{
		"a": "1",
		"b": "2",
		"c": 3,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
