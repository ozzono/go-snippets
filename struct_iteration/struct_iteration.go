package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := "foo"
	b := "bar"

	x := struct {
		Foo *string
		Bar *string
	}{&f, &b}
	iterate(x)

	/*
		name: Foo, value: foo (*string)
		name: Bar, value: bar (*string)
	*/
}

func iterate(input interface{}) {
	v := reflect.ValueOf(input)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("name: %+v, value: %+v (%T)\n",
			v.Type().Field(i).Name, // Name attribute gives us the struct's key
			v.Field(i).Elem(),      // Elem() dereferences the pointer value
			v.Field(i).Interface()) // Interface() provides memory address of the value
	}
}
