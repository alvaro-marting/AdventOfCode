package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type round struct {
	opponent input
	you      input
}

func (r round) value() int {
	v := 0
	if r.you.value()%3 == (r.opponent.value()+1)%3 {
		v = 6
	} else if r.you.value() == r.opponent.value() {
		v = 3
	}
	return r.you.value() + v
}

type input string

func (i input) value() int {
	switch {
	case i == "A" || i == "X":
		return 1
	case i == "B" || i == "Y":
		return 2
	case i == "C" || i == "Z":
		return 3
	default:
		return 0
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	currentCal := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		xs := strings.Split(s, " ")
		if len(xs) != 2 {
			log.Fatal("MORE THAN 2 INPUTS READ")
		}
		r := round{opponent: input(xs[0]), you: input(xs[1])}
		currentCal += r.value()
		log.Println(currentCal, r.value(), r)
	}

	fmt.Println("Program finished: The score would be", currentCal)
}
