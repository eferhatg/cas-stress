package cashttpclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestNewMultipartRequest(t *testing.T) {

	remotePath := "http://exampleurl:exampleport/cas/filename"

	d1 := []byte("hello cas http client")
	err := ioutil.WriteFile("text.txt", d1, 0644)
	if err != nil {
		t.Error(err)
	}

	request, err := NewMultipartRequest(remotePath, "file", "text.txt")
	if err != nil {
		t.Errorf("NewMultipartRequest returned error %v", err)
	}
	if request.Method != "PUT" {
		t.Errorf("NewMultipartRequest returned wrong method %s", request.Method)
	}
	if request.URL.String() != remotePath {
		t.Errorf("NewMultipartRequest returned wrong URL %s", request.URL)
	}
	if !strings.Contains(request.Header.Get("Content-Type"), "multipart/form-data") {
		t.Errorf("NewMultipartRequest returned wrong Header %s", request.Header.Get("Content-Type"))
	}
	os.Remove("text.txt")
}

func TestRequestExecuter(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte("body"))
		if err != nil {
			t.Error(err)
		}
	}))
	defer func() { testServer.Close() }()

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Errorf("Error when creating request %v", err)
	}

	res, _, err := RequestExecuter(req)
	if err != nil {
		t.Errorf("Error when request executing %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("RequestExecuter returned wrong response %v", res.StatusCode)
	}

}
