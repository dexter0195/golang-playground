package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/dexter0195/golang-playground/exercise-http-jwt/pkg/api"
)

func main() {
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)
	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "use a password to access our api")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Help: ./http-get -h\nURL is not valid URL: %s\n", requestURL)
		os.Exit(1)
	}

	client := http.Client{}

	api.New()

	if password != "" {
		token, err := doLoginRequest(client, parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			if requestErr, ok := err.(RequestError); ok {
				fmt.Printf("Error occurred: %s (HTTP Error: %d, Body: %s)\n", requestErr.Error(), requestErr.HTTPCode, requestErr.Body)
				os.Exit(1)
			}
			fmt.Printf("Error occurred: %s\n", err)
			os.Exit(1)
		}
		client.Transport = &MyJWTTransport{
			Transport: http.DefaultTransport,
			Token:     token,
		}
	}

	res, err := doRequest(client, parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error occurred: %s (HTTP Error: %d, Body: %s)\n", requestErr.Error(), requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error occurred: %s\n", err)
		os.Exit(1)
	}
	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}
	fmt.Printf("Response: %s\n", res.GetResponse())
}
