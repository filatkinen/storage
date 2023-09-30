package storage

import "github.com/filatkinen/storage/v2/internal/storage"

func NewFileStorage(name string, data []byte, i int) *storage.FileStorage {
	return storage.NewFileStorage(name, data, i)
}
