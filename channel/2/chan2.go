package main

import (
	"fmt"
	"time"
)

func main() {
	control := make(chan bool)
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().Nanosecond()/100, "i", i)
		go func(i int, c chan bool) {
			if i%2 == 0 {
				fmt.Println(time.Now().Nanosecond()/100, "mod", i)
				control <- true
			}
		}(i, control)
	}
	resp := <-control
	if resp {
		fmt.Printf("%#v\n", resp)
	}
}
