package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type cid struct {
	Origin  string `json:"origin"`
	Version string `json:"version"`
	Data    []Data `json:"data"`
	
}

type Data struct {
	Cid  string `json:"cid"`
	Desc string `json:"desc"`
}

func main() {
	cidlist, err := readJson("/home/hugo/Projects/ifbra-app/src/assets/json/cid10.json")
	if err != nil {
		log.Printf("readjson err: %v", err)
	}
	cidlist.Data = parseCid(cidlist.Data)
	cidByte, err := json.MarshalIndent(cidlist, "", "	")
	if err != nil {
		log.Printf("MarshalIndent err: %v", err)
	}
	WriteToFile("newcid10.json", string(cidByte))
}

func readJson(path string) (cid, error) {
	cidlist := cid{}
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return cidlist, err
	}
	err = json.Unmarshal([]byte(jsonFile), &cidlist)
	if err != nil {
		return cid{}, err
	}
	return cidlist, nil
}

func parseCid(list []Data) []Data {
	output := []Data{}
	for i := range list {
		if strings.Contains(list[i].Cid, ", ") {
			for _, item := range strings.Split(list[i].Cid, ", ") {
				output = append(output, Data{Cid: item, Desc: list[i].Desc})
			}
			fmt.Printf("Parsing %v\n", list[i].Cid)
		} else {
			output = append(output, list[i])
		}
	}
	return output
}

func WriteToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
