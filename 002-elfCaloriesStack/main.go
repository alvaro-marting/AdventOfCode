package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type topThree [3]int

func (t *topThree) check(i int) (bool, int) {
	maxPos := 2
	insert := false
	for n := 2; n >= 0; n-- {
		if t[n] >= i {
			break
		} else {
			insert = true
			maxPos = n
		}
	}

	return insert, maxPos
}

func (t *topThree) insert(index int, value int) {
	if index < 2 {
		for n := 1; n >= index; n-- {
			t[n+1] = t[n]
		}
	}
	t[index] = value
}

func (t *topThree) value() int {
	acc := 0
	for _, v := range t {
		acc += v
	}
	return acc
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	t := topThree{}
	currentCal := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			if insert, pos := t.check(currentCal); insert {
				t.insert(pos, currentCal)

			}
			currentCal = 0
		} else {
			c, e := strconv.Atoi(s)
			if e != nil {
				log.Fatal("Could not parse string as int: ", s)
			}
			currentCal += c
		}
	}

	fmt.Println("Program finished: The sum of the topThree is", t.value())
}
