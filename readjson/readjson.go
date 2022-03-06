package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	data, err := readJson("sample.json")
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	log.Println(data)
	for key := range data {
		log.Printf("data[%s]: %v", key, data[key])
	}
}

func readJson(path string) (map[string]interface{}, error) {
	var data map[string]interface{}
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	log.Println(string(jsonFile))
	err = json.Unmarshal([]byte(jsonFile), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
