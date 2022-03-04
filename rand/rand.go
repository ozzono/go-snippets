package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("%v\n", randomPhoneNumber())
}

func oldrand() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	// newid := strconv.FormatUint(rnd, 16)
	fmt.Printf("%s\n", rnd%len("lalala"))
}

func simplerand() {
	test := "meu teste de tamanho da string lalalala"
	r := rand.Intn(len(test))
	fmt.Printf("r: %v\n", r)
}

func diceroller(myvar []string) int {
	flip := 0
	if len(myvar) > 1 {
		mymod := len(myvar) - 1
		if mymod == 1 {
			mymod++
		}
		flip = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(myvar))
		fmt.Printf("flip: %v\n", flip)
	}
	return flip
}

func randomPhoneNumber() string {
	ddd := []string{"11", "12", "13", "14", "15", "16", "17", "18", "19", "21", "22", "24", "27", "28", "31", "32", "33", "34", "35", "37", "38", "41", "42", "43", "44", "45", "46", "47", "48", "49", "51", "53", "54", "55", "61", "62", "63", "64", "65", "66", "67", "68", "69", "71", "73", "74", "75", "77", "79", "81", "82", "83", "84", "85", "86", "87", "88", "89", "91", "92", "93", "94", "95", "96", "97", "98", "99"}

	return ddd[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(67)] + "9" + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999999))
}
