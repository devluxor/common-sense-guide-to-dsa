package main

import (
	"fmt"
)

/*
This problem is known as the “Unique Paths” problem: Let’s say you have
a grid of rows and columns. Write a function that accepts a number of rows
and a number of columns, and calculates the number of possible “shortest”
paths from the upper-leftmost square to the lower-rightmost square.


From the starting position, we have only two choices of movement. We
can either move one space to the right or one space downward.

the total number of unique shortest paths will
be:

the number of paths from space to the right of S + the number of paths
from space below S.


*/

func main() {
	fmt.Println(uniquePaths(7, 3))
}


func uniquePaths(rows, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	return uniquePaths(rows - 1, columns) + uniquePaths(rows, columns - 1)
}