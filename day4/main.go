package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
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

		// Store winning numbers in a map
		winningNums := map[int]bool{}
		for _, token := range strings.Fields(string(matches[2])) {
			n, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			winningNums[n] = true
		}

		// Find number of matching winning numbers
		win := 0
		for _, token := range strings.Fields(string(matches[3])) {
			n, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			_, ok := winningNums[n]
			if ok {
				win += 1
			}
		}

		// Calculate result
		if win > 0 {
			ans += int(math.Pow(float64(2), float64(win-1)))
		}
	}

	return ans
}
