package main

import (
	"fmt"
	"log"
	"strings"

	shell "github.com/ozzono/go-shell"
)

func main() {
	items := filter(shell.Cmd("snap list --all"))
	if len(items) > 0 {
		log.Printf("Removing %d disabled snaps", len(items))
	} else {
		log.Println("There are no disabled snap available for removal")
	}
	for _, item := range items {
		shell.Cmd(fmt.Sprintf("snap remove %s --revision=%s", item[0], item[1]))
	}
}

func filter(input string) [][]string {
	output := [][]string{}
	for i, item := range strings.Split(input, "\n") {
		if len(item) > 0 && i > 0 && strings.Contains(item, "disabled") {
			element := []string{}
			for _, el := range strings.Split(item, " ") {
				if len(el) > 0 {
					element = append(element, el)
				}
			}
			output = append(output, []string{element[0], element[2]})
		}
	}
	return output
}
