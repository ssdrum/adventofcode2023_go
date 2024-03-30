package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	//fmt.Println(part1(input))

	input, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	fmt.Println(part2(input))
}

type number struct {
	val, line, start, end int
}

type symbol struct {
	line, pos int
}

func (s symbol) toString() string {
	return strconv.Itoa(s.line) + " " + strconv.Itoa(s.pos)
}

func part1(file *os.File) (ans int) {
	fscanner := bufio.NewScanner(file)
	numbers := []number{}
	symbols := []symbol{}

	// Store position and value of all numbers and position of all symbols
	lineNum := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		start := -1
		for i, c := range line {
			//  Current character is a digit
			if unicode.IsDigit(c) {
				// Current character is the starting digit of a number
				if start == -1 {
					start = i
				}
				// Current character is the last digit of a number, at the end of the line
				if i == len(line)-1 {
					val, err := strconv.Atoi(line[start : i+1])
					if err != nil {
						log.Fatal(err)
					}
					numbers = append(numbers, number{val, lineNum, start, i - 1})
				}
			} else {
				//  Current character is not a digit
				if start != -1 {
					// Current character is the character next the ending digit of a number
					val, err := strconv.Atoi(line[start:i])
					if err != nil {
						log.Fatal(err)
					}
					// Append new number
					numbers = append(numbers, number{val, lineNum, start, i - 1})
					start = -1
				}
				if string(c) != "." {
					// Current character is a symbol
					symbols = append(symbols, symbol{lineNum, i})
				}
			}
		}

		lineNum++
	}

	// Find all numbers with an adjacent symbol and compute result
	for _, n := range numbers {
		for _, s := range symbols {
			// Check if number has an adjacent symbol on the same line
			if s.line == n.line && (s.pos == n.start-1 || s.pos == n.end+1) {
				ans += n.val
				break
			}
			// Check if number has an adjacent symbol on the lines above or below
			if (s.line == n.line+1 || s.line == n.line-1) && (n.start-1 <= s.pos && s.pos <= n.end+1) {
				ans += n.val
				break
			}
		}
	}

	return ans
}

func part2(file *os.File) (ans int) {
	fscanner := bufio.NewScanner(file)
	numbers := []number{}
	symbols := []symbol{}

	// Store position and value of all numbers and position of all symbols
	lineNum := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		start := -1
		for i, c := range line {
			//  Current character is a digit
			if unicode.IsDigit(c) {
				// Current character is the starting digit of a number
				if start == -1 {
					start = i
				}
				// Current character is the last digit of a number, at the end of the line
				if i == len(line)-1 {
					val, err := strconv.Atoi(line[start : i+1])
					if err != nil {
						log.Fatal(err)
					}
					numbers = append(numbers, number{val, lineNum, start, i - 1})
				}
			} else {
				//  Current character is not a digit
				if start != -1 {
					// Current character is the character next the ending digit of a number
					val, err := strconv.Atoi(line[start:i])
					if err != nil {
						log.Fatal(err)
					}
					// Append new number
					numbers = append(numbers, number{val, lineNum, start, i - 1})
					start = -1
				}
				if string(c) == "*" {
					// Current character is a symbol
					symbols = append(symbols, symbol{lineNum, i})
				}
			}
		}

		lineNum++
	}

	// Store a map of each gear with its adjacent numbers
	adjMap := map[string][]int{}

	// Find all gears and it's adjacent numbers
	for _, n := range numbers {
		for _, s := range symbols {
			// Check if number has an adjacent symbol on the same line
			if s.line == n.line && (s.pos == n.start-1 || s.pos == n.end+1) {
				adjMap[s.toString()] = append(adjMap[s.toString()], n.val)
			}
			// Check if number has an adjacent symbol on the lines above or below
			if (s.line == n.line+1 || s.line == n.line-1) && (n.start-1 <= s.pos && s.pos <= n.end+1) {
				adjMap[s.toString()] = append(adjMap[s.toString()], n.val)
			}
		}
	}

	// Calculate final result
	for _, values := range adjMap {
		if len(values) == 2 {
			ans += values[0] * values[1]
		}
	}

	return ans
}
