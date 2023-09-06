# Simplemutex

This is a very simple package based on atomic that implements read/write locking on competitive execution.

## Installation

```go get -u github.com/alaleks/simplemutex```

## Usage
```
// read/write lock
mu   = simplemutex.New()
mu.Lock()
mu.Unlock()

// non-blocking reading
mu2   = simplemutex.New()
mu2.RLock()
mu2.RUnlock()
``` 

## Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/alaleks/simplemutex
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
BenchmarkRWMutexWrite-12        154265481                7.943 ns/op           0 B/op          0 allocs/op
BenchmarkMutexWrite-12          265443903                4.509 ns/op           0 B/op          0 allocs/op
BenchmarkSimpleMutex-12         300189049                3.887 ns/op           0 B/op          0 allocs/op
BenchmarkMutexRead-12           165793032                6.985 ns/op           0 B/op          0 allocs/op
BenchmarkSimpleMutexRead-12     178749810                6.629 ns/op           0 B/op          0 allocs/op
```