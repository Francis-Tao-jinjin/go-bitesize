package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

type requestContextKey struct{}
type requestContextValue struct {
	requestID string
}

func addRequestID(r *http.Request, requestID string) *http.Request {
	c := requestContextValue{
		requestID: requestID,
	}
	currentCtx := r.Context()
	// 使用 空结构体 requestContextKey 最为 key，将 requestContextValue 作为 value 存入 context
	// 空结构体和字符串相比，空结构体更加节省内存，因为字符串是一个指针，而空结构体是一个零长度的结构体
	// 而且，空结构体作为 key 是唯一的，不会和其他 key 冲突。可以理解成 js 中的 Symbol
	newCtx := context.WithValue(currentCtx, requestContextKey{}, c)
	return r.WithContext(newCtx)
}

func logRequest(r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(requestContextKey{})
	if m, ok := v.(requestContextValue); ok {
		log.Printf("Processing request: %s", m.requestID)
	}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Write([]byte("Hello World from GO!"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	requestID := "request-123-abc"
	r = addRequestID(r, requestID)
	processRequest(w, r)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api", apiHandler)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
