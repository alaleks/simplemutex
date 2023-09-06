// package simplemutex implements synchronization primitive providing
// mutual exception on a competitive read or write.
package simplemutex

import (
	"sync/atomic"
)

type (
	Mutex struct {
		numReaders atomic.Int32 // number of readers
		lockWrite  atomic.Bool  // lock flag, where true - locked
	}
)

// New creates a new poiner of mutex.
func New() *Mutex {
	return &Mutex{
		numReaders: atomic.Int32{},
		lockWrite:  atomic.Bool{},
	}
}

// Lock implements blocking read/write.
func (m *Mutex) Lock() {
	for m.numReaders.Load() != 0 ||
		!m.lockWrite.CompareAndSwap(false, true) {
	}
}

// Unlock implements unblocking read/write.
func (m *Mutex) Unlock() {
	m.lockWrite.Store(false)
}

// RLock implements blocking for read during write, but not blocking read.
func (m *Mutex) RLock() {
	for m.lockWrite.Load() {
	}

	m.numReaders.Add(1)
}

// RUnlock implements reduces the number of readers.
func (m *Mutex) RUnlock() {
	m.numReaders.Add(-1)
}
