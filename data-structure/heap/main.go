package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	slice []int
}

// Insert adds an elemnet to the heap
func (h *MaxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.slice[0]

	if len(h.slice) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	l := len(h.slice) - 1
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	// for h.slice[left(index)] == h.slice[len(h.slice)-1] {
	// 	curr := h.slice[index]
	// 	l := h.slice[left(index)]
	// 	r := h.slice[right(index)]
	// 	if l > r {
	// 		if l > curr {
	// 			h.swap(curr, left(index))
	// 			index = left(index)
	// 		}
	// 	} else {
	// 		if r > curr {
	// 			h.swap(curr, right(index))
	// 			index = right(index)
	// 		}
	// 	}
	// }

	lastIndex := len(h.slice) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.slice[l] > h.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.slice[index] < h.slice[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get left child index
func left(p int) int {
	return p*2 + 1
}

// Get left child index
func right(p int) int {
	return p*2 + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 22, 33, 44, 82, 20, 56, 86}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 8; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
