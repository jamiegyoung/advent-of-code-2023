package main

import (
	"advent_of_code/solutions"
	"fmt"
)

type SelectionError struct {
	Field string
	Msg   string
}

func (e SelectionError) Error() string {
	return fmt.Sprintln("User inputted, %s: %s", e.Field, e.Msg)
}

var days = map[string]func() error{
	"1": solutions.Day1,
	"2": solutions.Day2,
	"3": solutions.Day3,
  "3.5": solutions.Day3Point5,
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
