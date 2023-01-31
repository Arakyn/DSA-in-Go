package main

import "fmt"

// MaxHeap struct has  a slice which holds the array
type MaxHeap struct {
	array []int
}

// The insert method adds an element into the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// extract returns the largest key from the heap and removes it
func (h *MaxHeap) Extract() int {

	extracted := h.array[0]

	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because the lenght of the array is 0")
		return -1
	}

	// taking out the last index and placing it in the first index position
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.maxHeapifyDown(0)

	// returned the root node before putting the top one to the bottom
	return extracted
}

// func maxheapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {

	for h.array[getParent(index)] < h.array[index] {
		h.swap(getParent(index), index)
		index = getParent(index)
	}

}

// funx maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

}

// gets the parent index
func getParent(i int) int {

	return (i - 1) / 2
}

// gets the  left child index
func left(i int) int {
	return (i * 2) + 1
}

// gets the right child index
func right(i int) int {
	return (i * 2) + 2
}

// swaps keys in the index

func (h *MaxHeap) swap(parent, child int) {
	h.array[parent], h.array[child] = h.array[child], h.array[parent]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 40, 50, 700, 23}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)

}
