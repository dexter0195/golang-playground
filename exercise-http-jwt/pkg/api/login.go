package api

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
)

type LoginRequest struct {
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client ClientInterface, options Options) (string, error) {
	loginRequest := LoginRequest{
		Password: options.Password,
	}

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("marshal error: %s", err)
	}

	res, err := client.Post(options.LoginUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("post error: %s", err)
	}

	if err != nil {
		return "", fmt.Errorf("get error: %s", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("readAll error: %s", err)
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("invalid output (HTTP Code %d): %s", res.StatusCode, string(body))
	}

	var loginResponse LoginResponse

	if !json.Valid(resBody) {
		return "", RequestError{
			Err:      "response is not a json",
			HTTPCode: res.StatusCode,
			Body:     string(body),
		}
	}

	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		return "", RequestError{
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
			HTTPCode: res.StatusCode,
			Body:     string(body),
		}
	}
	return loginResponse.Token, nil
}
