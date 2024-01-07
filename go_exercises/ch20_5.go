package main

import (
	"fmt"

)
/*
You are to write a function that sorts these readings from lowest to highest.
*/

func main() {
	a := []int{986, 980, 971, 990, 989, 978, 985, 982, 980, 971}
	fmt.Println(sortTemps(a))
}

/*

*/

func sortTemps(a []int) []int{
	r := []int{}
	temps := make(map[int]int)

	for _, t := range a {
		if _, e := temps[t]; e {
			temps[t] += 1
		} else {
			temps[t] = 1
		}
	}

	temp := 970
	times := 0
	for temp <= 990 {
		if _, e := temps[temp]; e {
			fmt.Print(temp, "\n")
			times = temps[temp]
			for times > 0 {
				r = append(r, temp)
				times -= 1
			}
		}
		temp += 1
	}

	return r
}
