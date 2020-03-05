package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var (
	input string
	tap   bool
)

func init() {
	flag.StringVar(&input, "i", "", "Sets the xml styled input")
	flag.BoolVar(&tap, "tap", false, "Enables adb tap for inputed coords")
}

func main() {
	flag.Parse()
	if len(input) > 0 {
		xml2tap(input)
		x, y := xml2tap(input)
		fmt.Println("teste")
		if tap && x != 0 && y != 0 {
			shell(fmt.Sprintf("adb shell input tap %d %d", x, y))
			return
		}
	}
}

func xml2tap(xmlcoords string) (int, int) {
	fmt.Println("Parsing coords")
	openbracket := "["
	closebracket := "]"
	joinedbracket := "]["
	if string(xmlcoords[0]) == openbracket && string(xmlcoords[len(xmlcoords)-1]) == closebracket && strings.Contains(xmlcoords, joinedbracket) {
		stringcoords := strings.Split(xmlcoords, "][")
		leftcoords := strings.Split(string(stringcoords[0][1:]), ",")
		rightcoords := strings.Split(string(stringcoords[1][:len(stringcoords[1])-1]), ",")
		x1, err := strconv.Atoi(leftcoords[0])
		if err != nil {
			fmt.Printf("atoi err: %v", err)
			return 0, 0
		}
		y1, err := strconv.Atoi(leftcoords[1])
		if err != nil {
			fmt.Printf("atoi err: %v", err)
			return 0, 0
		}
		x2, err := strconv.Atoi(rightcoords[0])
		if err != nil {
			fmt.Printf("atoi err: %v", err)
			return 0, 0
		}
		y2, err := strconv.Atoi(rightcoords[1])
		if err != nil {
			fmt.Printf("atoi err: %v", err)
			return 0, 0
		}
		x := (x1 + x2) / 2
		y := (y1 + y2) / 2
		fmt.Printf("%s --- x: %d y: %d\n", xmlcoords, x, y)
		return x, y
	}
	return 0, 0
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
