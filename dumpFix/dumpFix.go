package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	inputString string
)

func init() {
	flag.StringVar(&inputString, "s", "", "Input string on terminal")
}

func main() {
	flag.Parse()
	dumpFix(inputString)
}

func dumpFix(input string) string {
	rows := strings.Split(input, "\n")
	for i := range rows {
		if strings.Contains(rows[i], "INSERT  IGNORE INTO `") {
			data := strings.Split(myregexp("INSERT\\s+IGNORE INTO `.*` VALUES \\((.*)\\)", rows[i]), ",")
			for j := range data {
				if strings.HasSuffix(data[j], "\\'") {
					fmt.Printf("Before fix: %s\n", rows[i])
					data[j] = fmt.Sprintf("%v%v", data[j][:len(data[j])-2], data[j][len(data[j])-1:])
					fmt.Printf("After  fix: %s\n", rows[i])
				}
			}
			rows[i] = fmt.Sprintf("%s VALUES %s", strings.Split(rows[i], "VALUES")[0], strings.Join(data, ","))
			fmt.Println()
		}
	}
	return strings.Join(rows, "\n")
}

func myregexp(exp, text string) string {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		log.Println("Unable to find match for regexp")
		return ""
	}
	if strings.Contains(exp, "|") && len(match[1]) == 0 && len(match) > 2 {
		log.Println("Found secondary match")
		return match[2]
	}
	return match[1]
}
