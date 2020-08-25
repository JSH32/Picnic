package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Http struct {
}

func (Http) Post(url, content, contenttype string) *Response {
	read := strings.NewReader(content)
	if contenttype == "" {
		contenttype = "application/json"
	}
	res, err := http.Post(url, contenttype, read)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		data = nil
	}

	headers := make(map[string]string)
	for k := range res.Header {
		v := res.Header.Get(k)
		headers[k] = v
	}

	return &Response{
		Status: res.StatusCode,
		Data: Data{
			Text:    string(data),
			Headers: headers,
		},
	}
}

// Get : get request JS bind
func (Http) Get(url string) *Response {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		data = nil
	}

	res.Body.Close()

	headers := make(map[string]string)
	for k := range res.Header {
		v := res.Header.Get(k)
		headers[k] = v
	}

	return &Response{
		Status: res.StatusCode,
		Data: Data{
			Text:    string(data),
			Headers: headers,
		},
	}
}
