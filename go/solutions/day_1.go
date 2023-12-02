package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_1() {
	numberStrings := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	longestNumberStringCount := 5
	smallestNumberStringCount := 3

	lines := Input()
	cumSum := 0

	for _, line := range lines {
		numericalLettersAcc := []int{}

		for i := 0; i < len(strings.Split(line, "")); i++ {
			charAsString := string(line[i])

			foundNumber, err := strconv.Atoi(charAsString)
			if err == nil {
				// handle numbers
				numericalLettersAcc = append(numericalLettersAcc, foundNumber)
				continue
			}

			// handle letters
			for j := smallestNumberStringCount; j <= longestNumberStringCount; j++ {
				if i+j > len(line) {
					continue
				}

				subString := line[i : i+j]
				if numberStrings[subString] != 0 {
					numericalLettersAcc = append(numericalLettersAcc, numberStrings[subString])
					continue
				}
			}
		}

		joined, err := strconv.Atoi(fmt.Sprintf("%d%d", numericalLettersAcc[0], numericalLettersAcc[len(numericalLettersAcc)-1]))
		if err != nil {
			fmt.Println("An error occured parsing the joined numbers")
			os.Exit(1)
		}
		cumSum += joined
	}

	fmt.Println()
	fmt.Println(cumSum)
}
