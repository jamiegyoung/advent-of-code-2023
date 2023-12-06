package day4

import (
	"advent_of_code/common"
	"fmt"
	"strings"
)

func Part1() error {
	fmt.Println("Please enter the puzzle input: ")
	var input []string = common.Input()
	acc := 0
	for _, line := range input {
		numbersString := strings.Split(line, ":")[1]
		split := strings.Split(numbersString, "|")

		winningNumbers, err := mapToIntArray(strings.Split(split[0], " "))
		if err != nil {
			fmt.Println("Error converting winning numbers")
			return err
		}

		drawnNumbers, err := mapToIntArray(strings.Split(split[1], " "))
		if err != nil {
			fmt.Println("Error converting drawn numbers")
			return err
		}

		points := 0
		for _, drawnNumber := range drawnNumbers {
			for _, winningNumber := range winningNumbers {
				if drawnNumber == winningNumber {
					if points == 0 {
						points++
						continue
					}
					points *= 2
				}
			}
		}
		acc += points
	}
	fmt.Println(acc)
	return nil
}
