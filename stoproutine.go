package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go func() {
		fmt.Println("inner code")
	}()
	time.Sleep((1 * time.Second))
	close(quit)
}
