package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []byte(string("abc"))
	b := make([]byte, len(a))
	copy(b, a)
	fmt.Printf("equal: %v\n", reflect.DeepEqual(a, b))
}
