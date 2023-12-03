package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game = map[string]*cube

type cube struct {
	color  string
	amount int
}

func generateLimitsMap() (map[string]int, error) {
	fmt.Println("Please enter a max amount of each color (e.g. 10 red<newline>5 blue<newline>eoi):")
	fmt.Println("If you just enter eoi it will default to 12 red, 13 green, and 14 blue")

	max := Input()
	if len(max) == 0 {
		return map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}, nil
	}

	limits := make(map[string]int)

	for _, val := range max {
		if len(val) < 2 {
			return nil, InputError("Input was too short")
		}

		splitVal := strings.Split(val, " ")

		parsedAmount, err := strconv.Atoi(splitVal[0])
		if err != nil {
			return nil, InputError("Amount was not a number")
		}

		limits[splitVal[1]] = parsedAmount
	}

	return limits, nil
}

func Day2() error {
	limits, err := generateLimitsMap()
	if err != nil {
		return err
	}

	// Scan for new input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a game string (e.g. Game 1: 1 blue, 2 pink; 3 blue, 4 pink) and then type eoi:")

	revealSuccessAcc := 0
	minPowerSetAcc := 0
	for scanner.Scan() {
		gameString := scanner.Text()

		// Just ignore empty lines, pasting an input can result in an additional
		// newline
		if gameString == "" {
			continue
		}

		if gameString == "eoi" || gameString == "q" {
			fmt.Println("Reveal success: " + strconv.Itoa(revealSuccessAcc))
			fmt.Println("Calculated power set: ", minPowerSetAcc)
			break
		}

		idReg := regexp.MustCompile(`Game (\d+)`)

		idMatch := idReg.FindStringSubmatch(gameString)
		if idMatch == nil {
			return InputError("Game string was not in the format 'Game <id>'")
		}

		revealRegs := regexp.MustCompile(`(?:(\d+ \w+)[,;]?)+`)

		revealMatches := revealRegs.FindAllStringSubmatch(gameString, -1)

		fail := false
		foundMaxColors := make(map[string]int)
		for _, revealMatch := range revealMatches {
			revealSplit := strings.Split(revealMatch[1], " ")
			foundLimit, exists := limits[revealSplit[1]]

			foundAmount, err := strconv.Atoi(revealSplit[0])
			if err != nil {
				return InputError("Number was not in the format <number> <color>")
			}

			prevCount, exists := foundMaxColors[revealSplit[1]]

			if !exists || (exists && prevCount < foundAmount) {
				foundMaxColors[revealSplit[1]] = foundAmount
			}

			if foundLimit < foundAmount {
				fail = true
			}
		}

		powerSet := 1
		for _, v := range foundMaxColors {
			powerSet *= v
		}

		minPowerSetAcc += powerSet

		if fail {
			fmt.Println()
			continue
		}

		idInt, err := strconv.Atoi(idMatch[1])
		if err != nil {
			return InputError("Game id was not a number")
		}

		revealSuccessAcc += idInt
	}
	return nil
}
