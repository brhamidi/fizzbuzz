package storage

//go:generate mockgen -package=mock -source=storage.go -destination=$MOCK_FOLDER/storage.go Storage

// Interface which all storage type must implement while http.Handler use this interface to be operational
type Storage interface {
	Increment(key string) error
	Max() (string, int, error)
	Reset() error
}
