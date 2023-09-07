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
BenchmarkRWMutexWrite-12        149813430                8.029 ns/op           0 B/op          0 allocs/op
BenchmarkMutexWrite-12          252883302                4.507 ns/op           0 B/op          0 allocs/op
BenchmarkSimpleMutex-12         275072392                3.955 ns/op           0 B/op          0 allocs/op
BenchmarkMutexRead-12           256254141                4.075 ns/op           0 B/op          0 allocs/op
BenchmarkSimpleMutexRead-12     284053776                3.757 ns/op           0 B/op          0 allocs/op
```

## Attention
Since the package is based on atomic operations, performance tests on different hardware and operating systems will vary.