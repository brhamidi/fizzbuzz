package storage

import (
	"sync"
)

type inmemory struct {
	mu   sync.RWMutex
	data map[string]int
}

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
	return nil
}

func (i *inmemory) Max() (string, int, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if len(i.data) == 0 {
		return "", 0, nil
	}

	key, hits := "", 0
	for k, v := range i.data {
		if v > hits {
			key, hits = k, v
		}
	}

	return key, hits, nil
}

func (i *inmemory) Reset() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data = make(map[string]int)
	return nil
}
