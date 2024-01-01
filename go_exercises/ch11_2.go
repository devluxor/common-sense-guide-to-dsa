package main

import (
	"fmt"
)

/*
Use recursion to write a function that accepts an array of numbers and
returns a new array containing just the even numbers.
*/

func main() {
	a := []int{1, 2, 3, 6, 40}

	fmt.Println(filterEven(a))
}

func filterEven(numbers []int) []int {
	if len(numbers) == 1 {
		if isEven(numbers[0]) {
			return numbers
		}

		return []int{}
	}

	nextNumbers := []int{}

	if isEven(numbers[0]) {
		nextNumbers = append(nextNumbers, numbers[0])
	}

	return append(nextNumbers, filterEven(numbers[1:len(numbers)])...)
}

func isEven(n int) bool {
	return n % 2 == 0
}

