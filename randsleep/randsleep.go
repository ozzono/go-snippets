package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// sample()
	nap()
}

func nap() {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(10) + 1
	fmt.Printf("delay: %vs\n", delay)
	time.Sleep(time.Duration(delay) * time.Second)
}
