package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

// Len implements heap.Interface.
func (iheap *IntegerHeap) Len() int {
	return len(*iheap)
}

// Less implements heap.Interface.
func (*IntegerHeap) Less(i int, j int) bool {
	panic("unimplemented")
}

// Pop implements heap.Interface.
func (*IntegerHeap) Pop() any {
	panic("unimplemented")
}

// Push implements heap.Interface.
func (*IntegerHeap) Push(x any) {
	panic("unimplemented")
}

// Swap implements heap.Interface.
func (iheap IntegerHeap) Swap(i int, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func main() {

	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
