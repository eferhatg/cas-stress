package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//CasServer keeps the cas server
type CasServer struct {
	casFolder   string
	casRoute    string
	casParam    string
	casEndPoint string
	casPort     string
}

//NewCasServer initialize new cas server
func NewCasServer() *CasServer {
	cserver := CasServer{casFolder: "cas", casRoute: "cas", casParam: "hash"}
	cserver.casEndPoint = fmt.Sprintf("/%s/:%s", cserver.casRoute, cserver.casParam)
	cserver.casPort = os.Getenv("CAS_SERVER_PORT")
	_, err := strconv.Atoi(cserver.casPort)
	if err != nil {
		panic(err)
	}
	return &cserver
}

//BuildRouter builds httprouter
func (cs *CasServer) BuildRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET(cs.casEndPoint, cs.GetContent)
	router.PUT(cs.casEndPoint, cs.PutContent)

	return router
}

//GetContent gets the content
func (cs *CasServer) GetContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	data, err := ioutil.ReadFile(fmt.Sprintf("./%s/%s", cs.casFolder, ps.ByName(cs.casParam)))

	if err == nil {
		_, err := w.Write(data)
		if err != nil {
			log.Error(err)
		}

	} else {
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)

		_, err := w.Write([]byte(http.StatusText(http.StatusNotFound)))
		if err != nil {
			log.Error(err)
		}
	}
}

//PutContent puts the content
func (cs *CasServer) PutContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(""))
		if err != nil {
			log.Error(err)
		}

		file.Close()
		return
	}
	defer file.Close()

	f, err := os.OpenFile(fmt.Sprintf("%s/%s", cs.casFolder, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(""))
		if err != nil {
			log.Error(err)
		}
		f.Close()
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(""))
		if err != nil {
			log.Error(err)
		}
		return
	}
	log.Infof("Got new file with name: %s", handler.Filename)
	w.WriteHeader(200)
	_, err = w.Write([]byte(http.StatusText(http.StatusOK)))
	if err != nil {
		log.Error(err)
	}

}
