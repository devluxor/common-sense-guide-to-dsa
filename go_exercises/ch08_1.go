package main

import "fmt"

/*
arrays can have repeated characters
arrays can have diff. lengths (even 0)
arrays are not sorted

get lenghts of boths a

if one of them is empty, return []

init output
init seen A
init seen B

for every i 0 to len of longest

		- if i < len of A, get char at array A[i]

				- if char in seen B < 1? 
						- add char to output
						- add 1 to seenB[char]
						- add 1 to seenA[char]
				- else
						- add to seenA[char] = 0 if not exists

		- if i < len of B, get char at array B[i]

				- if char in seen A < 1?
						- add char to output
						- add 1 to seenA[char]
						- add 1 to seenB[char]
				- else
						- add to seenB[char] = 0 if not exists

return output

A [ab]
B [bcb]

current char: 
i = 2
A = 
B = b

Seen in A:
a = 1

Seen in B
b = 1
a = 1
c = 0

OUTPUT [b]

*/

func main() {
	a1 := []int{1, 1, 1, 5}
	a2 := []int{2, 3, 1, 1}
	fmt.Println(intersection(a1, a2))
}

func intersection(arrayA []int, arrayB []int) []int {
	lenA := len(arrayA)
	lenB := len(arrayB)
	output := make([]int, 0)
	if lenA == 0 || lenB == 0 {
		return output
	}

	longestLen := longest(lenA, lenB)

	currentInt := 0
	seenA := make(map[int]int)
	seenB := make(map[int]int)
	for i := 0; i < longestLen; i++ {
		if i < lenA {
			currentInt = arrayA[i]
			if seenB[currentInt] == 1{
				output = append(output, currentInt)
				seenA[currentInt] += 1
				seenB[currentInt] += 1
			} else if seenA[currentInt] == 0 {
				seenA[currentInt] = 1
			}
		}

		if i < lenB {
			currentInt = arrayB[i]
			if seenA[currentInt] == 1 {
				output = append(output, currentInt)
				seenA[currentInt] += 1
				seenB[currentInt] += 1
			} else if seenB[currentInt] == 0 {
				seenB[currentInt] = 1
			}
		}
	}

	return output
}

func longest (a, b int) int {
	if a > b {
		return a
	}

	return b
}