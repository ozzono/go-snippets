package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mega struct {
	universe []int
	bet      []int
}

func main() {
	var mega Mega
	mega.DiceRoll(6)
}

func (mega *Mega) populate(size int) {
	for i := 0; i < size; i++ {
		mega.universe = append(mega.universe, i+1)
	}
}

func (mega *Mega) killIndex(i int) {
	mega.universe[i] = mega.universe[len(mega.universe)-1]
	mega.universe = mega.universe[:len(mega.universe)-1]
}

func (mega *Mega) DiceRoll(size int) {
	mega.populate(60)
	for i := 0; i < size; i++ {
		index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(mega.universe))
		mega.bet = append(mega.bet, mega.universe[index])
		mega.killIndex(index)
	}
	fmt.Println(mega.bet)
}
