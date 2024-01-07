package main

import (
	"fmt"
)

func main() {
	a := []int{2, 0, 4, 1, 7}

	fmt.Println(twoSum10(a))
}

func twoSum10(numbers []int) bool {
	// create hash + check condition in a single iteration
	valuesAndIndexes := make(map[int]bool)
	for _, val := range numbers {
		_, differenceExists := valuesAndIndexes[10 - val]
		if differenceExists {
			return true
		}
		
		valuesAndIndexes[val] = true
	}

	return false
}
