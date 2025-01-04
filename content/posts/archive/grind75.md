---
title: "Grind 75"
date: 2022-11-13T16:56:59+13:00
draft: true
---

## Two Sum

```
func twoSum(nums []int, target int) []int {

	indices := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if i != j {
					indices = append(indices, []int{i, j})
				}
			}
		}
	}

	fmt.Println(indices)
	return indices[0]
}

```

## Two Parenthesis

```
package main

import (
	"container/heap"
	"fmt"
)

type StringHeap []string

func (sheap StringHeap) Len() int { return len(sheap) }

func (sheap StringHeap) Less(i, j int) bool {
	return sheap[i] > sheap[j]
}

func (sheap StringHeap) Swap(i, j int) {
	sheap[i], sheap[j] = sheap[j], sheap[i]
}

func (sheap *StringHeap) Push(heapstrf interface{}) {
	*sheap = append(*sheap, heapstrf.(string))
}

func (sheap *StringHeap) Pop() interface{} {
	var n int
	var x1 string
	var previous StringHeap = *sheap
	n = len(previous)
	x1 = previous[n-1]
	*sheap = previous[0 : n-1]
	return x1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	data := []string{"{", "}", "(", ")", "[", "]"}
	var strHeap *StringHeap = &StringHeap{}
	heap.Init(strHeap)

	for _, j := range data {

		heap.Push(strHeap, j)

		if strHeap.Len() > 0 {

			switch j {

			case "{":
				if contains(data, "}") {
					heap.Pop(strHeap)
				}
			case "[":
				if contains(data, "]") {
					heap.Pop(strHeap)
				}

			case "(":
				if contains(data, ")") {
					heap.Pop(strHeap)
				}

			case "}":
				if contains(data, "{") {
					heap.Pop(strHeap)
				}
			case "]":
				if contains(data, "[") {
					heap.Pop(strHeap)
				}

			case ")":
				if contains(data, "(") {
					heap.Pop(strHeap)
				}
			}
		}

	}

	if strHeap.Len() == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
```

### Merge Two Sorted Lists

```
package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {

	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head

	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}

}

func (ll *LinkedList) GetAt(pos int) *Node {
	ptr := ll.head
	if pos < 0 {
		return ptr
	}

	if pos > (ll.len - 1) {
		return nil
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr

}

func (ll *LinkedList) Traverse() {
	currentNode1 := ll.head
	for currentNode1 != nil {
		fmt.Println(currentNode1.value)
		currentNode1 = currentNode1.next
	}
}

func (ll *LinkedList) AddAfter(nodeProperty int, property int) {

	var node = &Node{}
	node.value = property
	node.next = nil
	var nodeWith *Node
	nodeWith = ll.GetAt(nodeProperty)
	if nodeWith != nil {
		node.next = nodeWith.next
		nodeWith.next = node
	}

}

func mergedSortedLists(l1 *LinkedList, l2 *LinkedList) LinkedList {

	mergedList := LinkedList{}
	currentNode1 := l1.head
	currentNode2 := l2.head

	mergedList.Insert(currentNode1.value)
	mergedList.Insert(currentNode2.value)

	for currentNode1 != nil && currentNode2 != nil && currentNode1.next != nil && currentNode2.next != nil {
		currentNode1 = currentNode1.next
		currentNode2 = currentNode2.next
		mergedList.Insert(currentNode1.value)
		mergedList.Insert(currentNode2.value)
	}

	return mergedList

}

func main() {

	//populate original linked lists
	l1 := &LinkedList{}
	l1.Insert(1)
	l1.Insert(2)
	l1.Insert(4)

	l2 := &LinkedList{}
	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	result := mergedSortedLists(l1, l2)
	result.Traverse()

}

//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//insert between
```