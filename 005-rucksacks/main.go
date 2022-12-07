package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rucksack string

type compartment string

func (r rucksack) getCompartments() (compartment, compartment) {
	return compartment(r[:(len(r) / 2)]), compartment(r[len(r)/2:])
}

func (r rucksack) getRepeatedTypePriority() int {
	a, b := r.getCompartments()
	for _, v := range a {
		c := strings.Contains(string(b), string(v))
		if c {
			return getValue(v)
		}
	}

	return 0
}

func getValue(r rune) int {
	b := byte(r)
	if b >= 97 && b <= 122 {
		return int(b % 96)
	} else if b >= 65 && b <= 90 {
		return int(b%64 + 26)
	}
	return 0
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
		r := rucksack(s)
		currentCal += r.getRepeatedTypePriority()
	}

	fmt.Println("Program finished: The score would be", currentCal)
}
