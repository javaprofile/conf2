package main

import (
	"fmt"
	"sync"
	"time"
)

type VersionedValue struct {
	value     string
	timestamp int64
}

type MVCCStore struct {
	mu     sync.RWMutex
	store  map[string][]VersionedValue
	latest map[string]int64
}

func NewMVCCStore() *MVCCStore {
	return &MVCCStore{
		store:  make(map[string][]VersionedValue),
		latest: make(map[string]int64),
	}
}

func (m *MVCCStore) Write(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	timestamp := time.Now().UnixNano()
	m.store[key] = append(m.store[key], VersionedValue{value, timestamp})
	m.latest[key] = timestamp
}

func (m *MVCCStore) Read(key string, readTime int64) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	versions := m.store[key]
	for i := len(versions) - 1; i >= 0; i-- {
		if versions[i].timestamp <= readTime {
			return versions[i].value
		}
	}
	return ""
}

func main() {
	store := NewMVCCStore()
	store.Write("x", "v1")
	time.Sleep(1 * time.Millisecond)
	store.Write("x", "v2")
	readTime := time.Now().UnixNano()
	time.Sleep(1 * time.Millisecond)
	store.Write("x", "v3")
	fmt.Println("Read at", readTime, "->", store.Read("x", readTime))
}
