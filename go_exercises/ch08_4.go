package main

import "fmt"

/*
Write a function that returns the first non-duplicated character in a string.
For example, the string, "minimum" has two characters that only exist
once—the "n" and the "u", so your function should return the "n", since it
occurs first. The function should have an efficiency of O(N).
*/

func main() {
	// s := "minimum"
	// s := "iiiiiua"
	// s := "aiiiiiiiiu"
	s := "min"

	fmt.Println(firstUnique(s))
}

func firstUnique(sentence string) string {
	lastIndex := len(sentence) - 1

	seen := make(map[rune]int)
	candidates := make(map[rune]int)
	var inSeen bool
	var inCandidates bool
	for i, runeChar := range sentence {
		_, inSeen = seen[runeChar]
		_, inCandidates = candidates[runeChar]

		if !inSeen && i >= (lastIndex-1) && len(candidates) != 0 {
			break
		} else if !inSeen {
			seen[runeChar] = i
			candidates[runeChar] = i
		} else if inSeen && inCandidates {
			delete(candidates, runeChar)
		}
	}

	var res string
	for k := range candidates {
		res = string(k)
	}
	return res
}
