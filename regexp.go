package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	filename    string
	exp         string
	inputString string
)

func init() {
	flag.StringVar(&filename, "f", "", "File to be read")
	flag.StringVar(&exp, "e", "", "Regular expression")
	flag.StringVar(&inputString, "s", "", "Input string on terminal")
}

func main() {
	flag.Parse()
	start()
}

func start() {
	file := readFile(filename)
	if strings.Contains(file, "LOCK TABLES `funcionarios` WRITE;") {
		brokefile := strings.Split(file, "\n")
		for _, line := range brokefile {
			if strings.Contains(line, "INSERT INTO `funcionarios` VALUES") {
				insert := myregexp("INSERT INTO `funcionarios` VALUES \\((.*)\\);", line)
				slicedInsert := strings.Split(insert, ",")

				slicedInsert[2] = keepquotes(slicedInsert[2], firstName(slicedInsert[2])+" *****")    // [ 2]: employee name
				slicedInsert[4] = keepquotes(slicedInsert[4], slicedInsert[4][:3]+"****")             // [ 4]: rg
				slicedInsert[8] = keepquotes(slicedInsert[8], maskDate(slicedInsert[8]))              // [ 8]: birthday
				slicedInsert[11] = keepquotes(slicedInsert[11], firstName(slicedInsert[11])+" *****") // [11]: employee mothers name
				slicedInsert[12] = keepquotes(slicedInsert[12], "")                                   // [12]: address
				slicedInsert[13] = keepquotes(slicedInsert[13], "")                                   // [13]: address number
				slicedInsert[14] = keepquotes(slicedInsert[14], "")                                   // [14]: address complement

				output := strings.Join(slicedInsert, ",")
				newline := strings.Replace(line, insert, output, 1)
				file = strings.Replace(file, line, newline, 1)
			}
		}
		dumpToFile("compare.sql", file)
	}
}

func readFile(filename string) string {
	filebyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error fetching file: %v\n", err)
		return ""
	}
	file := string(filebyte)
	if len(strings.Split(file, "\n")) > 10 {
		fmt.Println("\nHEAD")
		for _, item := range strings.Split(file, "\n")[:5] {
			fmt.Printf("%v\n", item)
		}
		fmt.Printf("[%d lines]", len(strings.Split(file, "\n"))-10)
		for _, item := range strings.Split(file, "\n")[len(strings.Split(file, "\n"))-5:] {
			fmt.Printf("\n%v", item)
		}
		fmt.Println("TAIL\n")
	} else {
		fmt.Println(file)
	}
	return file
}

func myregexp(exp, text string) string {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		log.Println("Unable to find match for regexp")
		return ""
	}
	return match[1]
}
func firstName(fullname string) string {
	for i, item := range fullname {
		if string(item) == " " {
			return fullname[:i]
		}
	}
	return fullname
}

func maskDate(input string) string {
	index := len(input) - 1
	for i := index; i >= 0; i-- {
		if string(input[i]) == "/" {
			index = i
			break
		}
	}
	return input[:index+1] + "@@@@"
}

func keepquotes(quoted, input string) string {
	if strings.Contains(quoted, "'") {
		input = strings.Replace(input, "'", "", -1)
		return fmt.Sprintf("'%s'", input)
	}
	return input
}

func dumpToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data+"\n")
	if err != nil {
		return err
	}
	return file.Sync()
}
