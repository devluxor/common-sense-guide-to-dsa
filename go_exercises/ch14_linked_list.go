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

	// l.print()
	// fmt.Println(l.last())
	// fmt.Println(l.lastNoTail())
	// l.removeTarget(2)
	// l.print()
	l.reverse()
	l.print()
}


type List struct {
	head *Node
	tail *Node
}

type Node struct {
	data any
	next *Node
}

func createList(d any) *List {
	newList := &List{}
	newNode := &Node{data: d, next: nil}
	newList.head = newNode
	newList.tail = newNode

	return newList
}

func (list *List) insert(d any) {
	newNode := &Node{data: d, next: nil}

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

func (list *List) last() any {
	return list.tail.data
}

func (list *List) lastNoTail() any {
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (list *List) removeTarget(target int) *Node {
	dummyNode := &Node{data: nil, next: list.head}
	previous := dummyNode
	current := list.head

	for current != nil {
		if current.data == target {
			previous.next = current.next
		} else {
			previous = current
		}

		current = current.next
	}

	// a head, either the original or the new one if we removed the head
	return dummyNode.next 
}

func (list *List) reverse() *Node {
	var previous *Node
	current := list.head
	var next *Node
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	list.head = previous
	return list.head
}
