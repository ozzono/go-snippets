package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stop := looped("babe")
	defer stop()
	fmt.Println("out")
	time.Sleep(5 * time.Second)
}

func looped(input string) func() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for true {
			time.Sleep(1 * time.Second)
			fmt.Printf("sleep %s\n", input)
		}
	}()
	return func() {
		fmt.Printf("wake up %s\n", input)
		wg.Done()
	}
}
