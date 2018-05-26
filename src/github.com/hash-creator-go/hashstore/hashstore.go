package hashstore

import (
	"sync/atomic"
	"crypto/sha512"
	"encoding/base64"
	"time"
)

// Defines interface for HashStore implementations. Start will initialize the store and make it ready to receive
// messages. Stop will allow current processing to finish and immediately stop receiving new messages.
type HashStore interface {
	Start()
	Stop()
	GetHash(id uint64) string
	StoreHash(pw string) uint64
	GetStats() (total uint64, avg float64)
	IsShutdown() bool
}

// Backing collection is a simple golang map with an int key and string hash value. Unlike reads, writes are
// handled in a separate thread
type SimpleKVHashStore struct {
	m map[uint64]string
	sm chan storemsg
	ctr uint64
	// Total Time Processing, used to calculate avg stat
	ttp int64
	shutdown bool
}

type storemsg struct {
	id uint64
	pw string
	start time.Time
}

func (hs *SimpleKVHashStore) Start() {
	hs.m = make(map[uint64]string)
	hs.sm = make(chan storemsg, 1000);
	go func() {
		for {
			m, ok := <-hs.sm
			hs.m[m.id] = hash(m.pw)
			totalTime := time.Now().Sub(m.start);
			hs.ttp += totalTime.Nanoseconds()

			if !ok {
				break;
			}
		}
	}()
}

func hash(pw string) string {
	h := sha512.Sum512([]byte(pw))
	return base64.StdEncoding.EncodeToString(h[:])
}

func (hs *SimpleKVHashStore) Stop() {
	close(hs.sm)
	hs.shutdown = true
}

// Returns the hash stored at the given id, the read of the underlying collection is
// handled in the thread of the caller
func (hs *SimpleKVHashStore) GetHash(id uint64) string {
	return hs.m[id];
}

// Returns the id for eventual retrieval of the stored hash, store operation is handled in a separate thread.
// Returns 0 if store has already been shutdown
func (hs *SimpleKVHashStore) StoreHash(pw string) uint64 {
	startTime := time.Now()
	if (hs.shutdown) {
		return 0
	}

	id := atomic.AddUint64(&(hs.ctr), 1);
	// if our channel buffer is full, the store could cause the calling thread to block, so doing the store
	// in a new goroutine
	go func() {
		hs.sm <- storemsg{id, pw, startTime}
	}()

	return id;
}

func (hs *SimpleKVHashStore) GetStats() (total uint64, avg float64) {
	return hs.ctr, float64(hs.ttp) / float64(hs.ctr)
}

func (hs *SimpleKVHashStore) IsShutdown() bool {
	return hs.shutdown
}





