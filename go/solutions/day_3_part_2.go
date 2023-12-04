package solutions

import (
	"fmt"
	"regexp"
	"strconv"
)

func inRange3Point5(min int, max int, min2 int, max2 int) bool {
	// check if any number is within the range of the other
	aInB := min >= min2 && min <= max2
	bInA := min2 >= min && min2 <= max

	return aInB || bInA
}

func solveRow3Point5(
	strIndex int,
	i int,
	strings []string,
	symbolIndex int,
	foundIndexes map[string]bool,
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

		// this is awful
		uniqueForNum := fmt.Sprintf("%d,%d,%d", rowIndex, min, max)

		if inRange3(symbolIndex-margin, symbolIndex+margin, min, max) && !foundIndexes[uniqueForNum] {
			number, err := strconv.Atoi(row[otherRowNumIndexRange[0]:otherRowNumIndexRange[1]])
			if err != nil {
				return nil, err
			}

			foundIndexes[uniqueForNum] = true
			adjacent = append(adjacent, number)
		}
	}
	return adjacent, nil
}

func solve3Point5(strings []string) (int, error) {
	acc := 0
	foundIndexes := map[string]bool{}
	for strIndex, str := range strings {
		reg := regexp.MustCompile(`\*`)

		symbolsIndexes := reg.FindAllStringIndex(str, -1)

		if symbolsIndexes == nil {
			continue
		}

		for _, symbolIndexes := range symbolsIndexes {
			symbolIndex := symbolIndexes[0] - 1
			symbolIndex++

			adjacent := []int{}
			for i := -1; i < 2; i += 1 {
				vals, err := solveRow3Point5(strIndex, i, strings, symbolIndex, foundIndexes, 1)
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

func Day3Point5() error {
	input := Input()
	res, err := solve3Point5(input)
	if err != nil {
		return err
	}
	fmt.Println("Total ratio:", res)
	return nil
}
