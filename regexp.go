package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

var (
	exp         string
	inputString string
	filepath    string
)

func init() {
	flag.StringVar(&exp, "e", "", "Regular expression")
	flag.StringVar(&inputString, "s", "", "Input string on terminal")
	flag.StringVar(&filepath, "f", "", "Input file path")
}

func main() {
	flag.Parse()
	text := inputString
	filepath = "/home/hugo/Projects/bf/vtefortaleza/dump.xml"
	if filepath != "" && len(filepath) > 0 {
		fmt.Printf("Using content from %s\n", filepath)
		file, err := readfile(filepath)
		if err != nil {
			fmt.Printf("readfile err: %v", err)
			return
		}
		text = file
	}
	exp = "wprmenu_bar.*?(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])"
	output := myregexp(exp, text)
	if len(output) > 1 {
		for i := 1; i < len(output); i++ {
			fmt.Printf("match[%d]: %s\n", i, output[i])
		}
	}
}

func myregexp(exp, text string) []string {
	fmt.Println(text)
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		fmt.Printf("Unable to find match for exp %s\n", exp)
		return []string{}
	}
	return match
}

func readfile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
