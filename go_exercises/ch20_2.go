package main

import (
	"fmt"

)

// You are to write a function that accepts two arrays of players and returns
// an array of the players who play in both sports. In this case, that would be:

func main() {
	a := []int{3, 4, 1}
	fmt.Println(findMissing(a))
}

func findMissing(numbers []int) int{
	min := numbers[0]
	present := make(map[int]bool)
	for _, n := range numbers {
		present[n] = true
		if n < min {
			min = n
		}
	}

	i := min + 1
	for {
		if !present[i] {
			return i
		}
		i++
	}

	return i
}
