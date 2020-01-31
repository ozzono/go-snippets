package main

import (
	"flag"
	"fmt"
	"regexp"
)

var (
	exp         string
	inputString string
)

func init() {
	flag.StringVar(&exp, "e", "", "Regular expression")
	flag.StringVar(&inputString, "s", "", "Input string on terminal")
}

func main() {
	flag.Parse()
	exp = "bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"2\" text=\"DDD + Telefone\""
	output := myregexp(exp, inputString)
	if len(output) > 1 {
		for i := 1; i < len(output); i++ {
			fmt.Printf("match[%d]: %s\n", i, output[i])
		}
	}
}

func myregexp(exp, text string) []string {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		fmt.Println("Unable to find match for regexp")
		return []string{}
	}
	return match
}
