package main

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	ccli := NewCasClient()

	EmptyTempFolder(ccli.tempDir)

	ccli.GenerateFile()

	_, _, err := ccli.UploadFile()
	if err != nil {
		log.Error(err)
	}
	metrics := ccli.StartStressTest()

	val, _ := json.Marshal(metrics)
	log.Info(string(val))
}

//EmptyTempFolder empties the folder
func EmptyTempFolder(path string) {
	directory := path
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)

	for index := range dirFiles {
		fileHere := dirFiles[index]

		nameHere := fileHere.Name()
		fullPath := directory + nameHere

		os.Remove(fullPath)
	}
}
