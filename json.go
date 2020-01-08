package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsontest()
}

func jsontest() {
	jsonstring := []byte(`{"code": "000000","email": "cmoura322@susamicrosys.cf"}`)
	values := make(map[string]string)
	err := json.Unmarshal(jsonstring, &values)
	if err != nil {
		fmt.Printf("Unmarshal error: %v", err)
	}
	fmt.Printf("values['code'] : %v\n", values["code"])
	fmt.Printf("values['email']: %v\n", values["email"])
}
