package day3

import (
	"advent_of_code/common"
	"fmt"
	"regexp"
	"strconv"
)

func solveRowAlt(
	strIndex int,
	i int,
	strings []string,
	symbolIndex int,
	foundIndexes pointArr,
	margin int,
) ([]int, error) {
	rowIndex := strIndex + i
	// ignore out of bounds
	if rowIndex < 0 || rowIndex > len(strings)-1 {
		return nil, nil
	}

	row := strings[rowIndex]
	// check for numbers on other rows
	numReg := regexp.MustCompile(`[0-9]+`)
	otherRowNumIndexRanges := numReg.FindAllStringIndex(row, -1)
	adjacent := []int{}

	for _, otherRowNumIndexRange := range otherRowNumIndexRanges {
		min := otherRowNumIndexRange[0]
		max := otherRowNumIndexRange[1] - 1

		point := point{rowIndex, min, max}

		if inRange(symbolIndex-margin, symbolIndex+margin, min, max) && !foundIndexes.contains(point) {
			number, err := strconv.Atoi(row[otherRowNumIndexRange[0]:otherRowNumIndexRange[1]])
			if err != nil {
				return nil, err
			}

			foundIndexes = append(foundIndexes, point)
			adjacent = append(adjacent, number)
		}
	}
	return adjacent, nil
}

func solveAlt(strings []string) (int, error) {
	acc := 0
	foundIndexes := pointArr{}
	for strIndex, str := range strings {
		reg := regexp.MustCompile(`\*`)

		symbolsIndexes := reg.FindAllStringIndex(str, -1)

		if symbolsIndexes == nil {
			continue
		}

		for _, symbolIndexes := range symbolsIndexes {
			symbolIndex := symbolIndexes[0]

			adjacent := []int{}
			for i := -1; i < 2; i += 1 {
				vals, err := solveRowAlt(strIndex, i, strings, symbolIndex, foundIndexes, 1)
				if err != nil {
					return 0, err
				}
				for _, val := range vals {
					adjacent = append(adjacent, val)
				}
			}

			if len(adjacent) > 1 {
				ratio := 1
				for _, adjVal := range adjacent {
					ratio *= adjVal
				}

				acc += ratio
			}
		}
	}
	return acc, nil
}

func Part2() error {
	fmt.Println("Please enter the input and then enter \"eoi\" on the next line")
	input := common.Input()
	res, err := solveAlt(input)
	if err != nil {
		return err
	}
	fmt.Println("Total ratio:", res)
	return nil
}
