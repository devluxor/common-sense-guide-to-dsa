package main

import (
	"fmt"
)

func main() {
	l := createList(1)

	// i := 2
	// for i <= 10 {
	// 	l.insert(i)
	// 	i++
	// }
	
	l.insert(1)
	l.insert(1)
	l.insert(2)
	l.insert(1)
	l.insert(3)
	l.insert(1)
	l.insert(4)

	l.print()
	// fmt.Println(l.last())
	// fmt.Println(l.lastNoTail())
	// l.removeTarget(2)
	// l.print()
	// l.reverse()

	l.head.removeDuplicates()


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
		if current.data == target { // 
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

// func (list *List) removeDuplicates() *Node {
// 	dummy := &Node{data: nil, next: list.head}
// 	previous := dummy
// 	current := list.head
// 	seen := make(map[int]bool)
// 	var currentData int

// 	for current != nil {
// 		currentData = current.data.(int)
		
// 		if seen[currentData] {
// 			previous.next = current.next
// 		} else {
// 			previous = current
// 		}
		
// 		seen[currentData] = true
// 		current = current.next
// 	}

// 	return dummy.next 
// }

func (head *Node) removeDuplicates() *Node {
	dummy := &Node{data: nil, next: head}
	previous := dummy
	current := head
	seen := make(map[int]bool)
	var currentData int

	for current != nil {
		currentData = current.data.(int)
		
		if seen[currentData] {
			previous.next = current.next
		} else {
			previous = current
		}
		
		seen[currentData] = true
		current = current.next
	}

	return dummy.next 
}

/*
for every node

save node data

if node is seen, delete node

remove duplicates with space complexity of 1

try

current = head
next = current.next

for next != nil {
	if curr.val == next.val {
		next = next.next
	} else {
		current.next = next
		current = next
		next = current.next
	}
}

current.next = nil
return head
*/
