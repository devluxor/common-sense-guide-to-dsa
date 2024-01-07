package main

import (
	"fmt"

)
/*
You’re writing a function that accepts an array of unsorted integers and
returns the length of the longest consecutive sequence among them. The
sequence is formed by integers that increase by 1
*/

func main() {
	a := []int{119, 13, 15, 12, 18, 14, 17, 11}
	fmt.Println(longestSequence(a))
}

/*

*/

func longestSequence(a []int) int{
	hash := make(map[int]bool)

	for _, n := range a {
		hash[n] = true
	}

	
	var sequenceBase bool
	var isolatedNumber bool
	var nextNumber int
	longest := 0
	currentLength := 0
	for _, n := range a {
		sequenceBase = !hash[n - 1]
		isolatedNumber = sequenceBase && !hash[n + 1]
		
		if !sequenceBase || isolatedNumber {
			continue
		}
		
		nextNumber = n
		currentLength = 0
		for hash[nextNumber] {
			currentLength += 1
			nextNumber += 1
		}

		if currentLength > longest {
			longest = currentLength
		}
	}

	return longest
}
