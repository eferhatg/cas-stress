package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eferhatg/cas-stress/cashttpclient"
	"github.com/eferhatg/cas-stress/casutils"
	"github.com/golang/groupcache"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

//CacheServer keeps Cache details
type CacheServer struct {
	gcache        *groupcache.Group
	hitMissRatio  float64
	casParam      string
	casServerAddr string
	casServerPath string
	ownAddr       string
	peerPort      string
	listenPort    string
}

//NewCacheServer inits Cache object
func NewCacheServer() *CacheServer {

	casAddr := os.Getenv("CAS_SERVER_ADDR")
	peerPort := os.Getenv("PEER_PORT")
	listenPort := os.Getenv("LISTEN_PORT")

	return &CacheServer{
		casParam:      "hash",
		ownAddr:       "http://localhost:" + peerPort,
		casServerPath: "cas",
		casServerAddr: casAddr,
		peerPort:      peerPort,
		listenPort:    listenPort,
		hitMissRatio:  0.0,
	}
}

// GetContent handles GET requests
func (c *CacheServer) GetContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer casutils.TimeTaken(time.Now(), "RequestDuration")

	randomNum := casutils.Random(0, 100)

	if s, err := strconv.ParseFloat(r.Header.Get("HitMissRatio"), 64); err == nil {
		c.hitMissRatio = s
	}

	key := ps.ByName(c.casParam)
	if c.gcache.Stats.Loads >= 1 && int(c.hitMissRatio*100) < randomNum {
		key = key + "-" + casutils.GenerateRandString(10)
		log.Info("Will miss due to HitMissRatio")
	}

	var ctx groupcache.Context
	var data []byte

	if err := c.gcache.Get(ctx, key, groupcache.AllocatingByteSliceSink(&data)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(http.StatusText(http.StatusNotFound) + err.Error()))
		if err != nil {
			log.Error(err)
		}
		return
	}

	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
	}
	c.showStatus()
}

//CacheServerGetter gets and cache the real content if not exists
func (c *CacheServer) CacheServerGetter(ctx groupcache.Context, key string, dest groupcache.Sink) error {

	key = strings.Split(key, "-")[0]
	endPoint := fmt.Sprintf("http://%s/%s/%s", c.casServerAddr, c.casServerPath, key)

	request, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, body, err := cashttpclient.RequestExecuter(request)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = dest.SetBytes(body.Bytes())
	if err != nil {
		log.Error(err)
	}
	return nil
}

func (c *CacheServer) showStatus() []byte {
	metrics := c.gcache.Stats

	val, _ := json.Marshal(metrics)
	log.Info(string(val))
	return val
}

//StartGroupCache starts group cache and listens requests
func (c *CacheServer) StartGroupCache(groupname string) {
	peers := groupcache.NewHTTPPool(c.ownAddr)
	peers.Set(c.ownAddr)

	c.gcache = groupcache.NewGroup(groupname, 64<<20, groupcache.GetterFunc(c.CacheServerGetter))
	go http.ListenAndServe(":"+c.peerPort, http.HandlerFunc(peers.ServeHTTP))
	log.Info("Groupcache started on port " + c.peerPort)
}
