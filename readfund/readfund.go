package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Starting to read funds and dates")
	// xml, err := getXml("remote-20191024-161934.xml")
	xml, err := getXml("fail.xml")
	if err != nil {
		fmt.Printf(`err: %v`, err)
	}
	match, err := readFund(`Data de pagamento.*"(?:^|[^Data de entrada/])?.*index="3".*content-desc="(\d*?.?\d+,\d{1,2})"|Data de pagamento.*"(?:^|[^Data de entrada/])?.*index="3" text="(\d*?.?\d+,\d{1,2})"`, xml)
	if err != nil {
		fmt.Printf(`err: %v`, err)
	}
	if strings.Contains(match[0], "Data de entrada") {
		fmt.Println("Ahá")
	}
}

func readFund(exp, xml string) ([]string, error) {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(xml)
	if len(match) <= 1 {
		fmt.Println("Unable to find match for card fund")
	}
	if len(match) > 1 {
		for i, _ := range match {
			fmt.Printf("match[%v]: %#v\n\n", i, match[i])
		}
		return match, nil
	}
	return match, errors.New("Regex not found")
}

func getXml(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
