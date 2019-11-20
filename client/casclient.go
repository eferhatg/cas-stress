package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/eferhatg/cas-stress/cashttpclient"
	"github.com/eferhatg/cas-stress/casutils"
	log "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

//CasClient keeps the CAS Client
type CasClient struct {
	casRoute       string
	tempDir        string
	serverEndPoint string
	cacheEndPoint  string
	attackFreq     int
	attackDuration int64
	contentSize    int64
	fileName       string
	hitMissRatio   float64
}

//NewCasClient inits cas client
func NewCasClient() *CasClient {
	ccli := CasClient{casRoute: "cas", tempDir: "temp", contentSize: 2500, attackDuration: 4, attackFreq: 2, hitMissRatio: 0.0}

	if cz, err := strconv.ParseInt(os.Getenv("CONTENT_SIZE"), 10, 64); err == nil {
		ccli.contentSize = cz
	}
	if af, err := strconv.Atoi(os.Getenv("ATTACK_FREQ")); err == nil {
		ccli.attackFreq = af
	}
	if ad, err := strconv.ParseInt(os.Getenv("ATTACK_DURATION"), 10, 64); err == nil {
		ccli.attackDuration = ad
	}

	if hm, err := strconv.ParseFloat(os.Getenv("HIT_MISS_RATIO"), 64); err == nil {
		ccli.hitMissRatio = hm
	}

	return &ccli

}

//StartStressTest tests the cache server
func (cc *CasClient) StartStressTest() *vegeta.Metrics {

	rate := vegeta.Rate{Freq: cc.attackFreq, Per: time.Second}
	duration := time.Duration(cc.attackDuration) * time.Second

	headers := http.Header{}
	headers.Add("HitMissRatio", fmt.Sprintf("%.3f", cc.hitMissRatio))

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    cc.cacheEndPoint,
		Header: headers,
	})

	attacker := vegeta.NewAttacker()

	log.Infof("Attacking cache server with the hit miss ratio of %6.2f", cc.hitMissRatio)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "CAS Attack") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics

}

//UploadFile uploads file to cas server
func (cc *CasClient) UploadFile() (*http.Response, *bytes.Buffer, error) {
	request, err := cashttpclient.NewMultipartRequest(cc.serverEndPoint, "file", fmt.Sprintf("%s/%s", cc.tempDir, cc.fileName))
	if err != nil {
		log.Fatal(err)
	}
	return cashttpclient.RequestExecuter(request)
}

//GenerateFile generates random file
func (cc *CasClient) GenerateFile() string {

	serverAddr := os.Getenv("CAS_SERVER_ADDR")
	cacheAddr := os.Getenv("CACHE_SERVER_ADDR")

	randStr := casutils.GenerateRandString(cc.contentSize)
	fileName := casutils.HashString(randStr)

	localPath := fmt.Sprintf("%s/%s", cc.tempDir, fileName)

	err := ioutil.WriteFile(localPath, []byte(randStr), 0644)
	if err != nil {
		panic(err)
	}
	cc.fileName = fileName
	cc.serverEndPoint = fmt.Sprintf("http://%s/%s/%s", serverAddr, cc.casRoute, cc.fileName)
	cc.cacheEndPoint = fmt.Sprintf("http://%s/%s/%s", cacheAddr, cc.casRoute, cc.fileName)
	return fileName
}
