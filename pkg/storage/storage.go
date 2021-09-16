package storage

//go:generate mockgen -package=mock -source=storage.go -destination=$MOCK_FOLDER/storage.go Storage

type Storage interface {
	Increment(key string) error
	Max() (string, int, error)
	Reset() error
}
