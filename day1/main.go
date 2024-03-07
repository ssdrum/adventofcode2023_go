package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	fmt.Println(part1(input))

	input, _ = os.Open("input.txt")
	defer input.Close()
	fmt.Println(part2(input))
}

func part1(input *os.File) int {
	tot := 0
	fscanner := bufio.NewScanner(input)

	for fscanner.Scan() {
		var first, last int
		// Read line
		line := fscanner.Text()
		// Find first digit
		for _, c := range line {
			if unicode.IsDigit(c) {
				first = int(c - '0')
				break
			}
		}
		// Find last digit
		for i := len(line) - 1; i >= 0; i-- {
			curr := rune(line[i])
			if unicode.IsDigit(curr) {
				last = int(curr - '0')
				break
			}
		}
		tot += first*10 + last
	}

	return tot
}

func part2(input *os.File) int {
	numsmap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	tot := 0
	fscanner := bufio.NewScanner(input)

	for fscanner.Scan() {
		line := fscanner.Text()
		var first, last int
		firstIndex, lastIndex := len(line), -1

		// Find first and last numbers as numbers
		for i, c := range line {
			if unicode.IsDigit(c) {
				if i < firstIndex {
					firstIndex = i
					first = int(c - '0')
				}
				if i > lastIndex {
					lastIndex = i
					last = int(c - '0')
				}
			}
		}

		// Find first and last numbers as strings
		for str, num := range numsmap {
			firstI := strings.Index(line, str)
			if firstI != -1 && firstI < firstIndex {
				firstIndex = firstI
				first = num
			}
			lastI := strings.LastIndex(line, str)
			if lastI != -1 && lastI > lastIndex {
				lastIndex = lastI
				last = num
			}
		}
		tot += first*10 + last
	}

	return tot
}
