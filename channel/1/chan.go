package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	control := false
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().Nanosecond()/100, "i", i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				mu.Lock()
				fmt.Println(time.Now().Nanosecond()/100, "mod", i)
				control = true
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	if control {
		fmt.Printf("%v\n", control)
	}
}
