package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type genericMap struct {
	dStart, srStart, rLen uint // Using uint since some input values exceed the range of regular ints
}

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

func part1(file *os.File) (ans uint) {
	fscanner := bufio.NewScanner(file)

	seeds := []uint{}
	seedToSoilMaps := []genericMap{}
	soilToFertilizerMaps := []genericMap{}
	fertilizerToWaterMaps := []genericMap{}
	waterToLightMaps := []genericMap{}
	lightToTemperatureMaps := []genericMap{}
	temperatureToHumidityMaps := []genericMap{}
	humidityToLocationMaps := []genericMap{}

	// Parse input file
	for fscanner.Scan() {
		line := fscanner.Text()
		switch {
		case strings.HasPrefix(line, "seeds:"):
			fields := strings.Fields(line)
			for _, s := range fields[1:] {
				n, err := strconv.ParseUint(s, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, uint(n))
			}
		case strings.HasPrefix(line, "seed-to-soil"):
			seedToSoilMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "soil-to-fertilizer"):
			soilToFertilizerMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "fertilizer-to-water"):
			fertilizerToWaterMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "water-to-light"):
			waterToLightMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "light-to-temperature"):
			lightToTemperatureMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "temperature-to-humidity"):
			temperatureToHumidityMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "humidity-to-location"):
			humidityToLocationMaps = storeMap(fscanner)
		}
	}

	// Find min location
	var minLocation uint = math.MaxUint
	for _, seed := range seeds {
		soil := sourceToDest(seed, seedToSoilMaps)
		fertilizer := sourceToDest(soil, soilToFertilizerMaps)
		water := sourceToDest(fertilizer, fertilizerToWaterMaps)
		light := sourceToDest(water, waterToLightMaps)
		temperature := sourceToDest(light, lightToTemperatureMaps)
		humidity := sourceToDest(temperature, temperatureToHumidityMaps)
		location := sourceToDest(humidity, humidityToLocationMaps)

		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func storeMap(fscanner *bufio.Scanner) (maps []genericMap) {
	for {
		fscanner.Scan()
		line := fscanner.Text()
		if line == "" {
			break
		}
		fields := strings.Fields(line)
		dStart, err := strconv.ParseUint(fields[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		srStart, err := strconv.ParseUint(fields[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		rLen, err := strconv.ParseUint(fields[2], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		maps = append(maps, genericMap{
			srStart: uint(srStart),
			dStart:  uint(dStart),
			rLen:    uint(rLen),
		})
	}

	return maps
}

// Matches number n from source map to destination map
func sourceToDest(n uint, maps []genericMap) (loc uint) {
	for _, m := range maps {
		if m.srStart <= n && n < m.srStart+m.rLen {
			return m.dStart + (n - m.srStart)
		}
	}

	return n
}

func part2(file *os.File) (ans uint) {
	fscanner := bufio.NewScanner(file)

	seeds := []uint{}
	seedToSoilMaps := []genericMap{}
	soilToFertilizerMaps := []genericMap{}
	fertilizerToWaterMaps := []genericMap{}
	waterToLightMaps := []genericMap{}
	lightToTemperatureMaps := []genericMap{}
	temperatureToHumidityMaps := []genericMap{}
	humidityToLocationMaps := []genericMap{}

	// Parse input file
	for fscanner.Scan() {
		line := fscanner.Text()
		switch {
		case strings.HasPrefix(line, "seeds:"):
			fields := strings.Fields(line)
			for _, s := range fields[1:] {
				n, err := strconv.ParseUint(s, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, uint(n))
			}
		case strings.HasPrefix(line, "seed-to-soil"):
			seedToSoilMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "soil-to-fertilizer"):
			soilToFertilizerMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "fertilizer-to-water"):
			fertilizerToWaterMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "water-to-light"):
			waterToLightMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "light-to-temperature"):
			lightToTemperatureMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "temperature-to-humidity"):
			temperatureToHumidityMaps = storeMap(fscanner)
		case strings.HasPrefix(line, "humidity-to-location"):
			humidityToLocationMaps = storeMap(fscanner)
		}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(len(seeds) / 2)

	// Find min location
	var minLocation uint = math.MaxUint
	for i := 0; i < len(seeds); i += 2 {
		start, length := seeds[i], seeds[i+1]

		// Use multi-threading
		go func(i int) {
			defer wg.Done()
			for seed := start; seed < start+length; seed++ {
				soil := sourceToDest(seed, seedToSoilMaps)
				fertilizer := sourceToDest(soil, soilToFertilizerMaps)
				water := sourceToDest(fertilizer, fertilizerToWaterMaps)
				light := sourceToDest(water, waterToLightMaps)
				temperature := sourceToDest(light, lightToTemperatureMaps)
				humidity := sourceToDest(temperature, temperatureToHumidityMaps)
				location := sourceToDest(humidity, humidityToLocationMaps)

				// Prevent race conditions when updating minLocation
				if location < minLocation {
					mu.Lock()
					minLocation = location
					mu.Unlock()
				}
			}
		}(i)
	}

	wg.Wait()

	return minLocation
}
