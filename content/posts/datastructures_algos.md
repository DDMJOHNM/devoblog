---
title: "Data Structures and Algorithms"
date: 2022-11-13T16:56:59+13:00
draft: false
---

# Learn Data Structures and Algorithms with Golang
Data structure organization of a computers memory, in order to retrive it quickly for processing. 

## Classification of data structures.
Linear -Lists, sets, tuples, queues, stacks, and heaps.

### Doubly Linked List
```
package main

import (
	"fmt"
)

//connected to two nodes can traverse backward and forward

type Node struct{
	property int
	nextNode *Node
	previousNode *Node	

}

type LinkedList struct{
	headNode *Node
}


func (linkedList *LinkedList) AddToHead(property int){

}

func (linkedList *LinkedList) IterateList(){
	
}

func (linkedList * LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node{

} 

func main(){

	//add to head method
	linkedList := LinkedList{}
	
	
}
```

### Heaps

```
package main

import (
	"container/heap"
	"fmt"
)

//Heaps - data structure based on a heap property
//used in selection, graph and kway merge alogrithms
//operations such as finding, insertion keys changes and deleting all perfromed on heaps
//heap order value is stored at each node is greater than equal to its children
//order descending max heap otherwise minimum heap
//not sorted but partial sorted
//heap structure, heap sorting alogrithm
//a php example https://www.sitepoint.com/data-structures-3/
//https://www.w3resource.com/php-exercises/searching-and-sorting-algorithm/searching-and-sorting-algorithm-exercise-2.php

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] > iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

//interheap method pushes the item
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

//interheap method pops the item from the heap

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minumum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("d% \n", heap.Pop(intHeap))
	}

}
```

### Linked List

```
package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property

	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {

	var node *Node
	var nodeWith *Node

	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}

}

func main() {

	//add to head method
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	//fmt.Println(linkedList.headNode.property)

	//iterate list method
	//linkedList.IterateList();
	//fmt.Println(linkedList.NodeWithValue(3));
	// /linkedList.AddAfter(4, 5)
	linkedList.IterateList()

}

```

### List 
```
package main

import (
	"fmt"
	"container/list"
)

func main() {

//doubly linked list can be traversed in both directions	
	
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
```

### Tuples

```
package main

import (
	"fmt"
)

//Tuples - finite list of sorted items  


func powerSeries(a int)(int,int){

	return a*a, a*a*a
}


func main(){

	var square int
	var cube int 
	
	square, cube = powerSeries(3)
	fmt.Print("square", square, "cube", cube)  

}
```

Non Linear - Trees, tables, and containers.

Honogenious - Two-dimensional and multidimensional.

Hetrogenious - Linked Lists, ordered lists and unordered lists.

Dynamic - dictionaries, tree sets, and sequences.

## Structural design patterns
Define relationships between entities, create system with different system blocks in a flexibile manner Adapter, bridge, composite, decorator, facade, flyweight, private class data, and proxy are the Gang of Four (GoF) structural design patterns.

- Adapter (DIP, Delegation) - principles adhered to
The adapter uses the interface of a class to be a class with another compatible interface.
The adapter translates the incompatible interface of the adaptee into an interface that the client wants.

```
package main

import(
    "fmt"
) 

type IProcess interface {
    process()
}

type Adapter struct {
    adaptee Adaptee
}    

func (adapter Adapter) process() {
    fmt.Println("adpater process")
    adapter.adaptee.convert()
}

type Adaptee struct{
    adapterType int 
}

func (adaptee Adaptee) convert(){
    fmt.Println("Adaptee convert method")
}

func main(){
    var processor IProcess = Adapter{}
    processor.process();    
}
```

## Chapter One
Classification of data structures and structural design patterns
Representation of algorithms
Complexity and performance analysis
Brute force algorithms
Divide and conquer algorithms
Backtracking algorithms


- Reduce storage space and differculty while performing specific tasks
- Data Structures (Organization of Data ) - handle large amount of data in various fields such as database mgmt and internet indexing services

### Abstract Data types
Classify into linear, non linear, homgeneuous, hetrogeneous and dynamic types

Abstract datatypes
Container 
List 
Set
Map 
Graph 
Stack 
Queue

1. basic alogorithms using the the right data structure and structural design patterns
2. performance analysis

Time and space analysis for different algorithms helps compare them and use the optimal one 

## Classification of Data Structure 
Orgnization of Data into a Computer MemoryScheme for Data to decouple fucntional defination of Data from its implementation.

### Linear
Lists, Sets, Tuples, Queues, Stacks, heaps

### Non-Linear 
Trees, Tables, Containers

### Honegenous
2D Arrays, Multi Dimensional Arrays 

### Hetrogeneous
Linked Lists, Ordered Lists, Unordered Lists

### Dynamic 
Dictionaries, TreeSets, Sequences


## Lists 
A sequence of elements, connected with each other forward or backward direction, element can have payload properties 

## Tuples
A is a finite sorted list of elements
groups data immutable sequential collections 
element has a related field of different datatypes

# Heaps
A heap is a data structure based on the heap property 
used in selection, graph and k way merge alogrithms
finding, merging, insertion and key changes and deleting are performed on heaps
according to the heap order (maximum heap property) value is stored at each node is greater than or equal to it's children.

Order is descending then referred to as a maximum heap otherwise minimum heap
partially ordered datastructure 










