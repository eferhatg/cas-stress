package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStartStressTest(t *testing.T) {
	var counter uint64

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		counter++
		if req.Header.Get("HitMissRatio") == "" {
			t.Fatalf("StartStressTest doesnt acquired header as expected")
		}
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte("body"))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer func() { testServer.Close() }()

	ccli := CasClient{casRoute: "cas", contentSize: 2500, attackDuration: 4, attackFreq: 2, cacheEndPoint: testServer.URL}
	metrics := ccli.StartStressTest()

	if counter != metrics.Requests {
		t.Errorf("StartStressTest doesnt requests as expected, expected %v, got %v", counter, metrics.Requests)
	}
}

func TestGenerateFile(t *testing.T) {
	ccli := CasClient{contentSize: 2500, tempDir: "temp"}

	fileName := ccli.GenerateFile()

	if ccli.fileName == "" {
		t.Errorf("GenerateFile doesnt set filename")
	}
	file, err := os.Open(ccli.tempDir + "/" + fileName)
	if err != nil {
		t.Errorf("GenerateFile couldn't read file")
	}
	defer file.Close()
	finfo, err := file.Stat()
	if err != nil {
		t.Errorf("GenerateFile couldn't get file size")
	}
	if ccli.contentSize != finfo.Size() {
		t.Errorf("GenerateFile content size is not as expected, expected %v, got %v", ccli.contentSize, finfo.Size())
	}

}
