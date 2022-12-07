package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type assignment string

func (a assignment) getIndices() (int, int) {
	xs := strings.Split(string(a), "-")
	start := parseInt(xs[0])
	end := parseInt(xs[1])
	return start, end
}

func (a assignment) contains(a2 assignment) bool {
	s, e := a.getIndices()
	s2, e2 := a2.getIndices()
	return s <= s2 && e >= e2
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
		xs := strings.Split(s, ",")
		a1, a2 := assignment(xs[0]), assignment(xs[1])
		if a1.contains(a2) || a2.contains(a1) {
			currentCal++
		}
	}

	fmt.Println("Program finished: The score would be", currentCal)
}
