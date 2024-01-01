package main

import (
	"fmt"
)

/*
Use recursion to write a function that accepts a string and returns the
first index that contains the character “x.” 

For example, the string,

"abcdefghijklmnopqrstuvwxyz" 

has an “x” at index 23. 

To keep things simple,
assume the string definitely has at least one “x.”
*/

func main() {
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(findX(s))
}


func findX(s string) int {
	if s[0] == 'x' {
		return 0
	}

	return findX(s[1:len(s)]) + 1
}