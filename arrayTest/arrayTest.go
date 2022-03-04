package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// definiying the slices
	arr1 := []string{"a", "b", "c", "e", "f"}
	arr2 := arr1
	arr3 := make([]string, len(arr1))
	i := copy(arr3, arr1)
	fmt.Printf("copied %d\n", i)

	arr1[0] = "test"
	// changed a value of arr1

	// testing what happened to arr2 and arr3
	fmt.Println("arr2[0]: %v", arr2[0])
	fmt.Println("arr3[0]: %v", arr3[0])
}

func boggleMe() {
	arr1 := []string{"a", "b", "c", "d", "e", "f"}
	log.Printf("before bugMe: %s", strings.Join(arr1, ", "))
	bugMe(arr1)
	log.Printf("after bugMe: %s", strings.Join(arr1, ", "))
}

func bugMe(arr []string) {
	for len(arr) > 0 {
		log.Printf("	[%d] before: %s", len(arr), strings.Join(arr, ", "))
		arr = draw(arr, randInt(len(arr)))
		log.Printf("	[%d] after:  %s", len(arr), strings.Join(arr, ", "))
		fmt.Println("")
	}
}

func randInt(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano() + int64(n))).Intn(n)
}

func draw(arr []string, index int) []string {
	arr = make([]string, len(arr))
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}
