package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	sqlInClauseList("teste1 teste2")
}

func sqlInClauseList(input string) string {
	// formats a sequence of inputs into the IN clause of MySQL
	inputList := strings.Split(input, " ")
	output := ""
	for i, item := range inputList {
		if len(item) > 0 {
			if i > 0 {
				output += fmt.Sprintf(",")
			}
			output += fmt.Sprintf("\"%s\"", item)
		}
	}
	if shlog {
		log.Printf("input : %s", input)
		log.Printf("output: %s", output)
	}
	return output
}
