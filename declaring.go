package main

import (
	"flag"
	"fmt"
)

var t1 bool

func init() {
	flag.BoolVar(&t1, "t", false, "Testing bool flag")
}

func main() {
	flag.Parse()
	fmt.Println(t1)
}
