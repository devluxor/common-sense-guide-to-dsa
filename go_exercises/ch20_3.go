package main

import (
	"fmt"

)
/*
You’re working on some more stock-prediction software. The function
you’re writing accepts an array of predicted prices for a particular stock
over the course of time.

For example, this array of seven prices:

[10, 7, 5, 8, 11, 2, 6]

predicts that a given stock will have these prices over the next seven days.
(On Day 1, the stock will close at $10; on Day 2, the stock will close at
$7; and so on.)

Your function should calculate the greatest profit that could be made
from a single “buy” transaction followed by a single “sell” transaction.
In the previous example, the most money could be made if we bought the
stock when it was worth $5 and sold it when it was worth $11. This yields
a profit of $6 per share.

Note that we could make even more money if we buy and sell multiple
times, but for now, this function focuses on the most profit that could be
made from just one purchase followed by one sale.

Now, we could use nested loops to find the profit of every possible buyand-
sell combination. However, this would be O(N2) and too slow for our
hotshot trading platform. Your job is to optimize the code so that the
function clocks in at just O(N).
*/

func main() {
	a := []int{10, 7, 5, 8, 11, 2, 6}
	fmt.Println(bestOperation(a))
}

/*

sample array

init best buy found (i) to a 0

init best sell found (j) to a 1

init best operation found (op) to - Inf

i points to array 0, 

j points to array 1

calculate last index

for every element, until j touches the last index; begin from 1 (i = 1, j = 2) ???

	if i < best buy found, save i

	if j > best sell found, save j

	if i < j, calculate operation, if opeartion result > op. saved; set op

	add 1 to i
	add 1 to j

return best op found
 0   1  2  3  4   5  6
[10, 7, 5, 8, 11, 2, 6]
           i  j
i: 5
j: 6

best buy: 2
best sell: 6

best op.: 6 (11 - 5)

*/

func bestOperation(a []int) int{
	bestBuy := a[0]
	bestSell := a[1]
	bestOperation := bestSell - bestBuy
	lastIndex := len(a) - 1
	currentOperation := 0
	for i, j := 1, 2; j <= lastIndex; i, j = i + 1, j + 1 {
		if a[i] < bestBuy {
			bestBuy = a[i]
		}

		bestSell = a[j]

		currentOperation = bestSell - bestBuy
		if currentOperation > bestOperation {
			bestOperation = currentOperation
		}
	}

	return bestOperation
}
