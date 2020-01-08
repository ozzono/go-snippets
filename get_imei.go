package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	getimei()
}

func getimei() {
	cmd("adb shell service call iphonesubinfo 1 | awk -F \"'\" '{print $2}' | sed '1 d' | tr -d '.' | awk '{print}' ORS=")
}

func cmd(arg string) {
	log.Printf("Executing command with Go: %v", arg)
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
