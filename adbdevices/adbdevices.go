package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	output := shell("adb devices")
	devices(output)
}

func devices(input string) {
	count := 0
	deviceID := []string{}
	for i, item := range strings.Split(input, "\n") {
		if strings.HasSuffix(item, "device") {
			fmt.Printf("[%d] - %s\n", i, item)
			count++
			deviceID = append(deviceID, deviceDesc(item))
		}
	}
	fmt.Printf("device count: %d\n", count)
	fmt.Printf("deviceID: %s\n", strings.Join(deviceID, ", "))
}

func deviceDesc(input string) string {
	return strings.Split(input, "	")[0]
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
