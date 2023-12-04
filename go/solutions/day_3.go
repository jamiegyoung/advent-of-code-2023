package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type point3 struct {
	rowIndex int
	min      int
	max      int
}

type pointArr3 []point3

func (s pointArr3) contains(e point3) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func inRange3(min int, max int, min2 int, max2 int) bool {
	// check if any number is within the range of the other
	aInB := min >= min2 && min <= max2
	bInA := min2 >= min && min2 <= max

	return aInB || bInA
}

func solveRow3(
	strIndex int,
	i int,
	strings []string,
	symbolIndex int,
	foundIndexes pointArr3,
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

		point := point3{rowIndex, min, max}

		if inRange3(symbolIndex-margin, symbolIndex+margin, min, max) && !foundIndexes.contains(point) {
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

func solve3(strings []string) error {
	acc := 0
	foundIndexes := pointArr3{}
	for strIndex, str := range strings {
		reg := regexp.MustCompile(`[^0-9.]`)

		symbolsIndexes := reg.FindAllStringIndex(str, -1)

		if symbolsIndexes == nil {
			continue
		}

		for _, symbolIndexes := range symbolsIndexes {
			symbolIndex := symbolIndexes[0] - 1
			symbolIndex++

			for i := -1; i < 2; i += 1 {
				err := solveRow3(strIndex, i, strings, symbolIndex, foundIndexes, &acc, 1)
				if err != nil {
					return err
				}
			}
		}
	}
	fmt.Println("Sum:", acc)
	return nil
}

func Day3() error {
	input := Input()
	// input := []string{
	// 	"1.2",
	// 	"@..",
	// }
	fmt.Println(strings.Join(input, "\n"))
	fmt.Println()
	return solve3(input)
}
