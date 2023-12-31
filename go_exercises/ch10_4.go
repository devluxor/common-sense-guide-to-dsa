package main

import (
	"fmt"
	"reflect"
)

/*
Here is an array containing both numbers as well as other arrays, which
in turn contain numbers and arrays:

array = [ 
	1,
	2,
	3,
	[4, 5, 6],
	7,
	[8, [9, 10, 11,	[12, 13, 14]]],
	[15, 16, 17, 18, 19, [20, 21, 22,	[23, 24, 25, [26, 27, 29] ], 30, 31], 32], 
	33
]
Write a recursive function that prints all the numbers (and just numbers).

*/

func main() {
	a := 
		[]interface{}{1,
			2,
			3,
			[]interface{}{4, 5, 6},
			7,
			[]interface{}{8, []interface{}{9, 10, 11, []interface{}{12, 13, 14}}},
			[]interface{}{15, 16, 17, 18, 19, []interface{}{20, 21, 22, []interface{}{23, 24, 35}}},
			25}

	traverse(a)
}

func traverse(element any) {
	if isNumber(element) {
		fmt.Println(element)
		return
	}

	for _, e := range element.([]interface{}) {
		traverse(e)
	}
}

func isNumber(a interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(1)
}


