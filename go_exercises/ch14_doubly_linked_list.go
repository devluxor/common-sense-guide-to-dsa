package main

import (
	"fmt"
)

func main() {
	l := createList(1)
	i := 2
	for i <= 10 {
		l.insert(i)
		i++
	}

	l.print()
	l.printReverse()
}

// Doubly Linked List
type List struct {
	head *Node
	tail *Node
}

// Doubly Linked List Node
type Node struct {
	data any
	previous *Node
	next *Node
}

func createList(d any) *List {
	newList := &List{}
	newNode := &Node{data: d, previous: nil, next: nil}
	newList.head = newNode
	newList.tail = newNode

	return newList
}

func (list *List) insert(d any) {
	newNode := &Node{data: d, previous: list.tail, next: nil}

	list.tail.next = newNode
	list.tail = newNode
}

func (list *List) print() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (list *List) printReverse() {
	currentNode := list.tail
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.previous
	}
}


