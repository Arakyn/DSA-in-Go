package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParent(index)] < h.array[index] {
		h.swap(getParent(index), index)
		index = getParent(index)

	}
}
func getParent(index int) int {

	return (index - 1) / 2
}
func (h *MaxHeap) swap(parent, child int) {
	h.array[parent], h.array[child] = h.array[child], h.array[parent]
}

func main() {
	m := &MaxHeap{}
	buildingHeap := []int{1, 2, 3, 4, 5, 6}
	for _, v := range buildingHeap {
		m.Insert(v)
	}
	fmt.Println(m)

}
