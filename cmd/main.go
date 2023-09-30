package main

import (
	"fmt"

	"github.com/filatkinen/storage/pkg/storage"
)

func main() {

	file := storage.NewFileStorage("One", []byte("Hello"))
	fmt.Println(file)
}
