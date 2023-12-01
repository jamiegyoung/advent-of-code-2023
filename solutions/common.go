package solutions

import (
	"bufio"
	"os"
)

func Input() []string {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "eoi" {
			break
		}

		lines = append(lines, line)
	}

	return lines
}
