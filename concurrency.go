package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	outer sync.WaitGroup
	inner sync.WaitGroup
)

const count = 4

func main() {
	clock := time.Now().Unix()
	for i := 0; i < count; i++ {
		outer.Add(1)
		go outerTask(i)
	}
	outer.Wait()
	fmt.Printf("Done in %ds\n", time.Now().Unix()-clock)
}

func outerTask(i int) {
	defer outer.Done()
	for j := 0; j < count; j++ {
		inner.Add(1)
		go innerTask(i, j)
	}
	inner.Wait()
	fmt.Printf("outer done %d\n", i)
}

func innerTask(i, j int) {
	defer inner.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("inner [%d:%d]\n", i, j)
}
