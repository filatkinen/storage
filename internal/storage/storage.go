package storage

import "fmt"

type FileStorage struct {
	Name  string
	Data  []byte
	Inode int
}

func NewFileStorage(name string, data []byte, i int) *FileStorage {
	return &FileStorage{
		Name:  name,
		Data:  data,
		Inode: i,
	}
}
func (f *FileStorage) String() string {
	return fmt.Sprintf("file:(%s,%s,%d)", f.Name, f.Data, f.Inode)
}
