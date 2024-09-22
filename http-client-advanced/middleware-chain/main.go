package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Middleware defines the interface for HTTP middlewares
type Middleware interface {
	RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error)
}

// MiddlewareFunc is a helper type to create middlewares from functions
type MiddlewareFunc func(r *http.Request, next http.RoundTripper) (*http.Response, error)

func (f MiddlewareFunc) RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error) {
	return f(r, next)
}

// LoggingMiddleware logs the request and response
type LoggingMiddleware struct {
	log *log.Logger
}

func (m LoggingMiddleware) RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error) {
	m.log.Printf("Request: %s %s\n", r.Method, r.URL.String())
	// resp, err := http.DefaultTransport.RoundTrip(r)
	resp, err := next.RoundTrip(r)
	if err != nil {
		m.log.Printf("Encounter an error: %s\n", err)
		return nil, err
	}
	m.log.Printf("Got back a response over: %s\n", resp.Proto)
	return resp, nil
}

// HeaderMiddleware adds custom headers to the request
type HeaderMiddleware struct {
	headers map[string]string
}

func (m HeaderMiddleware) RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error) {
	for key, value := range m.headers {
		r.Header.Set(key, value)
	}
	// return http.DefaultTransport.RoundTrip(r)
	return next.RoundTrip(r)
}

type middlewareRoundTripper struct {
	middleware Middleware
	next       http.RoundTripper
}

func (mrt *middlewareRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return mrt.middleware.RoundTrip(r, mrt.next)
}

// Chain is a middleware chain
type Chain struct {
	middlewares []Middleware
}

func (c *Chain) Add(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

func (c *Chain) RoundTrip(r *http.Request) (*http.Response, error) {
	var final http.RoundTripper = http.DefaultTransport
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		m := c.middlewares[i]
		currentNext := final // Capture the current next

		final = &middlewareRoundTripper{
			middleware: m,
			next:       currentNext,
		}
	}
	return final.RoundTrip(r)
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{
		Timeout: d,
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from\n")
		os.Exit(1)
	}

	// Create a middleware chain
	chain := &Chain{}
	chain.Add(LoggingMiddleware{
		log: log.New(os.Stdout, "[1st Logger @]:", 0),
	})
	chain.Add(LoggingMiddleware{
		log: log.New(os.Stdout, "[2nd Logger #]:", 0),
	})
	// chain.Add(HeaderMiddleware{
	// 	headers: map[string]string{
	// 		"User-Agent":     "Middleware-Chain-Client",
	// 		"X-Cutom-Header": "Custom-Value",
	// 	},
	// })

	client := &http.Client{
		Transport: chain,
	}

	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error fetching data: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Data: %s\nBytes in response: %d\n", string(body), len(body))
}
