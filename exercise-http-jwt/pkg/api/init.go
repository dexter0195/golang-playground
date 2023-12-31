package api

import (
	"io"
	"net/http"
)

type ClientInterface interface {
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type Options struct {
	Password string
	LoginUrl string
}

type APIIface interface {
	DoRequest(requestURL string) (Response, error)
}

type API struct {
	Options Options
	Client  ClientInterface
}

func New(options Options) APIIface {
	return API{
		Options: options,
		Client: &http.Client{
			Transport: &MyJWTTransport{
				Transport:  http.DefaultTransport,
				Password:   options.Password,
				LoginUrl:   options.LoginUrl,
				HTTPClient: &http.Client{},
			},
		},
	}
}
