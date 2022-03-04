package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", exists("samples/"))
}

// exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
