package day5Part2

import (
	"advent_of_code/common"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type almanac struct {
	seeds seedsRange
	maps  map[string]seedMap
}

func (a *almanac) getSmallest(start string, end string) int {
	smallest := -1
	for i := 0; i < a.seeds.mapRange; i++ {
		seed := a.seeds.start + i
		currentMap := a.maps[start]
		v := seed
		for currentMap.destination != end {
			v = currentMap.getSeedOrSelf(v)
			currentMap = a.maps[currentMap.destination]
		}

		location := currentMap.getSeedOrSelf(v)
		if location < smallest || smallest == -1 {
			smallest = location
		}
	}
	return smallest
}

func (a *almanac) setSeed(seeds seedsRange) {
	a.seeds = seeds
}

func (a *almanac) addMap(sm seedMap) {
	if a.maps == nil {
		a.maps = map[string]seedMap{}
	}
	a.maps[sm.source] = sm
}

type seedsRange struct {
	start    int
	mapRange int
}

type rangeMapping struct {
	minSource      int
	minDestination int
	mapRange       int
}

type seedMap struct {
	source      string
	destination string
	mappings    []rangeMapping
}

func (s *seedMap) addRange(minSource int, minDestination int, mapRange int) {
	s.mappings = append(s.mappings, rangeMapping{minSource, minDestination, mapRange})
}

func mapStringsToInts(strs []string) []int {
	final := []int{}

	for _, numString := range strs {
		number, err := strconv.Atoi(strings.TrimSpace(numString))
		if err != nil {
			continue
		}

		final = append(final, number)
	}

	return final
}

func (s *seedMap) addFromString(str string) {
	nums := mapStringsToInts(strings.Split(str, " "))
	s.addRange(nums[1], nums[0], nums[2])
}

func (s *seedMap) getSeedOrSelf(source int) int {
	for _, mapping := range s.mappings {
		maxSource := mapping.minSource + mapping.mapRange
		if mapping.minSource <= source && maxSource >= source {
			offset := mapping.minDestination - mapping.minSource
			return source + offset
		}
	}
	return source
}

func newSeedMap(source string, destination string) seedMap {
	s := seedMap{}
	s.source = source
	s.destination = destination
	s.mappings = []rangeMapping{}
	return s
}

func parseSeedsString(seedString string) []seedsRange {
	seedsUntrimmed := strings.Split(strings.Split(seedString, ":")[1], " ")

	var seeds = []seedsRange{}

	seedNumbers := mapStringsToInts(seedsUntrimmed)

	maxSeedNumber := 0
	for _, seedNumber := range seedNumbers {
		if seedNumber > maxSeedNumber {
			maxSeedNumber = seedNumber
		}
	}

	maxNum := maxSeedNumber / 1000

	for i := 0; i < len(seedNumbers); i += 2 {
		iterations := seedNumbers[i+1] / maxNum
		remainder := seedNumbers[i+1] % maxNum

		for j := 0; j < iterations; j++ {
			seeds = append(seeds, seedsRange{seedNumbers[i] + maxNum*j, maxNum})
		}
		seeds = append(seeds, seedsRange{seedNumbers[i] + maxNum*iterations, remainder})
	}

	return seeds
}

func processStringsToAlmanac(mapStrings []string, currentMap *seedMap, newAlmanac *almanac) {
	for _, ms := range mapStrings {
		currentMap.addFromString(ms)
	}
	newAlmanac.addMap(*currentMap)
}

func almanacsFromStrings(str []string) ([]almanac, error) {

	almanacs := []almanac{}
	seeds := parseSeedsString(str[0])

	for _, seed := range seeds {
		var newAlmanac = almanac{}
		newAlmanac.setSeed(seed)

		currentMap := newSeedMap("unknown", "unknown")

		var mapStrings = []string{}
		for _, v := range str[2:] {
			if v == "" {
				// handle generated map, then make new
				processStringsToAlmanac(mapStrings, &currentMap, &newAlmanac)
				mapStrings = []string{}
				continue
			}

			if strings.Contains(v, "map:") {
				sourceToDest := strings.Split(v, " ")[0]
				splitSourceToDest := strings.Split(sourceToDest, "-")
				src := strings.TrimSpace(splitSourceToDest[0])
				dest := strings.TrimSpace(splitSourceToDest[2])
				currentMap = newSeedMap(src, dest)
				continue
			}

			mapStrings = append(mapStrings, v)
		}
		processStringsToAlmanac(mapStrings, &currentMap, &newAlmanac)
		almanacs = append(almanacs, newAlmanac)
	}
	return almanacs, nil
}

func getSmallestToChan(a almanac, c chan int) {
	c <- a.getSmallest("seed", "location")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func Part2() error {
	inputString := common.Input()
	// inputString := strings.Split(`seeds: 79 14 55 13

	// seed-to-soil map:
	// 50 98 2
	// 52 50 48

	// soil-to-fertilizer map:
	// 0 15 37
	// 37 52 2
	// 39 0 15

	// fertilizer-to-water map:
	// 49 53 8
	// 0 11 42
	// 42 0 7
	// 57 7 4

	// water-to-light map:
	// 88 18 7
	// 18 25 70

	// light-to-temperature map:
	// 45 77 23
	// 81 45 19
	// 68 64 13

	// temperature-to-humidity map:
	// 0 69 1
	// 1 0 69

	// humidity-to-location map:
	// 60 56 37
	// 56 93 4`, "\n")

	inputAlmanacs, err := almanacsFromStrings(inputString)
	if err != nil {
		return err
	}

	c := make(chan int, len(inputAlmanacs))

	smallest := -1

	remaining := len(inputAlmanacs)
	fmt.Println("Creating", len(inputAlmanacs), "goroutines")
  defer timeTrack(time.Now(), "calc")
	for _, a := range inputAlmanacs {
		go getSmallestToChan(a, c)
	}

	for range inputAlmanacs {
		val := <-c
		if val < smallest || smallest == -1 {
			smallest = val
		}

		remaining -= 1
		fmt.Println("Waiting for", remaining, "goroutines")
	}
	fmt.Println(smallest)
	return nil
}
