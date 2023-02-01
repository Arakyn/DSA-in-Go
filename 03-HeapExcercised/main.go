package main

import "fmt"

type MaxHeap struct {
	array []int
}

// actual functions
// insert adds the key into the array and sorts it out
func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyDown(len(h.array) - 1)

}

// this is the actual function that sorts the last key to its place
func (h *MaxHeap) MaxHeapifyDown(index int) {

	for h.array[getParent(index)] < h.array[index] {

		h.swap(getParent(index), index)
		index = getParent(index)
	}

}

// func extract retracts the first index key and puts the last index key in the first index position and sorts it out to its right place
func (h *MaxHeap) extract() int {
	extracted := h.array[0]
	if len(h.array) == 0 {
		fmt.Println("The lenght of the array is 0 hence cannot extract anything")
		return -1
	}
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	return extracted

}

// actual extraction function
func (h *MaxHeap) maxHeapifyUp(index int) {

	LastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childtocompare := 0

	for l <= LastIndex {
		if l == LastIndex {
			childtocompare = l
		} else if h.array[l] > h.array[r] {
			childtocompare = l

		} else if h.array[l] < h.array[r] {
			childtocompare = r
		}

		if h.array[index] < h.array[childtocompare] {
			h.swap(index, childtocompare)
		} else {
			return
		}
	}

}

// helper functions
// swaps the parent key with the child key
func (h *MaxHeap) swap(parent int, child int) {

	h.array[parent], h.array[child] = h.array[child], h.array[parent]
}

// get parent gets the parent index

func getParent(index int) int {
	return (index - 1) / 2

}

// left gets the left index key

func left(index int) int {
	return (index * 2) + 1
}

// right gets the right index key

func right(index int) int {
	return (index * 2) + 2
}

func main() {

	m := &MaxHeap{}
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for _, v := range array {
		m.insert(v)
		fmt.Println(m.array)
	}

}
