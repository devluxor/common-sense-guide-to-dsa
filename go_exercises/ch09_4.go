package main

import "fmt"

/*
Write a function that uses a stack to reverse a string. (For example, "abcde"
would become "edcba".)
*/

func main() {
	fmt.Println(reverse("abcde"))
}

func reverse(sentence string) string {
	q := newQueue()
	for _, runeChar := range sentence {
		q.push(runeChar)
	}
	
	r := ""
	for !q.empty() {
		r += string(q.pop())
	}
	
	return r
}

type Queue struct {
	queue []rune
}

func newQueue() *Queue {
	q := Queue{}

	return &q
}

func (q *Queue) read() rune {
	return q.queue[len(q.queue) - 1]
}

func (q *Queue) push(n rune) rune {
	q.queue = append(q.queue, n)

	return n
}

func (q *Queue) pop() rune {
	qLength := len(q.queue)
	n := q.queue[qLength - 1]
	q.queue = q.queue[:qLength - 1]
	return n
}

func (q *Queue) print() {
	fmt.Println(q.queue)
}

func (q *Queue) empty() bool {
	return len(q.queue) == 0
}

