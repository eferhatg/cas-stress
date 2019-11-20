package casutils

import (
	"testing"
	"time"
)

func TestGenerateRandString(t *testing.T) {
	expected := GenerateRandString(10)
	notwant := GenerateRandString(10)
	if expected == notwant {
		t.Errorf("GenerateRandString doesnt work as expected")
	}
}

func TestHashString(t *testing.T) {
	raw := GenerateRandString(10)
	expected := HashString(raw)
	want := HashString(raw)
	if expected != want {
		t.Errorf("HashString doesnt work as expected")
	}
}

func TestRandom(t *testing.T) {

	expected := Random(1, 100000000)
	want := Random(1, 100000000)
	if expected == want {
		t.Errorf("Random doesnt work as expected")
	}
}
func TestTimeTaken(t *testing.T) {
	now := time.Now()
	time.Sleep(2 * time.Second)
	elapsed := TimeTaken(now, "RequestDuration")
	t.Log(elapsed)
	if *elapsed < 2*time.Second {
		t.Errorf("TimeTaken doesnt work as expected")
	}
}
