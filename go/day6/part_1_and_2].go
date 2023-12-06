package day6

import (
	"advent_of_code/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type race struct {
	time   int
	record int
}

func filterEmpty(strs []string) []string {
	res := []string{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		res = append(res, str)
	}
	return res
}

func getCleanInputTable(input []string) [][]string {
	inputClean := [][]string{}
	for i, v := range input {
		for _, s := range filterEmpty(strings.Split(v, " "))[1:] {
			if i >= len(inputClean) {
				inputClean = append(inputClean, []string{})
			}
			inputClean[i] = append(inputClean[i], s)
		}
	}
	return inputClean
}

func parseRaces(str [][]string) ([]race, error) {

	races := []race{}
	for columnIndex := range str[0] {
		time, err := strconv.Atoi(str[0][columnIndex])
		if err != nil {
			return nil, err
		}

		record, err := strconv.Atoi(str[1][columnIndex])
		if err != nil {
			return nil, err
		}

		races = append(races, race{time: time, record: record})

	}
	return races, nil
}

func solve(input []string) error {
	inputClean := getCleanInputTable(input)

	races, err := parseRaces(inputClean)
	if err != nil {
		return err
	}

	winningAcc := 1
	for _, race := range races {
		time := float64(race.time)
		record := float64(race.record)
		minSpeed := (time - math.Sqrt(math.Pow(time, 2)-4*record)) / 2
		maxSpeed := (time + math.Sqrt(math.Pow(time, 2)-4*record)) / 2
		winningAcc *= int(math.Ceil(maxSpeed) - 1 - math.Floor(minSpeed))
	}
	fmt.Println(winningAcc)
	return nil
}

func Part1() error {
	input := common.Input()

	inputPartTwo := []string{}

	for _, line := range input {
		cols := filterEmpty(strings.Split(line, " "))
		var restOfCols string
		for _, r := range cols[1:] {
			restOfCols = restOfCols + r
		}
		newLine := cols[0] + " " + restOfCols
		inputPartTwo = append(inputPartTwo, newLine)
	}

	err := solve(input)
	if err != nil {
		return err
	}

	err = solve(inputPartTwo)
	if err != nil {
		return err
	}
	return nil
}
