package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type MockRoundTripper struct {
	Response *http.Response
}

func (c *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return c.Response, nil
}

func TestRoundTrip(t *testing.T) {
	LoginResponse := LoginResponse{
		Token: "123",
	}
	LoginResponseJson, err := json.Marshal(LoginResponse)
	if err != nil {
		t.Fatalf("Marshal error: %s", err)
	}
	myJWTTransport := MyJWTTransport{
		Password: "123",
		Transport: &MockRoundTripper{
			Response: &http.Response{
				StatusCode: 200,
			},
		},
		HTTPClient: &MockClient{
			Response: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(LoginResponseJson)),
			},
		},
	}

	req := &http.Request{
		Header: make(http.Header),
	}
	res, err := myJWTTransport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip error: %s", err)
	}
	if res == nil {
		t.Fatalf("Response is Empty: %s", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Response is not correct. Got: %d", res.StatusCode)
	}
}
