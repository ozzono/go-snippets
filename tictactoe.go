package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	coords := [3][3]bool{}
	i0, i1 := readInput()
	coords[i0][i1] = !coords[i0][i1]
	fmt.Printf("output: %d-%d\n", i0, i1)
	// fmt.Println(coords[i0][i1])
}

func readInput() (int, int) {
	log.Println("Insert a coord pair")
	log.Println("Allowed format: n-n; n |0--2")
	input := ""
	for matched := false; !matched; {
		something, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break
		}
		input = string(something)
		matched = match("(\\d-\\d)", input)
		if !matched {
			log.Println("Invalid format; try again")
		}
	}
	return parseCoord(input)
}

func parseCoord(input string) (int, int) {
	split := strings.Split(input, "-")
	i0, _ := strconv.ParseInt(strings.Replace(strings.Replace(split[0], "\n", "", -1), " ", "", -1), 10, 0)
	i1, _ := strconv.ParseInt(strings.Replace(strings.Replace(split[1], "\n", "", -1), " ", "", -1), 10, 0)
	return int(i0), int(i1)
}

func match(exp, text string) bool {
	return regexp.MustCompile(exp).MatchString(text)
}
