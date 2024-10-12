package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	log.Fatal(http.ListenAndServe(
		listenAddr,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World from GO!"))
		}),
	))
}

// runs the server on port 9090
// LISTEN_ADDR=:9090 go run main.go
