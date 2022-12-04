package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	elfindex := 1
	currentCal := 0
	maxElfIndex := 0
	maxCal := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			if currentCal > maxCal {
				maxElfIndex = elfindex
				maxCal = currentCal
			}
			currentCal = 0
			elfindex++
		} else {
			c, e := strconv.Atoi(s)
			if e != nil {
				log.Fatal("Could not parse string as int: ", s)
			}
			currentCal += c
		}
	}

	fmt.Println("Program finished: The elf with the most calories is the", maxElfIndex, "one with ", maxCal, "calories")
}
