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
	for i := 0; i < 600; i++ {
		fmt.Println(time.Now().Nanosecond()/100, "i", i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%10 == 0 {
				mu.Lock()
				fmt.Println(time.Now().Nanosecond()/100, "mod", i)
				control = true
				mu.Unlock()
			}
			time.Sleep(1 * time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Printf("%v\n", control)
}
