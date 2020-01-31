package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var input string

func init() {
	flag.StringVar(&input, "i", "", "Sets the xml styled input")
}


func main() {
	flag.Parse()
	if len(input) > 0 {
		if strings.Contains(input, ",") {
			for _, item := range strings.Split(input, "|") {
				xml2tap(item)
			}
			return
		}
		xml2tap(input)
		return
	}
	xml2tap("[90,1204][990,1303]")
	xml2tap("[90,1225][376,1281]")
}

func xml2tap(xmlcoords string) {
	openbracket := "["
	closebracket := "]"
	joinedbracket := "]["
	if string(xmlcoords[0]) == openbracket && string(xmlcoords[len(xmlcoords)-1]) == closebracket && strings.Contains(xmlcoords, joinedbracket) {
		stringcoords := strings.Split(xmlcoords, "][")
		leftcoords := strings.Split(string(stringcoords[0][1:]), ",")
		rightcoords := strings.Split(string(stringcoords[1][:len(stringcoords[1])-1]), ",")
		x1, err := strconv.Atoi(leftcoords[0])
		if err != nil {
			return
		}
		y1, err := strconv.Atoi(leftcoords[1])
		if err != nil {
			return
		}
		x2, err := strconv.Atoi(rightcoords[0])
		if err != nil {
			return
		}
		y2, err := strconv.Atoi(rightcoords[1])
		if err != nil {
			return
		}
		x := (x1 + x2) / 2
		y := (y1 + y2) / 2
		fmt.Printf("%s --- x: %d y: %d\n", xmlcoords, x, y)
	}
}
