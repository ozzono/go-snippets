package main

import (
	"log"
	"os/exec"
	"strings"
)

func main() {
	memfree()
}

func memfree() {
	arg := "adb shell cat /proc/meminfo |grep MemFree"
	args := strings.Split(arg, " ")
	x := string(args[0])
	args = args[1:]
	out, err := exec.Command(x, args...).CombinedOutput()
	if err != nil {
		log.Printf("Error: %v; Command: '%v'; Output: %v\n", err, arg, string(out))
		return
	}
	log.Printf("out: %v", string(out))
	// memfree, err := readXML("(\\d+)", out)
	// if err != nil {
	// 	log.Printf("%v", err)
	// }
	// if len(memfree) > 1 {
	// 	logToFile("memfree.log", memfree[1])
	// }
}
