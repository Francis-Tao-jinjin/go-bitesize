package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type LoggingClient struct {
	log *log.Logger
}

func (c LoggingClient) RoundTrip(r *http.Request) (*http.Response, error) {
	c.log.Printf("Request: %s %s\n", r.Method, r.URL.String())
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		c.log.Printf("Encounter an error: %s\n", err)
		return nil, err
	}
	c.log.Printf("Got back a response over: %s\n", resp.Proto)
	return resp, nil
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func crerateHTTPClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		Timeout: d,
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}
	myTransport := LoggingClient{
		log: log.New(os.Stdout, "[@ Logger @]:", 0),
	}
	client := crerateHTTPClientWithTimeout(5 * time.Second)
	client.Transport = myTransport
	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error fetching data: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Data: %s\nBytes in response: %d\n", string(body), len(body))
}

// go run . https://ifconfig.me/all.json
