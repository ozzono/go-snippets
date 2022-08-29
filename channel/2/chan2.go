package main

import (
	"fmt"
	"time"
)

func main() {
	control := make(chan bool)
	for i := 0; i < 600; i++ {
		fmt.Println(time.Now().Nanosecond()/100, "i", i)
		go func(i int, c chan bool) {
			time.Sleep(1 * time.Second)
			if i%100 == 0 {
				fmt.Println(time.Now().Nanosecond()/100, "mod", i)
				c <- true
			}
		}(i, control)
	}
	resp := <-control
	if resp {
		fmt.Printf("%#v\n", resp)
	}
}
