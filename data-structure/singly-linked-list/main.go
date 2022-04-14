package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) Prepend(n *node) {
	secondNode := l.head
	l.head = n
	l.head.next = secondNode
	l.length++
}

func (l *linkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		// handling value not found in linked list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete.next = previousToDelete.next.next
		l.length--
	}
}

func (l linkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	testList := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 20}
	node3 := &node{data: 300}
	node4 := &node{data: 444}

	testList.Prepend(node1)
	testList.Prepend(node2)
	testList.Prepend(node3)
	testList.Prepend(node4)

	testList.PrintListData()
	testList.DeleteWithValue(20)
	testList.PrintListData()
}
