package pkgregister

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type pkgData struct {
	Name     string
	Version  string
	Filename string
	Bytes    io.Reader
}

type pkgregisterResult struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

func registerPackageData(url string, data pkgData) (pkgregisterResult, error) {
	p := pkgregisterResult{}
	payload, contentType, err := creareMultiPartMessage(data)
	// fmt.Printf(">>> creareMultiPartMessage payload: %v\n", payload)
	if err != nil {
		return p, err
	}
	reader := bytes.NewReader(payload)
	r, err := http.Post(url, contentType, reader)
	fmt.Printf(">>> http.Post response: %v\n", r)
	if err != nil {
		return p, err
	}
	defer r.Body.Close()
	respData, err := io.ReadAll(r.Body)

	if r.StatusCode != http.StatusOK {
		return p, errors.New(string(respData)) //fmt.Errorf("unexpected status code: %d", r.StatusCode)
	}
	if r.Header.Get("Content-Type") != "application/json" {
		return p, nil
	}
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(respData, &p)
	return p, err
}

// if the data is Json format, use this function
func registerJsonData(url string, data pkgData) (pkgregisterResult, error) {
	p := pkgregisterResult{}
	b, err := json.Marshal(data)
	if err != nil {
		return p, err
	}
	reader := bytes.NewReader(b)
	r, err := http.Post(url, "application/json", reader)
	if err != nil {
		return p, err
	}
	defer r.Body.Close()
	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return p, err
	}
	if r.StatusCode != http.StatusOK {
		return p, errors.New(string(respData))
	}
	err = json.Unmarshal(respData, &p)
	return p, err
}
