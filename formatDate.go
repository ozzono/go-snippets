package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	nowAsSqlDate()
}

func formatDate(data string) string {
	slicedData := strings.Split(data, "/")
	return slicedData[2] + "-" + slicedData[1] + "-" + slicedData[0] + "T00:00:00Z"
}

func DateFormat(date string) (time.Time, error) {
	newdate, err := time.Parse("2006-01-02T15:04:05Z07:00", date)
	if err != nil {
		return newdate, err
	}
	return newdate, nil
}

func testCase1() {
	date := "13/10/2019"
	// newdate, err := DateFormat(date)
	newdate, err := DateFormat(formatDate(date))
	if err != nil {
		log.Printf("Dateformat error: %v", err)
		return
	}
	log.Printf("%v", newdate)
}

func nowAsSqlDate() {
	format := "2006-01-02" // 15:04:05"
	t := time.Now().Format(format)
	fmt.Printf("t: %v\n", t)
}
