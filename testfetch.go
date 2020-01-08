package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://script.google.com/a/macros/beneficiofacil.com.br/s/AKfycbyAScA3iyLXbiqCs-9Es8dQQoSi_dreHNUANxaC-g/exec?teste=coisa%20linda%20de%20deus"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("NewRequest err: ", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("DefaultClient.Do err: ", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ReadAll err: ", err)
	}

	// fmt.Println(res)
	fmt.Println(string(body))
}
