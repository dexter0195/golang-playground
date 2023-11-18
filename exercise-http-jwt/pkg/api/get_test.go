package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	Response *http.Response
}

func (c *MockClient) Get(url string) (*http.Response, error) {
	return c.Response, nil
}

func TestDoGetRequest(t *testing.T) {

	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "abc",
			Words: []string{"a", "b"},
		},
	}

	wordsJson, err := json.Marshal(words)
	if err != nil {
		t.Fatalf("Marshal error: %s", err)
	}

	apiInstance := API{
		Options: Options{},
		Client: &MockClient{
			Response: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsJson)),
			},
		},
	}

	res, err := apiInstance.DoRequest("http://localhost:8080/words")

	if err != nil {
		t.Errorf("DoRequest error: %s", err)
	}

	if res == nil {
		t.Fatalf("Response is Empty: %s", err)
	}

	if res.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Fatalf("Response is not correct. Response is: %s", res.GetResponse())
	}
}
