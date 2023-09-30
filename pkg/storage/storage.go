package storage

import "github.com/filatkinen/storage/internal/storage"

func NewFileStorage(name string, data []byte) *storage.FileStorage {
	return storage.NewFileStorage(name, data)
}
