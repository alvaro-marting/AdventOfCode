package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Stack is a basic LIFO stack that resizes as needed.
type Stack []string

// Push adds an element to the stack.
func (s *Stack) Push(v string) {
	if len(v) > 0 {
		*s = append(*s, v)
	}
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() string {
	if s.Empty() {
		return ""
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Empty returns true if the stack is empty, false otherwise.
func (s *Stack) Empty() bool {
	return len(*s) == 0
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
	// Create a slice of stacks
	columns := make([]Stack, 9)
	columnsRead := false
	r := regexp.MustCompile(`\[(\w)\]`)

	r2 := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for scanner.Scan() {
		s := scanner.Text()
		if !columnsRead {
			if len(s) == 0 {
				columnsRead = true
				for _, v := range columns {
					for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
						v[i], v[j] = v[j], v[i]
					}
					fmt.Printf(v[len(v)-1])
				}
				fmt.Println()
			} else {
				for i, _ := range columns {
					sect := string(s[i*4 : ((i+1)*4 - 1)])
					if r.MatchString(sect) {
						columns[i].Push(r.FindStringSubmatch(sect)[1])
					}
				}
			}
		} else {
			matches := r2.FindStringSubmatch(s)
			n, from, to := parseInt(matches[1]), parseInt(matches[2]), parseInt(matches[3])
			f := &columns[from-1]
			t := &columns[to-1]
			for i := 0; i < n; i++ {
				t.Push(f.Pop())
			}
		}
	}
	for _, v := range columns {
		fmt.Printf(v[len(v)-1])
	}
}
