package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{7, 8, 1, 5}
	quicksort(a, 0, len(a) - 1)
	fmt.Println(a)
}

func quicksort(numbers []int, low, high int) {
	pivotIndex := partition(numbers, low, high)
	if low < pivotIndex -1 {
		quicksort(numbers, low, pivotIndex - 1) 
	}
	if pivotIndex < high {
		quicksort(numbers, pivotIndex, high)
	}
}

func partition(numbers []int, low, high int) int {
	pivotIndex := int(math.Floor(float64(low + high) / 2))
	pivot := numbers[pivotIndex]
	// fmt.Println(numbers, "low: ", low, "; high: ", high, "; pivotIndex: ", pivotIndex, "; pivot: ", pivot)
	left := low
	right := high

	for left <= right {
		for numbers[left] < pivot {
			left++
		}

		for numbers[right] > pivot {
			right--
		}

		if left > right {
			break
		}

		// fmt.Println("swap ", numbers[left], " for ", numbers[right])
		// swap values at left & right pointers 
		numbers[left], numbers[right] = numbers[right], numbers[left]

		left++
		right--
	}

	// fmt.Println("returns ", left)
	return left
}



