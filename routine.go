package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	newFun := func(i string) (int, error) {
		out, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		return out, nil
	}
	var wg sync.WaitGroup
	wg.Add(1)
	out, err := go newFun("1")
	if err != nil {
		log.Println(err)
		wg.Done()
		return
	}
	wg.Done()
	fmt.Printf("%d\n", out)
}
