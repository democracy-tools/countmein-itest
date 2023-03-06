package itest

import (
	"fmt"
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

func (c *HttpClient) Post(uri string, body io.Reader) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, getUrl(uri), body)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func (c *HttpClient) Get(uri string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, getUrl(uri), nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func getUrl(uri string) string {

	return fmt.Sprintf("%s/%s", GetEnvOrExit("BASE_URL"), uri)
}
