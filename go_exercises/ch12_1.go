package main

import (
	"fmt"
)

/*
The following function accepts an array of numbers and returns the sum,
as long as a particular number doesn’t bring the sum above 100. If adding
a particular number will make the sum higher than 100, that number is
ignored. However, this function makes unnecessary recursive calls. Fix
the code to eliminate the unnecessary recursion:

def add_until_100(array)
	return 0 if array.length == 0

	if array[0] + add_until_100(array[1, array.length - 1]) > 100
		return add_until_100(array[1, array.length - 1])
	else
		return array[0] + add_until_100(array[1, array.length - 1])
	end
end


*/

func main() {
	n := []int{95, 1, 1000}
	fmt.Println(addUntil100(n))
}


func addUntil100(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	remainder := addUntil100(numbers[1:])

	if (numbers[0] + remainder) > 100 {
		return remainder
	}

	return numbers[0] + remainder
}