package main

import "fmt"

func main() {
	a := []string{"a", "b", "b", "c", "c"}
	fmt.Println(findDuplicate(a))
}

func findDuplicate(words []string) string {
	seen := make(map[string]bool)
	firstDuplicate := ""
	for _, word := range words {
		if seen[word] {
			firstDuplicate = word
			break
		}
		seen[word] = true
	}

	return firstDuplicate
}