package main

import "net/http"

type MyJWTTransport struct {
	Transport http.RoundTripper
	Token     string
}

func (t *MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Token != "" {
		req.Header.Add("Authorization", "Bearer "+t.Token)
	}
	return t.Transport.RoundTrip(req)
}
