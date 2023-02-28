package itest

import (
	"io"
	"net/http"
	"sync"
)

type HttpClient struct{}

var (
	once     sync.Once
	instance *HttpClient
)

func GetHttpClient() *HttpClient {

	once.Do(func() {
		instance = &HttpClient{}
	})

	return instance
}

func (c *HttpClient) Post(url string, body io.Reader) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
