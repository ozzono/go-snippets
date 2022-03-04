package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i, cancel := chamada()
	fmt.Println(i)
	cancel()
}

func chamada() (int, func()) {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10),
		func() {
			fmt.Println("defer func")
		}
}
