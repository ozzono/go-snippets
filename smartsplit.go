package main

import "strings"

func main() {
	text := "'a, b, a'"
	SmartSplit(text)
}

func SmartSplit(input string) []string {
	if strings.Contains(input, ",") {
		commas := commaIndex(input)
		output := []string{}
		for i := len(commas) - 1; i >= 0; i-- {
			output = append([]string{string(input[commas[i]+1:])}, output...)
			input = input[:commas[i]]
			if i == 0 {
				output = append([]string{input}, output...)
			}
		}
		return output
	}
	return []string{input}
}

func commaIndex(input string) []int {
	output := []int{}
	var skip bool
	for i, _ := range input {
		if string(input[i]) == "'" {
			skip = !skip
		}
		if skip {
			continue
		}
		if string(input[i]) == "," {
			output = append(output, i)
		}
	}
	return output
}
