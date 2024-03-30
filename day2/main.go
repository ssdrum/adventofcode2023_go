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

func part1(input *os.File) (ans int) {
	maxCubes := map[string]int{
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
		gameId, _ := strconv.Atoi(strings.Split(tokens[0], " ")[1])
		// Extract handfuls of cubes picked
		cubesSets := strings.Split(tokens[1], ";")

		isGamePossible := true
		for _, set := range cubesSets {
			cubes := strings.Split(set, ",")

			for _, cube := range cubes {
				tokens = strings.Split(cube, " ")
				num, color := tokens[1], tokens[2]
				intNum, _ := strconv.Atoi(num)
				if intNum > maxCubes[color] {
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
