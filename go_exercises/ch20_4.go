package main

import (
	"fmt"

)
/*
You’re writing a function that accepts an array of numbers and computes
the highest product of any two numbers in the array.
*/

func main() {
	a := []int{5, -10, -6, 9, 4}
	fmt.Println(bestOperation(a))
}

/*

sample array

[5, -10, -6, 9, 4]

find two greatest

find two lowest

multiply them, and return the highest result

*/

func bestOperation(a []int) int{
	greatest := -99999
	secondGreatest := -999999

	lowest := 999999
	secondLowest := 999999


	for _, n := range a {
		if n > greatest {
			secondGreatest = greatest
			greatest = n
		} else if n > secondGreatest {
			secondGreatest = n
		}
		
		if n < lowest {
			secondLowest = lowest
			lowest = n
		} else if n < secondLowest && n > lowest {
			secondLowest = n
		}
	}

	r1 := greatest * secondGreatest
	r2 := lowest * secondLowest

	if r2 > r1 {
		return r2
	}

	return r1
}
