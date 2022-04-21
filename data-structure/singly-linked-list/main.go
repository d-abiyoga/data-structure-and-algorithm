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

func (l *LinkedList) Prepend(n *Node) {
	secondNode := l.head
	l.head = n
	l.head.next = secondNode
	l.length++
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		// handling value not found in linked list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	testList := LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 20}
	node3 := &Node{data: 300}
	node4 := &Node{data: 444}

	testList.Prepend(node1)
	testList.Prepend(node2)
	testList.Prepend(node3)
	testList.Prepend(node4)

	testList.PrintListData()
	testList.DeleteWithValue(20)
	testList.PrintListData()
}
