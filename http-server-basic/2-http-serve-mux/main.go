package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World from GO!")
	// or
	w.Write([]byte("Hello World from GO!"))
	fmt.Println("Request URL:", r.URL.Path)
}

func healthChechHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am alive!")
}

func setupHandlers(mux *http.ServeMux) {
	// mux.HandleFunc("/api/", apiHandler)
	// 如果你想要匹配所有以 /api 开头的请求，就可以使用 /api/，这样就可以匹配 /api/，/api/v1，/api/v1/users 等等
	// 否则，只使用 /api 的时候，就只能匹配 /api 这个路径，访问 /api/v1 就会返回 404
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthz", healthChechHandler)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
