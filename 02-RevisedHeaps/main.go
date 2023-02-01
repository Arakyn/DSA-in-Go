package main

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) extract() int {

	extracted := h.array[0]

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because the lenght of array is 0")
		return -1
	}

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	return extracted

}

func (h *Heap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childtocompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childtocompare = l
		} else if h.array[l] > h.array[r] {
			childtocompare = l

		} else {

			childtocompare = r
		}
		if h.array[index] < h.array[childtocompare] {
			h.swapper(index, childtocompare)
			index = childtocompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}

}

func (h *Heap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

func getParent(index int) int {

	return (index - 1) / 2
}
func left(parent int) int {
	return (parent * 2) + 1
}
func right(parent int) int {
	return (parent * 2) + 2
}

func (h *Heap) swapper(parent int, child int) {

	h.array[parent], h.array[child] = h.array[child], h.array[parent]

}

func (h *Heap) maxHeapifyUp(index int) {
	for h.array[getParent(index)] < h.array[index] {
		h.swapper(getParent(index), index)
		index = getParent(index)

	}

}

func main() {

	a := *&Heap{}
	array := []int{10, 20, 30, 40, 50}

	for _, v := range array {
		a.insert(v)
	}
	fmt.Println(a)

}
