// This package calculates the basic statistics from
// a log file with the given format: YYYY-MM-DD HH:mm:ss: SS
// where SS is this case is the total amount of seconds
// taken for each loop of another application

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	today string
	file  string
)

type AllNumbers struct {
	Average   float64
	Deviation float64
	Max       float64
	Min       float64
	Total     int
}

func init() {
	flag.StringVar(&file, "f", "", "Interval file path")
}

func main() {
	flag.Parse()
	today = time.Now().Format("2006-01-02")
	file, err := readlog(file)
	if err != nil {
		fmt.Printf("readlog err: %v", err)
	}
	allnumbers := stddeviation(slicelog(file))
	fmt.Printf("Todays standard deviation: -> %.2f\n", allnumbers.Deviation)
	fmt.Printf("Colection average: ---------> %.2f\n", allnumbers.Average)
	fmt.Printf("Max interval: --------------> %.2f\n", allnumbers.Max)
	fmt.Printf("Min interval: --------------> %.2f\n", allnumbers.Min)
	fmt.Printf("Verified registrations: ----> %d\n", allnumbers.Total)
}

func readlog(path string) (string, error) {
	logfile, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(logfile), nil
}

func slicelog(log string) []float64 {
	output := []float64{}
	for _, row := range strings.Split(log, "\n") {
		if strings.Contains(row, today) {
			item, err := strconv.ParseFloat(strings.Split(row, ": ")[1], 64)
			if err != nil {
				fmt.Println(err)
			}
			output = append(output, item)
		}
	}
	return output
}

func average(interval []float64) (float64, float64, float64) {
	var sum, max, min float64
	start := true
	for i, _ := range interval {
		sum += interval[i]
		if interval[i] > max {
			max = interval[i]
		}
		if start {
			min = max
			start = !start
		}
		if interval[i] < min {
			min = interval[i]
		}
	}
	return sum / float64(len(interval)), max, min
}

func stddeviation(interval []float64) AllNumbers {
	average, max, min := average(interval)
	var sqrsum float64
	for i, _ := range interval {
		diff := interval[i] - average
		sqrsum += diff * diff
	}
	sqrsum = sqrsum / float64(len(interval))
	sqrsum = math.Sqrt(sqrsum)
	return AllNumbers{
		average,
		sqrsum,
		max,
		min,
		len(interval),
	}
}
