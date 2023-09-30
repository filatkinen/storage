package main

import (
	"fmt"
	"github.com/filatkinen/storage/v2/pkg/storage"
)

func main() {
	file := storage.NewFileStorage("One", []byte("Hello"), 10)
	fmt.Println(file)
}
