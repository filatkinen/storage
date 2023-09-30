package storage

import "fmt"

type FileStorage struct {
	Name string
	Data []byte
}

func NewFileStorage(name string, data []byte) *FileStorage {
	return &FileStorage{
		Name: name,
		Data: data,
	}
}
func (f *FileStorage) stringer() string {
	return fmt.Sprintf("file:(%s,%s)", f.Name, f.Data)
}
