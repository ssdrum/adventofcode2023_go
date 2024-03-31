package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	fmt.Printf("Part 1 solution: %v\n", part1(input))

	//input, err = os.Open("input.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer input.Close()
	//fmt.Printf("Part 2 solution: %v\n", part2(input))
}

func part1(file *os.File) (ans int) {
	fscanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Card (\s*\d+): (\s*\d+(?: \s*\d+)*) \| (\s*\d+(?: \s*\d+)*)`)

	for fscanner.Scan() {
		line := fscanner.Text()
		matches := re.FindStringSubmatch(line)

		// Store winning numbers in a map for constant lookup time
		winningNums := map[string]bool{}
		for _, n := range strings.Fields(string(matches[2])) {
			winningNums[n] = true
		}

		// Find number of matching winning numbers
		wins := 0
		for _, n := range strings.Fields(string(matches[3])) {
			_, ok := winningNums[n]
			if ok {
				wins++
			}
		}

		// Calculate result
		if wins > 0 {
			ans += int(math.Pow(float64(2), float64(wins-1)))
		}
	}

	return ans
}
