package main

import (
	"fmt"
	"sync"
	"time"
)

const defaultSize = 10

func main() {
	// procedural()
	// concurentNonStop()
	// concurentWithWait_without_Channel()
	// concurentWithWait_with_Channel()
}

// each loop cicle will be executed in line after the other
func procedural() {
	start := time.Now()
	for i := 0; i < defaultSize; i++ {
		time.Sleep(time.Second)
		fmt.Printf("procedural timestamp: %d - %d\n", time.Now().Nanosecond(), i)
	}
	fmt.Printf("procedural took %dms\n", time.Now().Sub(start).Milliseconds())
}

// each loop cicle will be executed concurrently
func concurentNonStop() {
	start := time.Now()
	for i := 0; i < defaultSize; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Printf("concurentNonStop timestamp ns: %d - %d\n", time.Now().Nanosecond(), i)
		}(i)
	}
	// probably the runtime will finish the funcion before running the routines
	fmt.Printf("concurrentNonStop took %dms\n", time.Now().Sub(start).Milliseconds())
}

// each loop cicle will be executed concurrently using sync.WaitGroup
func concurentWithWait_without_Channel() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < defaultSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Printf("concurentWithWait_without_Channel timestamp ns: %d - %d\n", time.Now().Nanosecond(), i)
		}(i)
	}
	wg.Wait()
	// it will probably take about 1s to run all routines
	fmt.Printf("concurentWithWait_without_Channel took %dms\n", time.Now().Sub(start).Milliseconds())
}

// each loop cicle will be executed concurrently using channel
func concurentWithWait_with_Channel() {
	start := time.Now()
	control := make(chan struct{}, defaultSize)
	for i := 0; i < defaultSize; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Printf("concurentWithWait_with_Channel timestamp ns: %d - %d\n", time.Now().Nanosecond(), i)
			control <- struct{}{}
		}(i)
	}
	for i := 0; i < defaultSize; i++ {
		<-control
	}
	fmt.Printf("concurentWithWait_with_Channel took %dms\n", time.Now().Sub(start).Milliseconds())
}
