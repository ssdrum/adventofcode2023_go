package main

import (
	"bufio"
	"fmt"
	"log"
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

	input, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	fmt.Printf("Part 2 solution: %v\n", part2(input))
}

func part1(input *os.File) (ans int) {
	minCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	fscanner := bufio.NewScanner(input)

	for fscanner.Scan() {
		// Read line
		line := fscanner.Text()
		tokens := strings.Split(line, ":")
		// Extract game ID
		gameId, err := strconv.Atoi(strings.Split(tokens[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		// Extract handfuls of cubes picked
		cubesSets := strings.Split(tokens[1], ";")

		isGamePossible := true
		for _, set := range cubesSets {
			cubes := strings.Split(set, ",")

			for _, cube := range cubes {
				tokens = strings.Split(cube, " ")
				num, color := tokens[1], tokens[2]
				intNum, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}

				if intNum > minCubes[color] {
					isGamePossible = false
					break
				}
			}
		}

		if isGamePossible {
			ans += gameId
		}
	}

	return ans
}

func part2(input *os.File) (ans int) {
	fscanner := bufio.NewScanner(input)
	re := regexp.MustCompile(`(\d+)\s*(green|blue|red)?`)

	for fscanner.Scan() {
		// Read line
		line := fscanner.Text()
		tokens := strings.Split(line, ":")

		// Extract handfuls of cubes picked
		cubesSets := strings.Split(tokens[1], ";")
		// Find minimum number of cubes needed per game
		minRed, minGreen, minBlue := 0, 0, 0
		for _, set := range cubesSets {
			// Extract all matches in cube set
			matches := re.FindAllStringSubmatch(set, -1)
			for _, match := range matches {
				// Extract values
				number := match[1]
				intNumber, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				color := match[2]

				switch {
				case color == "red" && intNumber > minRed:
					minRed = intNumber
				case color == "green" && intNumber > minGreen:
					minGreen = intNumber
				case color == "blue" && intNumber > minBlue:
					minBlue = intNumber
				}
			}
		}

		// Add power set of cubes to answer
		ans += minRed * minGreen * minBlue
	}

	return ans
}
