package storage

import "sync"

type pg struct {
	mu sync.RWMutex
}

func NewPersistant() (Storage, error) {
	return &pg{}, nil
}

func (p *pg) Increment(key string) error    { return nil }
func (p *pg) Value(key string) (int, error) { return 42, nil }
func (p *pg) Reset() error                  { return nil }
