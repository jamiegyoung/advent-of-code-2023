package day4

import (
	"advent_of_code/common"
	"fmt"
	"strings"
)

type scratchcard struct {
	line  string
	index int
}

func Part2() error {
	fmt.Println("Please enter the puzzle input: ")

	var input []string = common.Input()
	var scratchcardsQueue = queue[scratchcard]{}

	// start with the first scratchcard
	for i := range input {
		scratchcardsQueue.push(scratchcard{input[i], i})
	}

	for scratchcardsQueue.hasMore() {
		sc := scratchcardsQueue.next()

		numbersString := strings.Split(sc.line, ":")[1]
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
					points++
				}
			}
		}

		// create the new scratchcards from the points
		for j := 1; j < points+1; j++ {
			scratchcardsQueue.push(scratchcard{input[sc.index+j], sc.index + j})
		}
	}

	fmt.Println("Final count", scratchcardsQueue.len())
	return nil
}
