package day3

import (
	"advent_of_code/common"
	"fmt"
	"regexp"
	"strconv"
)

func solveRow(
	strIndex int,
	i int,
	strings []string,
	symbolIndex int,
	foundIndexes pointArr,
	acc *int,
	margin int,
) error {
	rowIndex := strIndex + i
	// ignore out of bounds
	if rowIndex < 0 || rowIndex > len(strings)-1 {
		return nil
	}

	row := strings[rowIndex]
	// check for numbers on other rows
	numReg := regexp.MustCompile(`[0-9]+`)
	otherRowNumIndexRanges := numReg.FindAllStringIndex(row, -1)
	for _, otherRowNumIndexRange := range otherRowNumIndexRanges {
		min := otherRowNumIndexRange[0]
		max := otherRowNumIndexRange[1] - 1

		point := point{rowIndex, min, max}

		if inRange(symbolIndex-margin, symbolIndex+margin, min, max) && !foundIndexes.contains(point) {
			number, err := strconv.Atoi(row[otherRowNumIndexRange[0]:otherRowNumIndexRange[1]])
			if err != nil {
				return err
			}

			fmt.Println("Adding number", number, *acc)
			foundIndexes = append(foundIndexes, point)
			*acc += number
		}
	}
	return nil
}

func solve(strings []string) error {
	acc := 0
	foundIndexes := pointArr{}
	for strIndex, str := range strings {
		reg := regexp.MustCompile(`[^0-9.]`)

		symbolsIndexes := reg.FindAllStringIndex(str, -1)

		if symbolsIndexes == nil {
			continue
		}

		for _, symbolIndexes := range symbolsIndexes {
			symbolIndex := symbolIndexes[0]

			for i := -1; i < 2; i += 1 {
				err := solveRow(strIndex, i, strings, symbolIndex, foundIndexes, &acc, 1)
				if err != nil {
					return err
				}
			}
		}
	}
	fmt.Println("Sum:", acc)
	return nil
}

func Part1() error {
	fmt.Println("Please enter the input and then enter \"eoi\" on the next line")
	input := common.Input()
	return solve(input)
}
