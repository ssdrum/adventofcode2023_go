package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// input, err = os.Open("input.txt")
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// defer input.Close()
	// fmt.Printf("Part 2 solution: %v\n", part2(input))
}

type hand struct {
	cards string
	bet   int
}

type hands []hand

func (h hands) Len() int {
	return len(h)
}

// Implements the Less method required by sort.Interface
func (h hands) Less(a, b int) bool {
	countA := countCards(h[a])
	countB := countCards(h[b])

	// Sort by rank
	if calcRank(maxCount(countA), len(countA)) < calcRank(maxCount(countB), len(countB)) {
		return true
	}
	if calcRank(maxCount(countA), len(countA)) > calcRank(maxCount(countB), len(countB)) {
		return false
	}
	if greatestByLabel(h[a], h[b]) == h[b] {
		return true
	}
	return false
}

// Implement the Swap method required by sort.Interface
func (h hands) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

// Returns a map holding the number of occurrances of each card
func countCards(h hand) map[string]int {
	count := map[string]int{}
	for i := 0; i < 5; i++ {
		count[string(h.cards[i])]++
	}

	return count
}

// Finds the maximum number of occurrances for any card in the hand
func maxCount(c map[string]int) int {
	maxCount := 0
	for _, value := range c {
		if value > maxCount {
			maxCount = value
		}
	}

	return maxCount
}

// Assigns a rank value in range [0, 7] with 0 == no pairs and 7 == 5 of a kind
func calcRank(maxCount, countLen int) int {
	switch {
	// Two pair
	case maxCount == 2 && countLen == 3:
		return 3
		// High card or one pair
	case maxCount == 1 || maxCount == 2:
		return maxCount
		// Three of a kind
	case maxCount == 3 && countLen == 3:
		return 4
		// Fullhouse
	case maxCount == 3 && countLen == 2:
		return 5
		// Four or five of a kind
	default:
		return maxCount + 2
	}
}

// Returns true if a < b by label. Assumes that a != b
func greatestByLabel(a, b hand) hand {
	ranks := map[string]int{
		"A": 12,
		"K": 11,
		"Q": 10,
		"J": 9,
		"T": 8,
		"9": 7,
		"8": 6,
		"7": 5,
		"6": 4,
		"5": 3,
		"4": 2,
		"3": 1,
		"2": 0,
	}

	for i := 0; i < 5; i++ {
		if ranks[string(a.cards[i])] > ranks[string(b.cards[i])] {
			return a
		}
		if ranks[string(a.cards[i])] < ranks[string(b.cards[i])] {
			return b
		}
	}
	// Unreachable assuming that the input does not contain equal hands
	log.Fatal("Duplicate hand found")
	return a
}

func part1(file *os.File) (ans int) {
	fscanner := bufio.NewScanner(file)
	hands := hands{}

	// Create slice of hands
	for fscanner.Scan() {
		fields := strings.Fields(fscanner.Text())
		// Create struct and assign values
		h := hand{}
		h.cards = fields[0]
		bet, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		h.bet = bet
		hands = append(hands, h)
	}

	// Sort slice of hands and calculate result
	sort.Sort(hands)
	for i, h := range hands {
		ans += h.bet * (i + 1)
	}

	return ans
}
