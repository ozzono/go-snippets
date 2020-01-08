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
	x := string(args[0])
	args = args[1:]
	out, err := exec.Command(x, args...).Output()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if out != nil && len(out) > 0 {
		fmt.Printf("Output:\n %s", out)
	}
	fmt.Printf("\nDone\n")
}
