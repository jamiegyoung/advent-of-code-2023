package common

import (
	"bufio"
	"errors"
	"os"
)

func InputError(msg string) error {
	return errors.New("Invalid input: " + msg)
}

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

