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
	//fmt.Println(part_1(input))

	input, _ = os.Open("input.txt")
	defer input.Close()
	fmt.Println(part_2(input))
}

func part_1(input *os.File) int {
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

func part_2(input *os.File) int {
	numsmap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}
	tot := 0
	fscanner := bufio.NewScanner(input)

	for fscanner.Scan() {
		line := fscanner.Text()
		var first, last int
		first_index, last_index := len(line), -1

		// Find first and last numbers as numbers
		for i, c := range line {
			if unicode.IsDigit(c) {
				if i < first_index {
					first_index = i
					first = int(c - '0')
				}
				if i > last_index {
					last_index = i
					last = int(c - '0')
				}
			}
		}

		// Find first and last numbers as strings
		for str, num := range numsmap {
			first_i := strings.Index(line, str)
			if first_i != -1 && first_i < first_index {
				first_index = first_i
				first = num
			}
			last_i := strings.LastIndex(line, str)
			if last_i != -1 && last_i > last_index {
				last_index = last_i
				last = num
			}
		}
		tot += first*10 + last
	}

	return tot
}
