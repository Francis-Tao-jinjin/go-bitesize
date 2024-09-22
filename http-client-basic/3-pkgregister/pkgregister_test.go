package pkgregister

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func packageRegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Package registration response
		d := pkgregisterResult{}
		err := r.ParseMultipartForm(5000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mForm := r.MultipartForm
		// Get file data
		f := mForm.File["filedata"][0]
		// Construct an artificial package ID to return
		d.ID = fmt.Sprintf("%s-%s", r.FormValue("name"), r.FormValue("version"))
		d.Filename = f.Filename
		d.Size = f.Size
		// Marshal outgoing package registaion response
		jsonData, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(jsonData))
	} else {
		http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
		return
	}
}

func startTestPackageServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(packageRegHandler))
	return ts
}

func TestRegisterPackageData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()

	p := pkgData{
		Name:     "testpkg",
		Version:  "1.0.0",
		Filename: "testpkg.tar.gz",
		Bytes:    strings.NewReader("this is test data"),
	}
	pResult, err := registerPackageData(ts.URL, p)
	fmt.Printf(">>> pResult: %v\n", pResult)
	if err != nil {
		t.Fatal(err)
	}
	if pResult.ID != fmt.Sprintf("%s-%s", p.Name, p.Version) {
		t.Errorf("Expected %q, got %q", fmt.Sprintf("%s-%s", p.Name, p.Version), pResult.ID)
	}
	if pResult.Filename != p.Filename {
		t.Errorf("Expected %q, got %q", p.Filename, pResult.Filename)
	}
	if pResult.Size != 17 {
		t.Errorf("Expected %d, got %d", 17, pResult.Size)
	}
}
