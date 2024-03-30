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

	fmt.Println(part1(input))
}

type number struct {
	val, line, start, stop int
}

type symbol struct {
	line, pos int
}

func part1(file *os.File) (ans int) {
	fscanner := bufio.NewScanner(file)
	numbers := []number{}
	symbols := []symbol{}

	lineNum := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		currNum := ""
		start := 0
		for i, c := range line {
			//  Current character is a digit
			if unicode.IsDigit(c) {
				// Current character is the starting digit of a number
				if currNum == "" {
					start = i
				}
				currNum += string(c)
			} else {
				//  Current character is not a digit
				if currNum != "" {
					// Current character is the character next the ending digit of a number
					val, err := strconv.Atoi(currNum)
					if err != nil {
						log.Fatalf("cannot convert %v to int", val)
					}
					// Append new number
					numbers = append(numbers, number{val, lineNum, start, i - 1})
					// Reset variables
					currNum = ""
					start = 0
				} else {
					// Current character is not a number
					if string(c) != "." {
						// Current character is a symbol
						symbols = append(symbols, symbol{lineNum, i})
					}
				}
			}
		}
		lineNum++
	}

	for _, n := range numbers {
		fmt.Println(n)
	}

	for _, s := range symbols {
		fmt.Println(s)
	}

	return ans
}
