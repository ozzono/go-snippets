package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 10

func main() {
	list := []int{}
	for i := 0; i < defaultSize; i++ {
		seed := (time.Now().Nanosecond() + i)
		fmt.Println("seed", seed)
		list = append(list, rand.Intn(seed)%10)
		time.Sleep(time.Duration(seed%100) * time.Nanosecond)
	}
	control := make(chan bool, len(list)) //buffered channel with size equal to len(list)
	for _, item := range list {
		go func(item int) { // this triggers an ascynchronous goroutine
			// since it's asynchronous the runtime won't wait for the loop to move on
			time.Sleep(time.Second) // all routines will take at least 1s to be executed
			if item%2 == 1 {
				control <- false // add false to the channel buffer
				return
			}
			control <- true // add true to the channel buffer
		}(item)
	}

	for i := 0; i < len(list); i++ {
		fmt.Println(time.Now().Nanosecond(), <-control) // when using <- it will wait for data comming from the control channel
	}
}
