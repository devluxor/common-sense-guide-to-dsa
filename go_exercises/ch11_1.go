package main

import (
	"fmt"
)

/*
Use recursion to write a function that accepts an array of strings and
returns the total number of characters across all the strings. For example,
if the input array is ["ab", "c", "def", "ghij"], the output should be 10 since there
are 10 characters in total.
*/

func main() {
	a := []string{"ab", "c", "def", "ghij"}

	fmt.Println(count(a))
}

func count(words []string) int {
	if len(words) == 1 {
		return len(words[0])
	}
	
	return len(words[0]) + count(words[1:len(words)])
}



