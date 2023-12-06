package main

import (
	"advent_of_code/day1"
	"advent_of_code/day2"
	"advent_of_code/day3"
	"advent_of_code/day4"
	"advent_of_code/day5"
	"advent_of_code/day5Part2"
	"fmt"
)

var days = map[string]func() error{
	"1":   day1.Day1,
	"2":   day2.Day2,
	"3":   day3.Part1,
	"3.5": day3.Part2,
	"4":   day4.Part1,
	"4.5": day4.Part2,
	"5":   day5.Part1,
	"5.5": day5Part2.Part2,
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
