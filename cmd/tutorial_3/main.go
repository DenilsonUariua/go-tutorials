package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()
}

func printMe() {
	fmt.Println("Tutorial 3!")

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	} else if remainder == 0 {
		fmt.Printf("No remainder for %d / %d\n", numerator, denominator)
	}

	fmt.Printf("The result of %d / %d is %d with a remainder of %d\n", numerator, denominator, result, remainder)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")

		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
