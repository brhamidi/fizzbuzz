package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

// verify interface compliance
var _ Storage = (*persistent)(nil)

type persistent struct {
	mu     sync.RWMutex
	driver *redis.Client
}

const (
	MAX_KEY  = "max_key"
	MAX_HITS = "max_hits"
)

func initMaxValue(driver *redis.Client) error {
	err := driver.Set(context.Background(), MAX_KEY, "", 0).Err()
	if err != nil {
		return err
	}

	err = driver.Set(context.Background(), MAX_HITS, "0", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewPersistant(host, port string) (*persistent, error) {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s", host, port))
	if err != nil {
		return &persistent{}, err
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Get(context.Background(), MAX_KEY).Result()
	if err == nil {
		return &persistent{driver: rdb}, nil
	}

	if err != redis.Nil {
		return &persistent{}, err
	}

	if err := initMaxValue(rdb); err != nil {
		return &persistent{}, err
	}

	return &persistent{driver: rdb}, nil
}

func (p *persistent) Increment(key string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	result, err := p.driver.Incr(context.Background(), key).Result()
	if err != nil {
		return err
	}

	currMaxHits, err := p.driver.Get(context.Background(), MAX_HITS).Int64()
	if err != nil {
		return err
	}

	if currMaxHits >= result {
		return nil
	}

	err = p.driver.Set(context.Background(), MAX_KEY, key, 0).Err()
	if err != nil {
		return err
	}

	err = p.driver.Set(context.Background(), MAX_HITS, result, 0).Err()
	if err != nil {
		// TODO use transaction here to prevent corrupted data
		return err
	}
	return nil
}

func (p *persistent) Max() (string, int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, err := p.driver.Get(context.Background(), MAX_KEY).Result()
	if err != nil {
		return "", 0, err
	}

	hits, err := p.driver.Get(context.Background(), MAX_HITS).Int()
	if err != nil {
		return "", 0, err
	}

	return key, hits, nil
}

func (p *persistent) Reset() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.driver.FlushDB(context.Background()).Err(); err != nil {
		return err
	}

	if err := initMaxValue(p.driver); err != nil {
		return err
	}
	return nil
}
