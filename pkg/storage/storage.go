package storage

type Storage interface {
	Increment(key string) error
	Value(key string) (int, error)
	Reset() error
}
