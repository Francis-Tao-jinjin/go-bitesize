package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_DecodeHandler(t *testing.T) {
	const jsonStream = `
	{"user_ip": "172.121.19.21", "event": "click_on_add_cart"}{"user_ip": "172.121.19.21", "event": "click_on_checkout"}
	`
	body := strings.NewReader(jsonStream)
	r := httptest.NewRequest("POST", "http://localhost/decode", body)
	w := httptest.NewRecorder()
	decodeHandler(w, r)
}

// go test -v
// go test -v -run Test_DecodeHandler
