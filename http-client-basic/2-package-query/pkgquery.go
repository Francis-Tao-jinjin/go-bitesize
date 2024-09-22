package pkgquery

import (
	"encoding/json"
	"io"
	"net/http"
)

// 带有结构体标签的 struct：提供了更大的灵活性，
// 可以控制字段在序列化和反序列化时的行为。
// 例如，可以指定 JSON 键的名称。
type pkdData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Private bool   `json:"private"`
}

func fetchPackageData(url string) ([]pkdData, error) {
	// create a slice of pkdData to hold the package data
	//  注意，定义的是 slice，所以数据的格式应该要符合 slice 的格式
	var packages []pkdData
	// use a default http client to make a GET request
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// unmarshal the JSON data into the packages slice
	// json.Unmarshal 在解析 JSON 数据时会自动适配不存在的字段。
	// 具体来说，如果 JSON 数据中缺少某些字段，json.Unmarshal 不会报错，
	// 而是会将这些字段保留为其零值（zero value）。同样，如果 JSON
	// 数据中包含结构体中没有定义的字段，这些字段会被忽略。
	// 所以，即使 JSON 数据中没有 "private" 字段，也不会报错。
	err = json.Unmarshal(data, &packages)
	return packages, err
}

// go test -v
