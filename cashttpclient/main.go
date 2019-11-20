package cashttpclient

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// NewMultipartRequest creates new multipart request
func NewMultipartRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// RequestExecuter executes given request
func RequestExecuter(request *http.Request) (*http.Response, *bytes.Buffer, error) {

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	resp.Body.Close()

	return resp, body, nil

}
