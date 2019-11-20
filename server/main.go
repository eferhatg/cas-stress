package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	cserver := NewCasServer()

	_ = os.Mkdir(cserver.casFolder, os.ModeDir)

	log.Infof("CAS server started on port: %s", cserver.casPort)
	log.Fatal(http.ListenAndServe(":"+cserver.casPort, cserver.BuildRouter()))
}
