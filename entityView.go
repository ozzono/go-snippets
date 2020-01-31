package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"google.golang.org/appengine/datastore"
)

var (
	e    entity
	path string
)

func init() {
	flag.StringVar(&path, "f", "", "")
}

type entity struct {
	Key   *datastore.Key
	Props []property
}

type property struct {
	Name     string
	Value    interface{}
	NoIndex  bool
	Multiple bool
}

func main() {
	flag.Parse()
	e, err := getEntity(path)
	if err != nil {
		fmt.Println("path: ", path)
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", e[0].Key)
	fmt.Printf("%#v\n", e[0].Props)

	printEntity(e)
}

func getEntity(path string) ([]entity, error) {
	var e []entity
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return e, err
	}
	fmt.Printf("%v", string(jsonFile))
	err = json.Unmarshal(jsonFile, &e)
	if err != nil {
		return e, err
	}
	return e, nil
}

func printEntity(e []entity) {
	if len(e) > 0 {
		for i := range e {
			if len(e[i].Props) > 0 {
				for j := range e[i].Props {
					fmt.Printf("Prop[%d]\n", j)
					fmt.Printf("Name: -----> %v\n", e[i].Props[j].Name)
					fmt.Printf("Value: ----> %v\n", e[i].Props[j].Value)
					fmt.Printf("NoIndex: --> %v\n", e[i].Props[j].NoIndex)
					fmt.Printf("Multiple: -> %v\n", e[i].Props[j].Multiple)
					fmt.Println("####")
				}
				continue
			}
			fmt.Printf("e[%d].Props is empty\n", i)
		}
	}
	fmt.Println("e is empty")
}
