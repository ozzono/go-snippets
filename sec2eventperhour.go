package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var s string

func init() {
	flag.StringVar(&s, "s", "0", "")
}

func main() {
	flag.Parse()
	if len(s) > 0 && s != "0" {
		calc(strings.Split(s, ","))
		return
	}
	list := []string{
		"69.59",
		"64.22",
		"69.59",
		"55.72",
		"56.89",
		"56.50",
		"58.16",
		"57.84",
	}
	calc(list)
}

func average(i float64) int {
	if i > 0 {
		i = 3600 / i
		fmt.Printf("Average per hour: %.f\n", i)
		return int(i)
	}
	fmt.Printf("s: %s\n", s)
	return -1
}

func calc(input []string) {
	total := 0
	for _, item := range input {
		i, err := strconv.ParseFloat(item, 64)
		if err != nil {
			fmt.Println(err)
		}
		if partial := average(i); partial > 0 {
			total += partial
		}
	}
	fmt.Printf("total: %d\n", total)
}
