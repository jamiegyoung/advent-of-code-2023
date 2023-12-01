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

func main() {
	fmt.Print("Please enter a day to run: ")

	var day string
	fmt.Scanln(&day)

	if day == "" {
		fmt.Println("No day was entered.")
		return
	}

	solutions := map[string]func(){
		"1": solutions.Day_1,
	}

	// check function exists
	var foundSol = solutions[day]

	if foundSol == nil {
		fmt.Println("Day", day, "has not been solved")
		return
	}

	foundSol()
}
