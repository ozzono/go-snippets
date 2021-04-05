package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	go func() {
		for true {
			fmt.Print("inner code")
		}
	}()
	time.Sleep((1 * time.Second))
	close(quit)
}
