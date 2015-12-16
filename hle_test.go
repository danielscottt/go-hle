package hle

import (
	"sync"
	"testing"
)

var (
	mu  = sync.Mutex{}
	hle = New()
)

func useThing(t int) {}

func BenchmarkHle(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			hle.Lock()
			thing := i
			hle.Unlock()

			useThing(thing)
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			thing := i
			mu.Unlock()

			useThing(thing)
			wg.Done()
		}()
	}
	wg.Wait()
}
