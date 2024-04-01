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

	input, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	fmt.Printf("Part 2 solution: %v\n", part2(input))
}

func part1(file *os.File) (ans int) {
	times := []int{}
	distances := []int{}

	fscanner := bufio.NewScanner(file)

	// Parse times
	fscanner.Scan()
	line := fscanner.Text()
	fields := strings.Fields(line)
	for _, f := range fields[1:] {
		t, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, t)
	}

	// Parse distances
	fscanner.Scan()
	line = fscanner.Text()
	fields = strings.Fields(line)
	for _, f := range fields[1:] {
		d, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		distances = append(distances, d)
	}

	ans = 1
	for i, t := range times {
		ans *= calcWinWays(t, distances[i])
	}

	return ans
}

func calcWinWays(time int, record int) (winWays int) {
	for holdTime := 0; holdTime <= time; holdTime++ {
		if time*holdTime-holdTime*holdTime > record {
			winWays++
		}
	}

	return winWays
}

func part2(file *os.File) (ans int) {
	var time, distance uint64

	fscanner := bufio.NewScanner(file)

	// Parse times
	fscanner.Scan()
	line := fscanner.Text()
	fields := strings.Fields(line)[1:]
	time, err := strconv.ParseUint(strings.Join(fields, ""), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Parse distances
	fscanner.Scan()
	line = fscanner.Text()
	fields = strings.Fields(line)[1:]
	distance, err = strconv.ParseUint(strings.Join(fields, ""), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return calcWinWays2(time, distance)
}

// Same as above but with unsigned integer to handle very large values
func calcWinWays2(time uint64, record uint64) (winWays int) {
	for holdTime := uint64(0); holdTime <= uint64(time); holdTime++ {
		if time*holdTime-holdTime*holdTime > record {
			winWays++
		}
	}

	return winWays
}
