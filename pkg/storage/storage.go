package storage

type Storage interface {
	Increment(key string) error
	Max() (string, int, error)
	Reset() error
}
