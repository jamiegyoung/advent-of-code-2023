package main

import (
	"advent_of_code/day1"
	"advent_of_code/day2"
	"advent_of_code/day3"
	"fmt"
)

var days = map[string]func() error{
	"1":   day1.Day1,
	"2":   day2.Day2,
	"3":   day3.Part1,
	"3.5": day3.Part2,
}

func main() {
	fmt.Print("Please enter a day to run: ")

	var day string
	fmt.Scanln(&day)

	if day == "" {
		fmt.Println("No day was entered.")
		return
	}

	// check function exists
	var foundSol = days[day]

	if foundSol == nil {
		fmt.Println("Day", day, "has not been solved")
		return
	}

	err := foundSol()
	if err != nil {
		fmt.Println(err)
	}
}
