package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day_1() {
	var lines = Input()
	var cumSum = 0

	for _, line := range lines {
		var first int
		var second int

		for _, char := range strings.Split(line, "") {
			i, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			if first == 0 {
				first = i
			}

			second = i
		}

		fmt.Println("first:", first, "second:", second)
		joined, err := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
		if err != nil {
			fmt.Println("An error occured parsing the joined numbers")
			os.Exit(1)
		}
		cumSum += joined
	}

	fmt.Println()
	fmt.Println(cumSum)
}

