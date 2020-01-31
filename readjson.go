package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type SQL struct { //sql arguments
	User string
	Pass string
	Host string
	Db   string
}

func main() {
	json, err := readJson("samples/origin.json")
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("%#v\n", json)
}

func readJson(path string) (SQL, error) {
	var accessSql SQL
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return accessSql, err
	}
	err = json.Unmarshal([]byte(jsonFile), &accessSql)
	if err != nil {
		return accessSql, err
	}
	return accessSql, nil
}

func tapScreen(x, y, delay int) {

	xstring := strconv.Itoa(x)
	ystring := strconv.Itoa(y)
	cmd("adb shell input tap "+xstring+" "+ystring, delay, false, true)
}
