package main

import (
	"log"
	"regexp"
	"strings"
)

func main() {
	// coord :=
	// log.Printf("coord: %s", coord)
	// log.Printf("match: %v", match(coord, "Pagar.+?(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])"))
	exps := [][2]string{
		[2]string{"\\d+", "123"},
		[2]string{strings.Replace("!!![!!d!+,!!d!+!!!]!!![!!d!+,!!d!+!!!]", "!", "\\", -1), "Pagar.+?(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])"},
	}
	for i := range exps {
		log.Printf("% 5t  exp: %s - text: %s", match(exps[i][0], exps[i][1]), exps[i][0], exps[i][1])
	}
}

func match(exp, text string) bool {
	return regexp.MustCompile(exp).MatchString(text)
}
