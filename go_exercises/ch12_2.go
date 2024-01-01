package main

import (
	"fmt"
)

/*
The following function uses recursion to calculate the Nth number from
a mathematical sequence known as the “Golomb sequence.” It’s terribly
inefficient, though! Use memoization to optimize it. (You don’t have to
actually understand how the Golomb sequence works to do this exercise.)

def golomb(n)
	return 1 if n == 1

	return 1 + golomb(n - golomb(golomb(n - 1)));
end

*/

func main() {
	// n := 
	fmt.Println(golomb(9, make(map[int]int)))
}


func golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, stored := memo[n]; !stored {
		memo[n] = 1 + golomb(n - golomb(golomb(n - 1, memo), memo), memo)
	}

	return memo[n]
}