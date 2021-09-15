package storage

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pg struct {
	mu     sync.RWMutex
	driver *gorm.DB
}

func NewPersistant(host, user, passwd, name string) (*pg, error) {
	var err error
	var driver *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, passwd, name)
	driver, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return &pg{}, err
	}
	return &pg{driver: driver}, nil
}

func (p *pg) Increment(key string) error    { return nil }
func (p *pg) Value(key string) (int, error) { return 42, nil }
func (p *pg) Reset() error                  { return nil }
