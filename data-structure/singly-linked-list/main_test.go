package main

import (
	"testing"
)

func TestPrependList(t *testing.T) {
	testList := LinkedList{}
	node1 := &Node{data: 1}

	// Act
	testList.Prepend(node1)

	// Assertion
	if testList.head.data != 1 {
		t.Fatalf(`prepend 1 node to empty LinkedList should make the prepended node as head `)
	}

	node2 := &Node{data: 2}
	node3 := &Node{data: 3}

	// Act
	testList.Prepend(node2)
	testList.Prepend(node3)

	// Assertion
	if testList.head.data != 3 {
		t.Fatalf(`prepend node to non-empty LinkedList should update the head data`)
	}
}

func TestDeleteWithValue(t *testing.T) {
	list := LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	list.Prepend(node1)
	list.Prepend(node2)
	list.Prepend(node3)
	list.Prepend(node4)
	testList := list

	// Act
	// Delete head data
	testList.DeleteWithValue(4)

	// Assertion
	if testList.head.data == 4 {
		t.Fatalf(`delete head node should change head data`)
	}

	if testList.head.data != 3 {
		t.Fatalf(`delete head node should change head data to the second node data`)
	}

	if testList.head.data != 3 {
		t.Fatalf(`delete head node should change head data to the second node data`)
	}

	// Act
	testList = list
	testList.DeleteWithValue(2)

	// Assertion
	observedNode := testList.head.next.next.data
	if observedNode == 2 || observedNode != 1 {
		t.Fatalf(`delete node should change the data to be the next node data`)
	}
}
