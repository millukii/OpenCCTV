package client

import "net/http"

var DefaultClient = &http.Client{}

type HTTPClient interface {
}

type httpClient struct {
}

func NewHTTPClient() HTTPClient {

	return &httpClient{}
}
