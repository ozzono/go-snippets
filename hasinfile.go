package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"beneficiofacil.gopkg.net/site/util"
)

var (
	filepath   string
	searchitem string
)

func init() {
	flag.StringVar(&filepath, "f", "", "Sets the file path of the text to be evaluated")
	flag.StringVar(&searchitem, "i", "", "Sets the item to be searched withing the file")
}

func main() {
	flag.Parse()
	if len(filepath) == 0 || filepath == "" {
		fmt.Printf("Invalid value for -f; cannot be empty")
		return
	}

	if len(searchitem) == 0 || searchitem == "" {
		fmt.Printf("Invalid value for -i; cannot be empty")
		return
	}
	hasInFile()
}

func hasInFile() {
	filedata, err := readFile(filepath)
	if err != nil {
		fmt.Printf("readFile err:\n%v", err)
		return
	}
	filedata, err = util.ToUTF8(strings.ToLower(filedata), true)
	if err != nil {
		fmt.Printf("Unable to normalize filedata:\n%v", err)
		return
	}
	searchitem, err = util.ToUTF8(strings.ToLower(searchitem), true)
	if err != nil {
		fmt.Printf("Unable to normalize searchitem:\n%v", err)
		return
	}
	if strings.Contains(filedata, searchitem) {
		fmt.Printf("[Success] The item '%s' was found within the file '%s'\n", searchitem, filepath)
		return
	}
	fmt.Printf("[Fail] The item '%s' was not found within the file '%s'\n", searchitem, filepath)
}

func readFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
