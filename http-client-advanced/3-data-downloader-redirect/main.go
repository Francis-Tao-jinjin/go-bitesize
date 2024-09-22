package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	if len(via) >= 1 {
		// return an error to stop redirects
		return errors.New(fmt.Sprintf("Attempted redirect to %s", req.URL))
	}
	// return nil to allow redirects
	return nil
}

func crerateHTTPClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		CheckRedirect: redirectPolicyFunc,
		Timeout:       d,
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}
	client := crerateHTTPClientWithTimeout(5 * time.Second)
	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error fetching data: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Data: %s\n", string(body))
}

// go run . https://ifconfig.me/all.json
