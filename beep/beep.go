package main

import (
	"log"

	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Alert("Title", "Message body", "")
	if err != nil {
		log.Println(err)
	}
}
