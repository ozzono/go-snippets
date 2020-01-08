package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var Namespace string

func main() {
	Namespace = "sdx"
	t1 := time.Now()
	triggerFetchMail()
	t2 := time.Now()
	fmt.Printf("time diff: %v\n", time.Time{}.Add(t2.Sub(t1)).Format("15:04:05"))
}

func triggerFetchMail() bool {
	fmt.Println("Starting triggerFetchMail")
	// url := "https://script.google.com/macros/s/AKfycbwZqmMOxgbI6R-pr99YVqGVvqkaqt15pUBJ0JtEAu01Qp2w_i5D/exec"
	url := "https://script.google.com/macros/s/AKfycbysFrqnSi8v3l9YeN_ktVyYHG0jHI0lAwufKbXS4A/exec"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("NewRequest error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("DefaultClient.Do error: %v", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		fmt.Println("Returned code 200")
		output, err := strconv.ParseBool(string(body))
		if err != nil {
			fmt.Printf("ParseBool error: %v", err)
			return false
		}
		return output
	}
	return false
}
