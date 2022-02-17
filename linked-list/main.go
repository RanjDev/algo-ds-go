package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func main() {
	fmt.Println("Linked List")
	myList := LinkedList{}

	node1 := &Node{data: 21}
	node2 := &Node{data: 34}
	node3 := &Node{data: 65}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList)
	myList.printListData()

	fmt.Println("After Delete")
	myList.deleteWithValue(34)
	myList.printListData()
}

// takes a node tobe added at the front
func (l *LinkedList) prepend(n *Node) {
	// (l * LinkedList) is a reciever
	// because we want to work with a reference of the struct
	// second := l.head
	// l.head = n
	// l.head.next = second
	// l.length++

	// newNode := Node{data: value}
	if l.head != nil {
		n.next = l.head
		l.head = n
		l.length++
	} else {
		l.head = n
		l.length++
	}
}

func (l *LinkedList) printListData() {
	// toPrint := l.head
	// for l.length != 0 {
	// 	fmt.Printf("%d", toPrint.data)
	// 	toPrint = toPrint.next
	// 	l.length--
	// 	fmt.Printf("\n")
	// }
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) deleteWithValue(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val { //if prev.next.data was the value breaks the loop
		if previousToDelete.next.next == nil {
			return //means that prev.next is the last element so return
		}

		previousToDelete = previousToDelete.next

	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
