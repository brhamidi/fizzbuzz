package storage

import (
	"sync"
)

// verify interface compliance
var _ Storage = (*inmemory)(nil)

type inmemory struct {
	mu sync.RWMutex
	// represent map of [request]hit
	data    map[string]int
	maxKey  string
	maxHits int
}

// TODO apply the logic of keep the MAX value in seperate variable to prevent looping in GetStats operation

func NewInmemory() *inmemory {
	return &inmemory{data: make(map[string]int)}
}

func (i *inmemory) Increment(key string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	v, ok := i.data[key]
	if !ok {
		i.data[key] = 1
	} else {
		i.data[key] = v + 1
	}

	if i.data[key] > i.maxHits {
		i.maxHits = i.data[key]
		i.maxKey = key
	}
	return nil
}

func (i *inmemory) Max() (string, int, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.maxKey, i.maxHits, nil
}

func (i *inmemory) Reset() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data = make(map[string]int)
	i.maxHits = 0
	i.maxKey = ""
	return nil
}
