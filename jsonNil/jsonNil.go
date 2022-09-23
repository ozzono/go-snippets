package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type v struct {
		S string      `json:"s"`
		V interface{} `json:"v"`
	}

	val := v{S: "s", V: nil}
	data, err := json.MarshalIndent(val, "", "	")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
