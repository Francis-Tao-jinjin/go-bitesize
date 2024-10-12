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

// data = await fetch('/decode', {
// 	method: "POST",
// 	body: JSON.stringify({user_ip: 'test ip', event: 'test event'})
// })

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
	// http.ListenAndServe(":8080", mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Enable http2 support
	http2.ConfigureServer(server, &http2.Server{})
	fmt.Println("Server started on :8080")
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}

// Generate self-signed certificate
// openssl req -x509 -newkey rsa:2048 -nodes -keyout server.key -out server.crt -days 365
