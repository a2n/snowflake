package snowflake

import (
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	t.Log(Uint64())
}

// go test -v -run ^TestDataRacing$ -cpu 4 -race
func TestDataRacing(t *testing.T) {
	var rwMux sync.RWMutex
	m := map[uint64]struct{}{}
	put := func() {
		n := Uint64()
		rwMux.RLock()
		_, ok := m[n]
		rwMux.RUnlock()
		if ok == true {
			panic("found")
		}

		rwMux.Lock()
		m[n] = struct{}{}
		rwMux.Unlock()
	}

	sem := make(chan struct{}, 1024)
	for i := 0; i < 102400; i++ {
		sem <- struct{}{}
		go func() {
			put()
			<-sem
		}()
	}

	for i := 0; i < 1024; i++ {
		sem <- struct{}{}
	}
}
