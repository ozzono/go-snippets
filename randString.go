package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println(randString(randInt(5) + 5))
}

func randInt(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano() + int64(n))).Intn(n)
}

func randString(maxSize int) string {
	charSet := "abcdefghijklmnopqrstuvxywz"
	output := ""
	j := randInt(maxSize)
	for i := 0; i < j; i++ {
		output += string(charSet[randInt(len(charSet))])
	}
	return string(output)
}
