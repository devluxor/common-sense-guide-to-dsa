package main

import (
	"fmt"
)

/*
Here is a solution to the “Unique Paths” problem from an exercise in the
previous chapter. Use memoization to improve its efficiency:

def unique_paths(rows, columns)
	return 1 if rows == 1 || columns == 1
	return unique_paths(rows - 1, columns) + unique_paths(rows, columns - 1)
end

*/

func main() {
	fmt.Println(uniquePaths(7, 3, make(map[uint64]int)))
}

func uniquePaths(rows, columns int, memo map[uint64]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	rCsKey := rCsToKey(rows, columns)
	if _, stored := memo[rCsKey]; !stored {
		memo[rCsKey] = uniquePaths(rows-1, columns, memo) + uniquePaths(rows, columns-1, memo)
	}

	return memo[rCsKey]
}

func rCsToKey(rows, columns int) uint64 {
	return uint64(rows << 32 | columns)
}

func keyToRCs(key uint64) (uint32, uint32) {
	var rows uint32 = uint32(key >> 32)
	var columns uint32 = uint32(key)

	return rows, columns
}

// hashing function to store row + column (in C)

/*
#include <stdio.h>
#include <stdlib.h>

int main() {
    unsigned int low = 1;
    unsigned int high = 5000;

    unsigned long long data64;

    data64 = (unsigned long long) low << 32 | high;

    printf ("%llx\n", data64);  hexadecimal output
    printf ("%lld\n", data64);  decimal output
    // printf ("0b%b\n", data64);  binary output
    // while (data64) {
    //     if (data64 & 1)
    //         printf("1");
    //     else
    //         printf("0");

    //     data64 >>= 1;
    // }
    // printf("\n");

    unsigned int x, y;
    x = (unsigned long) data64 >> 32;
    y = (unsigned long)(data64);

    // while (data64) {
    //     if (data64 & 1)
    //         printf("1");
    //     else
    //         printf("0");

    //     data64 >>= 1;
    // }
    // printf("\n");

    printf ("%d\n", x);  decimal output
    printf ("%d\n", y);  decimal output

}


In Go:

func rCsToKey(rows, columns int) uint64 {
	return uint64(rows << 32 | columns)
}

func keyToRCs(key uint64) int, int {
	var rows uint32 = uint32(key >> 32)
	var columns uint32 = uint32(key)

	return rows, columns
}

*/
