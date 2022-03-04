package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

func main() {
	rawcookie, err := readFile("samples/cookie/cookies.txt")
	if err != nil {
		fmt.Println(err)
	}
	output := []string{}
	for _, item := range strings.Split(rawcookie, "\n") {
		if strings.Contains(item, "riocardmais") && neededCookies(item) {
			row := strings.Split(item, "	")
			output = append(output, fmt.Sprintf("%v=%v", row[len(row)-2], url.QueryEscape(row[len(row)-1])))
		}
	}
	fmt.Println(strings.Join(output, "; "))
}

func readFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func neededCookies(input string) bool {
	cookies := []string{
		"JSESSIONID",
		"visid",
		"incap",
	}
	for i := range cookies {
		if strings.Contains(input, cookies[i]) {
			return true
		}
	}
	return false
}
