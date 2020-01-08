package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	data()
}

func data() {
	DtUpdateBalance := "2019-10-15T15:04:05.000Z"
	DtOldRecharge, err := time.Parse("2006-01-02T15:04:05.000Z", DtUpdateBalance)
	if err != nil {
		log.Printf("Failed to parse date %v: %v", DtUpdateBalance, err)
	}
	fmt.Printf("formated date: %v\n", DtOldRecharge)
}
