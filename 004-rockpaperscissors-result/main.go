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
	result   result
}

type result string

func (r result) value() int {
	switch r {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		return 0
	}
}

func (r round) value() int {
	v := 0
	if r.result.value() == 3 {
		v = r.opponent.value()
	} else if r.result.value() == 6 {
		v = (r.opponent.value() % 3) + 1
	} else {
		if r.opponent.value() == 1 {
			v = 3
		} else if r.opponent.value() == 2 {
			v = 1
		} else {
			v = 2
		}
	}
	return r.result.value() + v
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
		r := round{opponent: input(xs[0]), result: result(xs[1])}
		currentCal += r.value()
		log.Println(currentCal, r.value(), r)
	}

	fmt.Println("Program finished: The score would be", currentCal)
}
