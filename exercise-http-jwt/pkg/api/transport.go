package api

import "net/http"

type MyJWTTransport struct {
	Transport  http.RoundTripper
	Token      string
	Password   string
	LoginUrl   string
	HTTPClient ClientInterface
}

func (t *MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Token == "" {
		if t.Password != "" {
			token, err := doLoginRequest(t.HTTPClient, Options{
				Password: t.Password,
				LoginUrl: t.LoginUrl,
			})
			if err != nil {
				return nil, err
			}
			t.Token = token
		}
		if t.Token != "" {
			req.Header.Add("Authorization", "Bearer "+t.Token)
		}
	}
	return t.Transport.RoundTrip(req)
}
