package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {

	c := NewCacheServer()

	c.StartGroupCache("cas_server")

	router := httprouter.New()
	router.GET(fmt.Sprintf("/%s/:%s", c.casServerPath, c.casParam), c.GetContent)

	log.Infof("listen port: %v", ":"+c.listenPort)
	log.Fatal(http.ListenAndServe(":"+c.listenPort, router))

}
