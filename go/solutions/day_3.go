package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func inRange(min int, max int, min2 int, max2 int) bool {
	// check if any number is within the range of the other
	aInB := min >= min2 && min <= max2
	bInA := min2 >= min && min2 <= max

	return aInB || bInA
}

func solveRow(
	strIndex int,
	i int,
	strings []string,
	symbolIndex int,
	foundIndexes map[string]bool,
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

		// this is awful
		uniqueForNum := fmt.Sprintf("%d,%d,%d", rowIndex, min, max)

		if inRange(symbolIndex-margin, symbolIndex+margin, min, max) && !foundIndexes[uniqueForNum] {
			number, err := strconv.Atoi(row[otherRowNumIndexRange[0]:otherRowNumIndexRange[1]])
			if err != nil {
				return err
			}

			fmt.Println("Adding number", number, *acc)
			foundIndexes[uniqueForNum] = true
			*acc += number
		}
	}
	return nil
}

func solve(strings []string) error {
	acc := 0
	foundIndexes := map[string]bool{}
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

func Day3() error {
	input := Input()
	// input := []string{
	// 	"1.2",
	// 	"@..",
	// }
	fmt.Println(strings.Join(input, "\n"))
	fmt.Println()
	return solve(input)
}
