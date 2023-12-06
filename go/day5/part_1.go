package day5

import (
	"advent_of_code/common"
	"fmt"
	"strconv"
	"strings"
)

type almanac struct {
	seeds []int
	maps  map[string]seedMap
}

func (a *almanac) getSmallest(start string, end string) int {
	smallest := -1
	for _, seed := range a.seeds {
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

func (a *almanac) setSeed(seeds []int) {
	a.seeds = seeds
}

func (a *almanac) addMap(sm seedMap) {
	if a.maps == nil {
		a.maps = map[string]seedMap{}
	}
	a.maps[sm.source] = sm
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

func parseSeedString(seedString string) []int {
	seedsUntrimmed := strings.Split(strings.Split(seedString, ":")[1], " ")

	var seeds = []int{}

	for _, v := range seedsUntrimmed {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		seeds = append(seeds, number)
	}

	return seeds
}

func processStringsToAlmanac(mapStrings []string, currentMap *seedMap, newAlmanac *almanac) {
	for _, ms := range mapStrings {
		currentMap.addFromString(ms)
	}
	newAlmanac.addMap(*currentMap)
}

func almanacFromStrings(str []string) (*almanac, error) {
	var newAlmanac = almanac{}

	seeds := parseSeedString(str[0])
	newAlmanac.setSeed(seeds)
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

	return &newAlmanac, nil
}

func Part1() error {
	inputString := common.Input()

	inputAlmanac, err := almanacFromStrings(inputString)
	if err != nil {
		return err
	}

	fmt.Println(inputAlmanac.getSmallest("seed", "location"))

	return nil
}
