package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(dt2gotime("2020-01-17 14:46:13"))
}

func dt2gotime(sqldt string) time.Time {
	layout := "2006-01-02T15:04:05"
	sqldt = strings.Replace(sqldt, " ", "T", 1)
	t, err := time.Parse(layout, sqldt)
	if err != nil {
		fmt.Println(err)
		return time.Time{}
	}
	return t
}
