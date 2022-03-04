package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := WriteToFile("txt.log", "oaoaoao\n")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

func WriteToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
