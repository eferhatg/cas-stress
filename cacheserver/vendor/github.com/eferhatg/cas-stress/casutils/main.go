package casutils

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//GenerateRandString generates n size random string
func GenerateRandString(n int64) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//HashString hashes string based on SHA1
func HashString(s string) string {
	hash := sha1.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		log.Error(err)
	}
	hashInBytes := hash.Sum(nil)[:20]
	return hex.EncodeToString(hashInBytes)
}

//Random generates random number between min and max
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//TimeTaken logs time since T timef
func TimeTaken(t time.Time, name string) *time.Duration {
	elapsed := time.Since(t)
	log.Infof("TIME: %s took %s\n", name, elapsed)
	return &elapsed
}
