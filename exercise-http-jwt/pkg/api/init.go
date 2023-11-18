package api

import (
	"net/http"
)

type Options struct {
	Password string
	Login    string
}

type APIIface interface {
}

type API struct {
	Options Options
	Client  http.Client
}

func New() APIIface {
	return API{
		Options: options,
		Client: http.Client{
			Transport: &MyJWTTransport{
				transport: http.DefaultTransport,
			},
		},
	}
}
