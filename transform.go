package main

import (
	"fmt"
	"strconv"
)

var Device string

func main() {
	Device = "vm"
	fmt.Println("meh")
	meh := 540 * 615 / 1081
	hm := fmt.Sprintf("%.0f", meh)
	// transform(540, 800)
	fmt.Printf("%s\n", hm)
}

func transform(x, y int) {
	if Device == "vm" {
		var err error
		fmt.Println("Using transformed coordinates for vm screen size")
		x, err = strconv.Atoi(fmt.Sprintf("%.0f", x*615/1080))
		if err != nil {
			fmt.Printf("Atoi err: %v", err)
		}
		fmt.Printf("Converted x: %v", x)
		y, err = strconv.Atoi(fmt.Sprintf("%.0f", y*799/1920))
		if err != nil {
			fmt.Printf("Atoi err: %v", err)
		}
		fmt.Printf("Converted y: %v", y)
	}
	xstring := strconv.Itoa(x)
	fmt.Printf("xstring: %v", xstring)
	ystring := strconv.Itoa(y)
	fmt.Printf("ystring: %v", ystring)
}
