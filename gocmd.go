package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	cmd(strings.Join(args, " "))
}

func cmd(arg string) {
	log.Printf("Executing command with Go: " + arg)
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if out != nil && len(out) > 0 {
		fmt.Printf("Output:\n %s", out)
	}
	fmt.Printf(">> Done <<\n")
}
