package pkgregister

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

func creareMultiPartMessage(data pkgData) ([]byte, string, error) {
	var b bytes.Buffer
	var err error
	var fw io.Writer

	mw := multipart.NewWriter(&b)

	// CreateFormField returns an io.Writer, which is used to write the field value
	fw, err = mw.CreateFormField("name")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprintf(fw, data.Name)

	fw, err = mw.CreateFormField("version")
	if err != nil {
		return nil, "", err
	}
	fmt.Fprintf(fw, data.Version)

	fw, err = mw.CreateFormFile("filedata", data.Filename)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(fw, data.Bytes)
	if err != nil {
		return nil, "", err
	}
	// Close the multipart writer !!!
	err = mw.Close()
	if err != nil {
		return nil, "", err
	}
	contentType := mw.FormDataContentType()
	return b.Bytes(), contentType, nil
}
