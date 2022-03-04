package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	lista := []string{
		"AL$%lx8700@",
		"BX!@qr4200@",
		"DQ!$bx4300@",
		"EF@?jv9900@",
		"EM$+qi4100@",
		"JB+$ap8800@",
		"JJb$wi6700@",
		"JPb$qz1100@",
		"JW$?hg5300@",
		"KO%@ty5900@",
		"LC%?yb2200@",
		"NR%+pn3700@",
		"PM%!bt9900@",
		"QSbbbj3400@",
		"RD@@op3500@",
		"SA!+is2600@",
		"TJ%bfe8700@",
		"UZ@$ea4500@",
		"VL?$ox7400@",
		"VN$$mk0400@",
		"XO$%sl9300@",
		"YC@+zx6500@",
		"ZZ@+ga1000@",
	}
	for i, _ := range lista {
		if ValidatePW(lista[i]) {
			fmt.Printf(" is invalid\n")
		} else {
			fmt.Printf(" is valid\n")
		}
	}
}

func ValidatePW(pw string) bool {
	log.Printf("Validating password %s", pw)
	if len(pw) < 10 || len(pw) > 12 || strings.Contains(pw, "*") {
		return false
	}
	expression := []string{"([A-Z]+)", "([a-z]+)", "(\\d+)", "([@|%]+)"}
	for i, _ := range expression {
		value := myregexp(expression[i], pw)
		if len(value) == 0 {
			fmt.Printf("Expression %s not satisfied", expression[i])
			return false
		}
	}
	return true
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
