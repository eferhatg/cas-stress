// handlers_test.go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/eferhatg/cas-stress/cashttpclient"
	"github.com/julienschmidt/httprouter"
)

func TestGetContent(t *testing.T) {

	cserver := NewCasServer()

	d1 := []byte("hello cas server")
	err := ioutil.WriteFile("cas/test", d1, 0644)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("GET", "/cas/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/cas/:hash", cserver.GetContent)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `hello cas server`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	os.Remove("/cas/test")

}

func TestGetContentError(t *testing.T) {

	cserver := NewCasServer()

	req, err := http.NewRequest("GET", "/cas/nofile", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/cas/:hash", cserver.GetContent)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("GetContent returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPutContent(t *testing.T) {
	cserver := NewCasServer()

	expected := http.StatusText(http.StatusOK)
	d1 := []byte(expected)
	err := ioutil.WriteFile("cas/text.txt", d1, 0644)
	if err != nil {
		t.Fatal(err)
	}
	router := httprouter.New()
	router.PUT("/cas/:hash", cserver.PutContent)

	req, err := cashttpclient.NewMultipartRequest("/cas/text", "file", "cas/text.txt")
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if body := rr.Body; body.String() != expected {
		t.Errorf("GetContent returned wrong status code: got %v want %v",
			body, expected)
	}
	os.Remove("cas/text.txt")

}
