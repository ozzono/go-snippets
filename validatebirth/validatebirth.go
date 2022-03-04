package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	ValidadeBirthDate(newdate())
}

func ValidadeBirthDate(date string) bool {
	log.Printf("Validating birthdate: %s", date)
	splitdata := strings.Split(date, "/")
	if len(splitdata) < 3 {
		log.Printf("Invalid birthdate: %s", date)
		return false
	}
	day, err := strconv.Atoi(splitdata[0])
	if err != nil {
		log.Printf("Invalid birthdate: %s\nerr: %v", date, err)
		return false
	}
	month, err := strconv.Atoi(splitdata[1])
	if err != nil {
		log.Printf("Invalid birthdate: %s\nerr: %v", date, err)
		return false
	}
	year, err := strconv.Atoi(splitdata[2])
	if err != nil {
		log.Printf("Invalid birthdate: %s\nerr: %v", date, err)
		return false
	}
	validday := day <= 31 && day >= 1
	validmonth := month <= 12 && month >= 1
	validyear := year >= 1900
	if !validday || !validmonth || !validyear {
		log.Printf("Invalid birthdate: %s", date)
		return false
	}
	fmt.Printf("Birthdate %s is valid\n", date)
	return true
}

func randn(i int) string {
	i = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
	if i == 0 {
		i = 1
	}
	return fmt.Sprintf("%02d", i)
}

func newdate() string {
	return strings.Join([]string{randn(29), randn(12), year()}, "/")
}

func year() string {
	if randn(2) == "1" {
		return "20" + randn(20)
	}
	return "19" + randn(99)
}
