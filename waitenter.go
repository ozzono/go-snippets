package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	waitEnter("gotcha!")
}

func waitEnter(text string) {
	log.Printf("Waiting for '%v'", text)
	log.Printf("Press <enter> to continue or <ctrl+c> to interrupt")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	log.Printf("Now, where was I?")
	log.Printf("Oh yes...\n\n")
}
