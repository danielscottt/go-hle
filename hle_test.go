package hle

import (
	"sync"
	"testing"
)

var (
	hle = New()
	mu  = sync.Mutex{}
)

func BenchmarkHle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hle.Lock()
		hle.Unlock()
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}
