# Middleware Chain HTTP Client in Go

This project demonstrates how to implement an HTTP client in Go with a middleware chain. The middleware chain allows you to intercept and modify HTTP requests and responses by adding custom middlewares. This implementation is particularly useful for tasks like logging, adding headers, and more.

## Features

- Support for multiple middlewares in a chain.
- Logging middleware to log requests and responses.
- Header middleware to add custom headers to HTTP requests.
- Easily extendable to add more custom middlewares.

## Code Structure

- Middleware interfaces and implementations:
  - `Middleware`: Interface for middlewares.
  - `MiddlewareFunc`: Helper type to create middlewares from functions.
  - `LoggingMiddleware`: Middleware for logging requests and responses.
  - `HeaderMiddleware`: Middleware for adding custom headers.
- `Chain`: Middleware chain to manage and execute middlewares.
- Utility functions:
  - `fetchRemoteResource`: Fetches data from a specified URL.
  - `createHTTPClientWithTimeout`: Creates an HTTP client with a specified timeout.

### Try running the Application

```sh
go run . https://ifconfig.me/all.json
```

### Adding Middlewares

You can add middlewares to the chain by using the `Add` method of `Chain`. Here is an example of how to add logging and header middlewares:

```go
chain := &Chain{}
chain.Add(LoggingMiddleware{
	log: log.New(os.Stdout, "[Logger]:", 0),
})
chain.Add(HeaderMiddleware{
	headers: map[string]string{
		"User-Agent":     "Middleware-Chain-Client",
		"X-Custom-Header": "Custom-Value",
	},
})
```

### Code Overview

#### Middleware Interfaces and Implementations

- **Middleware Interface:**
    ```go
    type Middleware interface {
        RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error)
    }
    ```

- **LoggingMiddleware:**
    ```go
    type LoggingMiddleware struct {
        log *log.Logger
    }
    
    func (m LoggingMiddleware) RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error) {
        m.log.Printf("Request: %s %s\n", r.Method, r.URL.String())
        resp, err := next.RoundTrip(r)
        if err != nil {
            m.log.Printf("Encounter an error: %s\n", err)
            return nil, err
        }
        m.log.Printf("Got back a response over: %s\n", resp.Proto)
        return resp, nil
    }
    ```

- **HeaderMiddleware:**
    ```go
    type HeaderMiddleware struct {
        headers map[string]string
    }
    
    func (m HeaderMiddleware) RoundTrip(r *http.Request, next http.RoundTripper) (*http.Response, error) {
        for key, value := range m.headers {
            r.Header.Set(key, value)
        }
        return next.RoundTrip(r)
    }
    ```

- **Adding and Executing Middlewares:**
    ```go
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
            currentNext := final
    
            final = &middlewareRoundTripper{
                middleware: m,
                next:       currentNext,
            }
        }
        return final.RoundTrip(r)
    }
    ```

#### Utility Functions

- **fetchRemoteResource:**
    ```go
    func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
        response, err := client.Get(url)
        if err != nil {
            return nil, err
        }
        defer response.Body.Close()
        return io.ReadAll(response.Body)
    }
    ```

- **createHTTPClientWithTimeout:**
    ```go
    func createHTTPClientWithTimeout(d time.Duration) *http.Client {
        return &http.Client{
            Timeout: d,
        }
    }
    ```

### Example

Here's an example of how to use the middleware chain with logging and header middlewares:

```go
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
```

## Further Improvements

- Adding more middlewares like caching, retry logic, authentication, etc.
- Improving error handling and logging.
- Configuring middlewares and clients through configuration files or environment variables.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

This README should provide a comprehensive overview of your project, its usage, and how to extend it with custom middlewares. If you have any further questions or need additional sections, feel free to ask!