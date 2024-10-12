1. 在使用 fetch API 发送流式请求时，需要指定 duplex 选项。duplex 选项允许你发送流式请求体。你可以将 duplex 设置为 "half" 来解决这个问题。
  - duplex: 'half'：在 fetch 选项中添加 duplex 选项，并将其设置为 "half"，以允许发送流式请求体。
  - Transfer-Encoding: 'chunked'：虽然这不是必须的，但添加这个头部可以明确表示请求体是分块传输的。

2. 浏览器出现 net::ERR_H2_OR_QUIC_REQUIRED 报错

这个错误通常是由于服务器不支持 HTTP/2 或 QUIC 协议，而浏览器尝试使用这些协议发送请求。为了确保浏览器和服务器之间的通信顺畅，可以尝试以下解决方案：
- 确保你的服务器配置支持 HTTP/2。你可以使用一个支持 HTTP/2 的库或框架来启动服务器。 默认情况下，Go 的 http.ListenAndServe 只支持 HTTP/1.1。如果你想让你的 Go 服务器支持 HTTP/2，你需要使用 http2 包来配置服务器。

```golang
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "golang.org/x/net/http2"
)

type logLine struct {
    UserIP string `json:"user_ip"`
    Event  string `json:"event"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
    dec := json.NewDecoder(r.Body)

    for {
        var l logLine
        err := dec.Decode(&l)
        if err == io.EOF {
            break
        }
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Println(l.UserIP, l.Event)
    }
    fmt.Fprintf(w, "OK")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/index.html")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/decode", decodeHandler)
    mux.HandleFunc("/", indexHandler)

    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    // Enable HTTP/2 support
    http2.ConfigureServer(server, &http2.Server{})

    fmt.Println("Server started on :8080")
    if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
```
## Explaination
导入 http2 包：导入 golang.org/x/net/http2 包来配置 HTTP/2 支持。
配置 HTTP/2：使用 http2.ConfigureServer 函数来配置服务器以支持 HTTP/2。

使用 ListenAndServeTLS：为了支持 HTTP/2，服务器需要使用 TLS（HTTPS）。因此，使用 ListenAndServeTLS 方法，并提供证书和密钥文件。

## 生成自签名证书
如果你没有证书和密钥文件，可以使用以下命令生成自签名证书：
`openssl req -x509 -newkey rsa:2048 -nodes -keyout server.key -out server.crt -days 365`



