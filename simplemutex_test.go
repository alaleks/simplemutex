package simplemutex_test

import (
	"sync"
	"testing"
	"time"

	"github.com/alaleks/simplemutex"
)

func TestMutexLock(t *testing.T) {
	var (
		counter = 0
		done    = make(chan struct{})
		mu      = simplemutex.New()
	)

	mu.Lock()

	go func() {
		mu.Lock()
		counter++
		mu.Unlock()
		done <- struct{}{}
	}()

	if counter != 0 {
		t.Errorf("expected counter to be 0, got %d", counter)
	}

	mu.Unlock()

	<-done

	if counter != 1 {
		t.Errorf("expected counter to be 1, got %d", counter)
	}
}

func TestMutexRLock(t *testing.T) {
	var (
		counter = 0
		done    = make(chan struct{})
		mu      = simplemutex.New()
	)

	go func() {
		mu.Lock()
		counter++
		mu.Unlock()
		done <- struct{}{}
	}()

	go func() {
		mu.RLock()
		time.Sleep(100 * time.Millisecond)

		if counter != 0 {
			t.Errorf("expected counter to be 0, got %d", counter)
		}

		mu.RUnlock()
	}()

	<-done

	if counter != 1 {
		t.Errorf("expected counter to be 1, got %d", counter)
	}
}

func TestMutexResult(t *testing.T) {
	var (
		num    = 1000
		result int
		wg     sync.WaitGroup
		mu     = simplemutex.New()
	)

	for i := 0; i < num; i++ {
		wg.Add(1)

		go func() {
			mu.Lock()
			result++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	if result != num {
		t.Errorf("expected result after goroutine's launch to be %d, got %d", num, result)
	}
}

func BenchmarkRWMutexWrite(b *testing.B) {
	var (
		mu      sync.RWMutex
		counter int
	)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkMutexWrite(b *testing.B) {
	var (
		mu      sync.Mutex
		counter int
	)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkSimpleMutex(b *testing.B) {
	var (
		mu      = simplemutex.New()
		counter int
	)

	for i := 1; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkMutexRead(b *testing.B) {
	var (
		mu   sync.RWMutex
		data = map[int]string{1: "test"}
	)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		mu.RLock()
		_, _ = data[i]
		mu.RUnlock()
	}
}

func BenchmarkSimpleMutexRead(b *testing.B) {
	var (
		mu   = simplemutex.New()
		data = map[int]string{1: "test"}
	)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		mu.RLock()
		_, _ = data[i]
		mu.RUnlock()
	}
}
