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
	shell(args)
}

func shell(arg string) string {
	fmt.Println("shell: " + arg)
	args := strings.Split(arg, " ")
	if len(args) == 1 {
		args = append(args, "")
	} else if len(args) < 1 {
		log.Println("Invalid command")
		return ""
	}
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v'; Output: %v; Error: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}
