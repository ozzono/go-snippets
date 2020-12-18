package main

import (
	"flag"
	"fmt"
)

func main() {
	interest := flag.Float64("interest", 0.002, "Set value for constant increase float")
	initial := flag.Float64("initial", 1, "Set value for constant increase float")
	period := flag.Int("period", 12, "Set value for constant increase")
	iValue := *initial
	instalment := iValue / float64(*period)
	flag.Parse()
	for i := 1; i <= *period; i++ {
		rendered := iValue * (*interest)
		iValue = iValue - instalment + rendered
	}
	fmt.Println("iValue ", iValue)
}
