package main

import (
	"fmt"
	"time"
)

const sleeptime = 1000

func main() {
	timestamp()
}

func timestamp() {
	t1 := time.Now().Unix()
	sleeper()
	t2 := time.Now().Unix()
	fmt.Printf("diff: %d\n", t2-t1)
}

func sleeper() {
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
}
