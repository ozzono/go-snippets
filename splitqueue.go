// This code takes a defined amount of parallel processing queues
// with defined amount of processes and splits among a defined
// amount of processing units trying to balance the load as even
// as possible

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Count struct {
	Queue []int
	Count int
}

var (
	queuecount int
	splitcount int
)

func init() {
	flag.IntVar(&queuecount, "q", 10, "Max parallel identified count")
	flag.IntVar(&splitcount, "s", 5, "Desired parallel count")
}

func main() {
	flag.Parse()
	count := queueSplit(Sort(countFlow(queuecount)))
	for i, _ := range count {
		fmt.Printf("Count: %03d - Queue: %v\n", count[i].Count, count[i].Queue)
	}
}

func randstuff(i int) []int {
	n := ((rand.New(rand.NewSource(time.Now().UnixNano())).Int() + 1) % queuecount) * 10
	if n != 0 {
		n = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
	}
	output := []int{}
	for i := 0; i < n; i++ {
		output = append(output, i)
	}
	return output
}

func countFlow(queuecount int) []Count {
	fmt.Printf("\nStarting count for %d rows\n", queuecount)
	balance := []Count{}
	count := []int{}
	for i := 0; i < queuecount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			count = randstuff(i)
			balance = append(balance, Count{[]int{i}, len(count)})
		}(i)
	}

	wg.Wait()
	return balance
}

func queueSplit(count []Count) []Count {
	split := []Count{}
	index := 0
	total := 0
	for len(split) < splitcount {
		split = append(split, Count{})
	}
	for i := 0; i < len(count); i++ {
		index = next(split)
		split[index].Count += count[i].Count
		split[index].Queue = append(split[index].Queue, count[i].Queue[0])
	}

	split = Sort(split)
	for j, _ := range split {
		total += split[j].Count
	}
	fmt.Printf("Total count: %d\n\n", total)
	return split
}

func Sort(count []Count) []Count {
	for i := len(count); i > 0; i-- {
		for j := 1; j < i; j++ {
			if count[j-1].Count < count[j].Count {
				tmp := count[j]
				count[j] = count[j-1]
				count[j-1] = tmp
			}
		}
	}
	return count
}

func next(split []Count) int {
	index := 0
	for i, _ := range split {
		if split[i].Count == 0 {
			return i
		} else if split[i].Count < split[index].Count {
			index = i
		}
	}
	return index
}
