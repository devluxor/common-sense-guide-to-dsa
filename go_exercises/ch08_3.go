package main

import "fmt"

/*
Write a function that accepts a string that contains all the letters of the
alphabet except one and returns the missing letter. For example, the string,
"the quick brown box jumps over a lazy dog" contains all the letters of the alphabet
except the letter, "f". The function should have a time complexity of O(N).
*/

func main() {
	s := "the quick brown box jumps over a lay dogf"
	fmt.Println(missingLetter(s))
}

func missingLetter(sentence string) string {
	seen := make(map[rune]bool) 
	letterCodeSum := sumLetterCodes()
	for _, runeChar := range sentence {
		if runeChar == ' ' {
			continue
		}

		if !seen[runeChar] {
			seen[runeChar] = true
			letterCodeSum -= int(runeChar)
		}
	}

	return string(letterCodeSum)
}

func sumLetterCodes() int {
	aCode := int('a')
	zCode := int('z')
	sum := 0
	for i := aCode; i <= zCode; i++ {
		sum += i
	}

	return sum
}