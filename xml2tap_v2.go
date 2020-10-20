package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	input string
)

func init() {
	flag.StringVar(&input, "i", "", "Sets the xml styled input")
}

func main() {
	flag.Parse()
	if len(input) > 0 {
		fmt.Println(xml2tap(input))
	}
}

func xml2tap(xmlcoords string) (int, int) {
	fmt.Println("Parsing coords")
	if match("(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])", xmlcoords) {
		stringcoords := [][]int{}
		for _, item := range strings.Split(xmlcoords, "][") {
			item = strings.Replace(item, "]", "", -1)
			item = strings.Replace(item, "[", "", -1)
			icoords := []int{}
			for _, item := range strings.Split(item, ",") {
				iItem, _ := strconv.Atoi(item)
				icoords = append(icoords, iItem)
			}
			stringcoords = append(stringcoords, icoords)
		}
		return stringcoords[0][0] + newRandNumber(stringcoords[0][1]-stringcoords[0][0]), stringcoords[1][0] + newRandNumber(stringcoords[1][1]-stringcoords[1][0])
	}
	log.Println("Invalid string input")
	return 0, 0
}

func match(exp, text string) bool {
	return regexp.MustCompile(exp).MatchString(text)
}

func newRandNumber(i int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}
