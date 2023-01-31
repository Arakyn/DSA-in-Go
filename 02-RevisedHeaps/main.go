package main

import "fmt"

type Heap struct {
	array []int
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
