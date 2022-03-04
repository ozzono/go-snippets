package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// fmt.Println(cmd("ls "))
	fmt.Println(shell("ls $(pwd)"))
}

func shell(arg string) string {
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v';\nOutput: %v;\nError: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}
