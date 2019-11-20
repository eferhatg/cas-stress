// handlers_test.go
package main

import (
	"os"
	"testing"

	"github.com/golang/groupcache"
)

func TestNewCacheServer(t *testing.T) {
	expCasServer := "localhost:5000"
	expPeerPort := "8001"
	expListenPort := "18001"

	os.Setenv("CAS_SERVER_ADDR", expCasServer)
	os.Setenv("PEER_PORT", expPeerPort)
	os.Setenv("LISTEN_PORT", expListenPort)
	cs := NewCacheServer()
	expOwnAddr := "http://localhost:" + cs.peerPort
	if cs.casServerAddr != expCasServer {
		t.Errorf("NewCacheServer returned CasServer : got %v want %v",
			cs.casServerAddr, expCasServer)
	}
	if cs.peerPort != expPeerPort {
		t.Errorf("NewCacheServer returned PeerPort : got %v want %v",
			cs.peerPort, expPeerPort)
	}
	if cs.listenPort != expListenPort {
		t.Errorf("NewCacheServer returned ListenPort : got %v want %v",
			cs.listenPort, expListenPort)
	}
	if cs.ownAddr != expOwnAddr {
		t.Errorf("NewCacheServer returned ListenPort : got %v want %v",
			cs.ownAddr, expOwnAddr)
	}

}

func TestStartGroupCache(t *testing.T) {
	expCasServer := "localhost:5000"
	expPeerPort := "8001"
	expListenPort := "18001"

	os.Setenv("CAS_SERVER_ADDR", expCasServer)
	os.Setenv("PEER_PORT", expPeerPort)
	os.Setenv("LISTEN_PORT", expListenPort)
	cs := NewCacheServer()

	cs.StartGroupCache("cas_test")

	if cs.gcache == nil {
		t.Errorf("StartGroupCache couldn't set groupcache object ")
	}

}
func TestShowStatus(t *testing.T) {
	expCasServer := "localhost:5000"
	expPeerPort := "8001"
	expListenPort := "18001"

	os.Setenv("CAS_SERVER_ADDR", expCasServer)
	os.Setenv("PEER_PORT", expPeerPort)
	os.Setenv("LISTEN_PORT", expListenPort)
	cs := NewCacheServer()
	cs.gcache = groupcache.NewGroup("cas_test_2", 64<<20, groupcache.GetterFunc(cs.CacheServerGetter))

	cs.gcache.Stats = groupcache.Stats{Gets: 6}

	val := cs.showStatus()
	expected := "{\"Gets\":6,\"CacheHits\":0,\"PeerLoads\":0,\"PeerErrors\":0,\"Loads\":0,\"LoadsDeduped\":0,\"LocalLoads\":0,\"LocalLoadErrs\":0,\"ServerRequests\":0}"
	if string(val) != expected {
		t.Errorf("showStatus returned wrong value : got %v want %v ",
			string(val), expected)

	}

}
