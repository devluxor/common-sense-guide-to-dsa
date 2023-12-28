package main 

import "fmt"

/*
The following function finds the greatest single number within an array,
but has an efficiency of O(N2). Rewrite the function so that it becomes a
speedy O(N)
*/

func main() {
	a := []int{1, 2, 99}
	fmt.Println(greatestNumber(a))
}

func greatestNumber(numbers []int) int {
	greatestFound := numbers[0]
	number := 0
	for i := 1; i < len(numbers); i++ {
		number = numbers[i]
		if number > greatestFound {
			greatestFound = number
		}
	}

	return greatestFound
}