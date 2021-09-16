package storage

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pg struct {
	mu     sync.RWMutex
	driver *gorm.DB
}

type Data struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Keys string    `gorm:"uniqueIndex"`
	Hits int       `gorm:"default:1"`
}

func NewPersistant(host, user, passwd, name string) (*pg, error) {
	var err error
	var driver *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, passwd, name)
	driver, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return &pg{}, err
	}

	if err := driver.AutoMigrate(&Data{}); err != nil {
		return &pg{}, err
	}

	return &pg{driver: driver}, nil
}

func (p *pg) Increment(key string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	entry := Data{Keys: key}

	err := p.driver.Where(&entry).Take(&entry).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return p.driver.Create(&entry).Error
	}
	if err != nil {
		return err
	}

	entry.Hits = entry.Hits + 1
	return p.driver.Save(&entry).Error
}

func (p *pg) Max() (string, int, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return "", 0, nil
}

func (p *pg) Reset() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.driver.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Data{}).Error
}
