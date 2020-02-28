package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsontest(`{"errorId":0,"taskId":1260524499}`)
	jsontest(`{"code": "000000","email": "cmoura322@susamicrosys.cf"}`)
}

func jsontest(data string) {
	jsonstring := []byte(data)
	// values := make(map[string]interface{})
	var response struct { // anti-captcha api response
		ErrorID int `json:"errorId"`
		TaskID  int `json:"taskId"`
	}
	err := json.Unmarshal(jsonstring, &values)
	if err != nil {
		fmt.Printf("Unmarshal error: %v", err)
	}
	fmt.Printf("converted data: %v\n", values)
}
