package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", triggerFetchMail())
}

func triggerFetchMail() bool {
	url := "https://script.google.com/macros/s/AKfycbwZqmMOxgbI6R-pr99YVqGVvqkaqt15pUBJ0JtEAu01Qp2w_i5D/exec"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("NewRequest error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("DefaultClient.Do error: %v", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		output, err := strconv.ParseBool(string(body))
		if err != nil {
			log.Printf("ParseBool error: %v", err)
			return false
		}
		return output
	}
	return false
}
