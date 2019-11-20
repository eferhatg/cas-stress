package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestEmptyFolder(t *testing.T) {
	d1 := []byte("hello cas  client")
	err := ioutil.WriteFile("temp/text.txt", d1, 0644)
	if err != nil {
		t.Fatal(err)
	}
	path := "temp/"
	EmptyTempFolder(path)
	directory := path
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)
	if len(dirFiles) > 0 {
		t.Errorf("EmptyFolder doesnt work as expected")
	}
}
