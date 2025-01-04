---
title: "Data Structures and Algorithms"
date: 2022-11-13T16:56:59+13:00
draft: true
---

## Chapter Two
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

## Structural design patterns 
Describe the Relationship between entities
They are used to form large structures using objects and classes.
These patterns are used to create a system with different building blocks in a flexible manner.

Gang of Four (Gof) structural design patterns
Adapter
Bridge 
Comopsite
Decorator
Facade
Flyweight
Private class data
Proxy

## Adapter (DIP)
Provides a wrapper required by an api client to incompatible types and acts as a translator between the two types.
A adapter uses the interface of a class to be to be a class with another compatible interface. 
When requirements change there are scenarios where fucntionality needs to be changed because of incompatible intefaces

target - is the inteface the client calls and invokes methods on the adapter and adaptee  
adaptee 
adapter - translates the incompatible interface implemented by the adapter
client - wants the incompatible interface implemented by the adapter 

## Bridge 

Let's say IDrawShape is an interface with the drawShape method. DrawShape implements the IDrawShape interface. We create an IContour bridge interface with the drawContour method. The contour class implements the IContour interface. The ellipse class will have a, b , r properties and drawShape (an instance of DrawShape). The ellipse class implements the contour bridge interface to implement the drawContour method. The drawContour method calls the drawShape method on the drawShape instance.


## Composite
A composite is a group of similar objects in a single object. Objects are stored in a tree form to persist the whole hierarchy. The composite pattern is used to change a hierarchical collection of objects. 

The composite pattern is modeled on a heterogeneous collection. New types of objects can be added without changing the interface and the client code.
component interface, component class, composite, and client:

Let's say there is an IComposite interface with the perform method and BranchClass that implements IComposite and has the addLeaf, addBranch, and perform methods. The Leaflet class implements IComposite with the perform method. BranchClass has a one-to-many relationship with leafs and branches. Iterating over the branch recursively, one can traverse the composite tree.

## Decorator 

In a scenario where class responsibilities are removed or added, the decorator pattern is applied. The decorator pattern helps with subclassing when modifying functionality, instead of static inheritance. An object can have multiple decorators and run-time decorators. The single responsibility principle can be achieved using a decorator. The decorator can be applied to window components and graphical object modeling. The decorator pattern helps with modifying existing instance attributes and adding new methods at run-time.

The decorator pattern participants are the component interface, the concrete component class, and the decorator class:

Let 's say IProcess is an interface with the process method. 
ProcessClass implements an interface with the process method. 
ProcessDecorator implements the process interface and has an instance of ProcessClass. 
ProcessDecorator can add more functionality than ProcessClass


## Facade 
Facade is used to abstract subsystem interfaces with a helper. The facade design pattern is used in scenarios when the number of interfaces increases and the system gets complicated. Facade is an entry point to different subsystems, and it simplifies the dependencies between the systems. The facade pattern provides an interface that hides the implementation details of the hidden code.

A loosely coupled principle can be realized with a facade pattern. You can use a facade to improve poorly designed APIs. In SOA, a service facade can be used to incorporate changes to the contract and implementation.

For example, account, customer, and transaction are the classes that have account, customer, and transaction creation methods. BranchManagerFacade can be used by the client to create an account, customer, and transaction

## Flyweight
Flyweight is used to manage the state of an object with high variation. The pattern allows us to share common parts of the object state among multiple objects, instead of each object storing it. Variable object data is referred to as extrinsic state, and the rest of the object state is intrinsic. Extrinsic data is passed to flyweight methods and will never be stored within it. Flyweight pattern helps reduce the overall memory usage and the object initializing overhead. The pattern helps create interclass relationships and lower memory to a manageable level.

The FlyWeight interface has a method through which flyweights can get and act on the extrinsic state.
ConcreteFlyWeight implements the FlyWeight interface to represent flyweight objects.
FlyweightFactory is used to create and manage flyweight objects. The client invokes FlyweightFactory to get a flyweight object. UnsharedFlyWeight can have a functionality that is not shared.
Client classes


## Private class data

The private class data pattern secures the data within a class. This pattern encapsulates the initialization of the class data. The write privileges of properties within the private class are protected, and properties are set during construction. The private class pattern prints the exposure of information by securing it in a class that retains the state. The encapsulation of class data initialization is a scenario where this pattern is applicable.

## Proxy
The proxy pattern forwards to a real object and acts as an interface to others. The proxy pattern controls access to an object and provides additional functionality. The additional functionality can be related to authentication, authorization, and providing rights of access to the resource-sensitive object. The real object need not be modified while providing additional logic. Remote, smart, virtual, and protection proxies are the scenarios where this pattern is applied. It is also used to provide an alternative to extend functionality with inheritance and object composition. A proxy object is also referred to as a surrogate, handle, or wrapper.

The proxy pattern comprises the subject interface, the RealSubject class, and the Proxy class:


## Representation of algorithms

A flow chart and pseudo code are methods of representing algorithms. An algorithm shows the logic of how a problem is solved. A flow chart has different representation symbols such as Entry, Exit, Task, Input/Output, Decision Point, and Inter Block. A structured program consists of a series of these symbols to perform a specific task. Pseudo code has documentation, action, and flow control keywords to visualize an algorithm. The documentation keywords are TASK and REM. SET, PUT, and GET are the action keywords.


The flow control keywords are SET, LOOP, (WHILE, UNTIL), REP, and POST. 

A flow chart has different representation symbols such as Entry, Exit, Task, Input/Output, Decision Point, and Inter Block. 

## Pseudo code
Pseudo code is a high-level design of a program or algorithm. Sequence and selection are two constructs used in pseudo code. Pseudo code is easier than a flow chart visualizes the algorithm while pseudo code can be easily modified and updated. Errors in design can be caught very early in pseudo code. This saves the cost of fixing defects later.

## Complexity and performance analysis

The efficiency of an algorithm is measured through various parameters, such as CPU time, memory, disk, and network. The complexity is how the algorithm scales when the number of input parameters increases. Performance is a measure of time, space, memory, and other parameters. Algorithms are compared by their processing time and resource consumption. Complexity measures the parameters and is represented by the Big O notation.

### Complexity analysis of algorithms

algorithm is a set of steps to achieve  a task 
the time taken for an algorithm to complete is the number of steps taken.

computational time grows as a factor of 10

### Big O notation
t(n) time function represents algorithm complexity based on big O notation
t(n) = o(n) - linear complexity 
logarithmic o(log n) eg binary search in a tree datastructure
quadratic time (o n2) Bubble sort, selection sort an insertion sort 
big omega and big theta are notation for the upper and lower bounds of the alogorithm
worst case, best case average case and amortised run-time complexity is used for the analysis of algorithms

Big O notation also used to determine how much space used by an algorithm 
best and worst scenarios relative to space and time 

## Linear Complexity
processing time or storage space directly proportional to the number of input elements to be processed
0(n)

```

func main (){
    var m [10]int
    var k int
    for k=0;k<10;k++{
        m[k]= k *200
        fmt.PrintF("Element [%d] = %d\n", k, m[k])
    }
}

```

## Quadratic complexity 
if the processing time is proportional to the square of the number of input elements 
complexity is 10*10=100 two loops have a max of 10 
quadratic complexity for a mutliplication table of n elements is o(n2)


```

func main (){
    var k,l int

    for k = 1; k <=10; k++{
        fmt.Println(" Multiplicatipn table",k)
        for l=1; l<=10; l++{
            var x int = l*k
            fmt.Println(x)
        }
    }
}


```

## Cubic Complexity 

In the case of cubic complexity, the processing time of an algorithm is proportional to the cube of the input elements. The complexity of the following algorithm is 10*10*10 = 1,000. The three loops have a maximum of 10. The cubic complexity for a matrix update is O(n3).

Cubic complexity O(n3) is explained in the following example:


```
func main (){
    var k,l,m int
    var arr[10][10][10]int
    for k = 0; k<10;k++{
        for l=0;l<10;l++{
            for m=0;m<10;m++{
                arr[k][l][m]=1 
                 fmt.Println("Element Value",k,l,m,"is", arr[k][l][m])
            }    
        }
    }
}

```

## Logarithmic complexity
An algorithm is of logarithmic complexity if the processing time is proportional to the logarithm of the input elements. The logarithm base is typically 2. The following tree is a binary tree with LeftNode and RightNode. The insert operation is of O(log n) complexity, where n is the number of nodes.

Logarithmic complexity is presented as follows:

```
See tree example 
```

### Brute force algorithms
A brute force algorithm solves a problem based on the statement and the problem definition. Brute force algorithms for search and sort are sequential search and selection sort. Exhaustive search is another brute force algorithm where the solution is in a set of candidate solutions with definitive properties. The space in which the search happens is a state and combinatorial space, which consists of permutations, combinations, or subsets.

Brute Force algorithms are known for wide applicability and simplicity in solving complex problems. Searching, string matching, and matrix multiplication are some scenarios where they are used. Single computational tasks can be solved using brute force algorithms. They do not provide efficient algorithms. The algorithms are slow and non-performant. Representation of a brute force algorithm is shown in the following code:

```
func findElement(arr[10] int, k int) bool {
    var i int
    for i=0; i< 10; i++ {
        if arr[i]==k {
            return true
        }
    }
    return false
}
// main method
func main() {
    var arr = [10]int{1,4,7,8,3,9,2,4,1,8}
    var check bool = findElement(arr,10)
    fmt.Println(check)
    var check2 bool = findElement(arr,9)
    fmt.Println(check2)
}
```

## Divide and conquer algorithms 
The approach is to recursively solve the sub-problems and merge the solutions of the sub-problems.
Recursion, quick sort, binary search, fast Fourier transform, and merge sort are good examples of divide and conquer algorithms. Memory is efficiently used with these algorithms. Performance is sometimes an issue in the case of recursion. On multiprocessor machines, these algorithms can be executed on different processors after breaking them down into sub-problems. A divide and conquer algorithm is shown in the following code:

```
func fibonacci(K int) int{
    if k<=1{
        return 1
    }
}

func main(){
    var m int =5
    for m=0; m<8; m++{
        var fib = finonacci(m)
        fmt.printLn(fib)
    }
}
```

## Backtracking algorithms

A backtracking algorithm solves a problem by constructing the solution incrementally. Multiple options are evaluated, and the algorithm chooses to go to the next component of the solution through recursion.

Backtracking is an algorithm that finds candidate solutions and rejects a candidate on the basis of its feasibility and validity. Backtracking is useful in scenarios such as finding a value in an unordered table. It is faster than a brute force algorithm, which rejects a large number of solutions in an iteration. Constraint satisfaction problems such as parsing, rules engine, knapsack problems, and combinatorial optimization are solved using backtracking.

The following is an example of a backtracking algorithm. The problem is to identify the combinations of elements in an array of 10 elements whose sum is equal to 18. The findElementsWithSum method recursively tries to find the combination. Whenever the sum goes beyond the k target, it backtracks:

```
//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main
// importing fmt package
import (
 "fmt"
)
//findElementsWithSum of k from arr of size

func findElementsWithSum(arr[10] int,combinations[19] int,size int, k int, addValue int, l int, m int) int {
var num int = 0
if addValue > k {
 return -1
 }
if addValue == k {
 num = num +1
 var p int =0
 for p=0; p < m; p++ {
  fmt.Printf("%d,",arr[combinations[p]])
  }
 fmt.Println(" ")
 }
var i int
for i=l; i< size; i++ {
//fmt.Println(" m", m)
combinations[m] = l
findElementsWithSum(arr,combinations,size,k,addValue+arr[i],l,m+1)
l = l+1
 }
 return num
}
// main method
func main() {
var arr = [10]int{1,4,7,8,3,9,2,4,1,8}
var addedSum int = 18
var combinations [19]int
findElementsWithSum(arr,combinations,10,addedSum,0,0,0)
//fmt.Println(check)//var check2 bool = findElement(arr,9)
//fmt.Println(check2)
}


```
This chapter covered the definition of abstract datatypes, classifying data structures into linear, nonlinear, homogeneous, heterogeneous, and dynamic types. Abstract datatypes such as container, list, set, map, graph, stack, and queue were presented in this chapter. The chapter covered the performance analysis of data structures and structural design patterns.

We looked at the classification of data structures and structural design patterns. You can use algorithms such as brute force, divide and conquer, and backtracking by calculating the complexity and performance analysis. The choice of algorithm and the use of design patterns and data structures are the key takeaways.

## Chapter Three 

In this chapter, we will discuss the following Go language-specific data structures:

Arrays
Slices
Two-dimensional slices
Maps
Database operations
Variadic functions
CRUD web forms

## Chapter Five - Linear Data Structures

Various applications, such as Facebook, Twitter, and Google, use lists and linear data structures. As we have discussed previously, data structures allow us to organize vast swathes of data in a sequential and organized manner, thereby reducing time and effort in working with such data. Lists, stacks, sets, and tuples are some of the commonly used linear data structures.

### Lists
A list is a collection of ordered elements that are used to store list of items. Unlike array lists, these can expand and shrink dynamically.
Lists also be used as a base for other data structures, such as stack and queue.

- Linked List 
is a sequence of nodes that have properties and a reference to the next node in the sequence. 
The data structure permits the addition and deletion of components from any node next to another node.
Linked list will have a set of nodes with integer properties, as follows:

```
//Node class

type Node struct {
    property int
    nextNode *Node
}

// LinkedList class
type LinkedList struct {
    headNode *Node
}

//AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
    var node = Node{}
    node.property = property
    if node.nextNode != nil {
        node.nextNode = linkedList.headNode
    }
    linkedList.headNode = &node
}

// main method
func main() {
    var linkedList LinkedList
    linkedList = LinkedList{}
    linkedList.AddToHead(1)
    linkedList.AddToHead(3)
    fmt.Println(linkedList.headNode.property)
}

//IterateList method iterates over LinkedList
func (linkedList *LinkedList) IterateList() {
    var node *Node
    for node = linkedList.headNode; node != nil; node = node.nextNode {
        fmt.Println(node.property)
    }
}

//LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node{
 var node *Node
 var lastNode *Node
 for node = linkedList.headNode; node != nil; node = node.nextNode {
 if node.nextNode ==nil {
 lastNode = node
 }
 }
 return lastNode
}

//AddToEnd method adds the node with property to the end

func (linkedList *LinkedList) AddToEnd(property int) {
 var node = &Node{}
 node.property = property
 node.nextNode = nil
 var lastNode *Node
 lastNode = linkedList.LastNode()
 if lastNode != nil {
 lastNode.nextNode = node
 }
}

//NodeWithValue method returns Node given parameter property

func (linkedList *LinkedList) NodeWithValue(property int) *Node{
 var node *Node
 var nodeWith *Node
 for node = linkedList.headNode; node != nil; node = node.nextNode {
 if node.property == property {
 nodeWith = node
 break;
 }
 }
 return nodeWith
}

//AddAfter method adds a node with nodeProperty after node with property

func (linkedList *LinkedList) AddAfter(nodeProperty int,property int) {
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

// main method
func main() {
    var linkedList LinkedList
    linkedList = LinkedList{}
    linkedList.AddToHead(1)
    linkedList.AddToHead(3)
    linkedList.AddToEnd(5)
    linkedList.AddAfter(1,7)
    linkedList.IterateList()
}

```

- Doubly Linked List
In a doubly linked list, all nodes have a pointer to the node they are connected to, on either side of them, in the list.
This means that each node is connected to two nodes, and we can traverse forward through to the next node or backward through to the previous node. 
Doubly linked lists allow insertion, deletion and, obviously, traversing operations. 

```
// Node class
type Node struct {
    property int
    nextNode *Node
    previousNode *Node
}

//NodeBetweenValues method of LinkedList
func (linkedList *LinkedList) NodeBetweenValues(firstProperty int,secondProperty int) *Node{
    var node *Node
    var nodeWith *Node
    for node = linkedList.headNode; node != nil; node = node.nextNode {
        if node.previousNode != nil && node.nextNode != nil {
            if node.previousNode.property == firstProperty && node.nextNode.property ==    
            secondProperty{
               nodeWith = node
               break;
            }
        }
    }
    return nodeWith
}

//AddToHead method of LinkedList
func (linkedList *LinkedList) AddToHead(property int) {
 var node = &Node{}
 node.property = property
 node.nextNode = nil
 if linkedList.headNode != nil {
 node.nextNode = linkedList.headNode
 linkedList.headNode.previousNode = node
 }
 linkedList.headNode = node
}

//AddAfter method of LinkedList
func (linkedList *LinkedList) AddAfter(nodeProperty int,property int) {
 var node = &Node{}
 node.property = property
 node.nextNode = nil
 var nodeWith *Node
 nodeWith = linkedList.NodeWithValue(nodeProperty)
 if nodeWith != nil {
 
 node.nextNode = nodeWith.nextNode
 node.previousNode = nodeWith
 nodeWith.nextNode = node
 }
}

//AddToEnd method of LinkedList
func (linkedList *LinkedList) AddToEnd(property int) {
 var node = &Node{}
 node.property = property
 node.nextNode = nil
 var lastNode *Node
 lastNode = linkedList.LastNode()
 if lastNode != nil {
 
 lastNode.nextNode = node
 node.previousNode = lastNode
 }
}

// main method
func main() {
 var linkedList LinkedList
 linkedList = LinkedList{}
 linkedList.AddToHead(1)
 linkedList.AddToHead(3) linkedList.AddToEnd(5)
 linkedList.AddAfter(1,7)
 fmt.Println(linkedList.headNode.property)
 var node *Node
 node = linkedList.NodeBetweenValues(1,5)
 fmt.Println(node.property)
}

```

### Sets 
A Set is a linear data structure that has a collection of values that are not repeated. A set can store unique values without any particular order. 
Dynamic and mutable sets allow for the insertion and deletion of elements.
Algebraic operations such as union, intersection, difference, and subset can be defined on the sets. 
The following example shows the Set integer with a map integer key and bool as a value:

```
package main
// importing fmt package
import (
 "fmt"
)
//Set class
type Set struct {
 integerMap map[int]bool
}

//create the map of integer and bool
func (set *Set) New(){
 set.integerMap = make(map[int]bool)
}

// adds the element to the set
func (set *Set) AddElement(element int){
 if !set.ContainsElement(element) {
  set.integerMap[element] = true
 }
}



```

### Tuples

Tuples are finite ordered sequences of objects. They can contain a mixture of other data types and are used to group related data into a data structure. In a relational database, a tuple is a row of a table. Tuples have a fixed size compared to lists, and are also faster. A finite set of tuples in the relational database is referred to as a relation instance. A tuple can be assigned in a single statement, which is useful for swapping values. Lists usually contain values of the same data type, while tuples contain different data.

### Queues

A queue consists of elements to be processed in a particular order or based on priority. A priority-based queue of orders is shown in the following code, structured as a heap. Operations such as enqueue, dequeue, and peek can be performed on queue. A queue is a linear data structure and a sequential collection. Elements are added to the end and are removed from the start of the collection. Queues are commonly used for storing tasks that need to be done, or incoming HTTP requests that need to be processed by a server. In real life, handling interruptions in real-time systems, call handling, and CPU task scheduling are good examples for using queues.

//new //add method order parameter added to queue.

Synchronized queue
A synchronized queue consists of elements that need to be processed in a particular sequence. Passenger queue and ticket processing queues are types of synchronized queues.

### Stacks
A stack is a last in, first out structure in which items are added from the top. Stacks are used in parsers for solving maze algorithms. Push, pop, top, and get size are the typical operations that are allowed on stack data structures. Syntax parsing, backtracking, and compiling time memory management are some real-life scenarios where stacks can be used. An example of stack implementation is as follows (stack.go):

## Chapter 6
Non Linear Data Structures 
Non-linear data structures are used in cryptography and other areas. A non-linear data structure is an arrangement in which an element is connected to many elements. These structures use memory quickly and efficiently. Free contiguous memory is not required for adding new elements.

The length of the data structures is not important before adding new elements. A non-linear data structure has multiple levels and a linear one has a single level. The values of the elements are not organized in a non-linear data structure. The data elements in a non-linear data structure cannot be iterated in one step. The implementation of these data structures is complicated.

Tree types such as binary search trees, treaps, and symbol tables are explained in this chapter.

This chapter covers the following non-linear data structures:

Trees
Tables
Containers
Hash functions


## Trees
A tree is a non-linear data structure. Trees are used for search and other use cases. A binary tree has nodes that have a maximum of two children. A binary search tree consists of nodes where the property values of the left node are less than the property values of the right node. The root node is at level zero of a tree. Each child node could be a leaf.

### Binary Search Tree
A binary search tree is a data structure that allows for the quick lookup, addition, and removal of elements. It stores the keys in a sorted order to enable a faster lookup. This data structure was invented by P. F. Windley, A. D. Booth, A. J. T. Colin, and T. N. Hibbard. On average, space usage for a binary search tree is of the order O(n), whereas the insert, search, and delete operations are of the order O(log n). A binary search tree consists of nodes with properties or attributes:

key integer
value integer
leftnode and right node instances of the treenode

### AVL Tree
height adjusting binary search trees 
The balance factor is obtained by finding the difference between the heights of the left and right sub-trees. Balancing is done using rotation techniques. If the balance factor is greater than one, rotation shifts the nodes to the opposite of the left or right sub-trees. The search, addition, and deletion operations are processed in the order of O(log n).

### B+ Tree
### B- Tree
### T-Tree


### Tables

A table class has an array of rows and column names

### Symbol Tables
A symbol table is present in memory during the program translation process. It can be present in program binaries. A symbol table contains the symbol's name, location, and address. In Go, the gosym package implements access to the Go symbol and line number tables. Go binaries generated by the GC compilers have the symbol and line number tables. A line table is a data structure that maps program counters to line numbers.


### Containers
The containers package provides access to the heap, list, and ring functionalities in Go. Containers are used in social networks, knowledge graphs, and other areas. Containers are lists, maps, slices, channels, heaps, queues, and treaps. 

### Ring is a circular linked list 
container ring 

### Hash functions
Hash functions are used in cryptography and other areas. These data structures are presented with code examples related to cryptography. There are two ways to implement a hash function in Go: with crc32 or sha256. Marshaling (changing the string to an encoded form) saves the internal state, which is used for other purposes later. A BinaryMarshaler (converting the string into binary form) example is explained in this section:

## Chapter 7
### Homogenous Data Structures

-simular types of data 

- Two-dimensional arrays
- Multi-dimensional arrays

Matrix representation
Multiplication
Addition
Subtraction
Determinant calculation
Inversion
Transposition

## Chapter 8
### Hetrogeneous data structures 

Hetrogeneous data structures contain diverse types of data such as integers, doubles and floats
linked lists and ordered lists used for memory management 
linked list is a chain of elements that are associated together by means of pointers 
memory utilised can be allocated dynamically dont have to take a block of memory 

linked lists
ordered lists
unorderered lists

singly linked list dynamic data structure in which additional and removal operations are easy
random retrieval not possible as youll have to travers the list of nodes for a positioned node

doubly linked list
data structure that consisits of nodes that have links to previous and next nodes
circular linked lists last node connected to the first


Ordered list: By creating a group of methods for the slice data type and calling sort

Unordered list: The other way is to invoke sort.Slice with a custom less function
sort by factor
mulitsorter

unordered list
relative positions in contiguous memory do not need to be maintained

## Chapter 9 
### Dynamic Data Structures 

```
A dynamic data structure is a set of elements in memory that has the adaptability to expand or shrink. This ability empowers a software engineer to control precisely how much memory is used. Dynamic data structures are used for handling generic data in a key-value store. They can be used in distributed caching and storage management. Dynamic data structures are valuable in many circumstances in which dynamic addition or deletion of elements occur. They are comparable in capacity to a smaller relational database or an in-memory database. These data structures are used in marketing and customer relationship management applications. Dictionaries, TreeSets, and sequences are examples of dynamic data structures.
```

### Dictionaries 

Collection of Unique Key Value Pairs 
Distributed Caching
In Memory Databases 
Keys can be any type 
can have duplicate values
add, remove and lookup operations 

```

type DictVal string 

type Dictionary Struct{
    elements map[DictKey]DictVal
    lock sync.RWMutex
}

// Put method
func (dict *Dictionary) Put(key DictKey, value DictVal) {
    dict.lock.Lock()
    defer dict.lock.Unlock()
    if dict.elements == nil {
        dict.elements = make(map[DictKey]DictVal)
    }
    dict.elements[key] = value
}

// Remove method
func (dict *Dictionary) Remove(key DictKey) bool {
    dict.lock.Lock()
    defer dict.lock.Unlock()
    var exists bool
    _, exists = dict.elements[key]
    if exists {
        delete(dict.elements, key)
    }
    return exists
}

// Contains method
func (dict *Dictionary) Contains(key DictKey) bool {
    dict.lock.RLock()
    defer dict.lock.RUnlock()
    var exists bool
    _, exists = dict.elements[key]
    return exists
}

// Find method
func (dict *Dictionary) Find(key DictKey) DictVal {
    dict.lock.RLock()
    defer dict.lock.RUnlock()
    return dict.elements[key]
}

// Reset method
func (dict *Dictionary) Reset() {
    dict.lock.Lock()
    defer dict.lock.Unlock()
    dict.elements = make(map[DictKey]DictVal)
}

// NumberOfElements method
func (dict *Dictionary) NumberOfElements() int {
    dict.lock.RLock()
    defer dict.lock.RUnlock()
    return len(dict.elements)
}

// GetKeys method
func (dict *Dictionary) GetKeys() []DictKey {
    dict.lock.RLock()
    defer dict.lock.RUnlock()
    var dictKeys []DictKey
    dictKeys = []DictKey{}
    var key DictKey
    for key = range dict.elements {
        dictKeys = append(dictKeys, key)
    }
    return dictKeys
}

// GetValues method 
func (dict *Dictionary) GetValues() []DictVal {
    dict.lock.RLock()
    defer dict.lock.RUnlock()
    var dictValues []DictVal
    dictValues = []DictVal{}
    var key DictKey
    for key = range dict.elements {
        dictValues = append(dictValues, dict.elements[key])
    }
    return dictValues
}

// main method
func main() {
  var dict *Dictionary = &Dictionary{}
  dict.Put("1","1")
  dict.Put("2","2")
  dict.Put("3","3")
  dict.Put("4","4")
  fmt.Println(dict)
}

```

### TreeSets
TreeSets are used in marketing and customer relationship management applications. 
TreeSet is a set that has a binary tree with unique elements. 
The elements are sorted in a natural order. 
In the following code snippet, TreeSet creation, insertion, search, and stringify operations are presented. 
TreeSet allows only one null value if the set is empty. 
The elements are sorted and stored as elements. 
The add, remove, and contains functions cost log(n) on TreeSets:

```
package main

// TreeSet class
type TreeSet struct {
  bst *BinarySearchTree
}

Insert TreeNode Method

// InsertTreeNode method
func (treeset *TreeSet) InsertTreeNode(treeNodes ...TreeNode) {
  var treeNode TreeNode
  for _, treeNode = range treeNodes {
    treeset.bst.InsertElement(treeNode.key, treeNode.value)
  }
}
//Delete Method
//inOrderTraverseTree Method - Traverses from the leftto the root and the to the right
//preorderTraverseTree Method - method traverses the tree from the root, to the left and right of the tree. The preOrderTraverseTree method takes treeNode and function as parameters. If treeNode is not nil, function is invoked with the value of treeNode, and the preOrderTraverseTree method is invoked with function and leftNode and rightNode as parameters
//search method -The Search method of the TreeSet class takes a variable argument named treeNodes of the TreeNode type and returns true if one of those treeNodes exists; otherwise, it returns false. The code following snippet outlines the Search method:
string method contains a string version of bst.


Synchronized TreeSets
Operations that are performed on synchronized TreeSets are synchronized across multiple calls that access the elements of TreeSets. Synchronization in TreeSets is achieved using a sync.RWMutex lock. The lock method on the tree's lock instance is invoked, and the unlock method is deferred before the tree nodes are inserted, deleted, or updated:


Mutable TreeSets
Mutable TreeSets can use add, update, and delete operations on the tree and its nodes. insertTreeNode updates the tree by taking the rootNode and treeNode parameters to be updated. The following code snippet shows how to insert a TreeNode with a given rootNode and TreeNode:

reeset.bst
The TreeNode's can be updated by accessing treeset.bst and traversing the binary search tree from the rootNode and the left and right nodes of rootNode, as shown here:


  var treeset *TreeSet = &TreeSet{}
  treeset.bst = &BinarySearchTree{} 
  var node1 TreeNode = TreeNode{8, 8, nil, nil}
  var node2 TreeNode = TreeNode{3, 3, nil, nil}
  var node3 TreeNode = TreeNode{10, 10, nil, nil}
  var node4 TreeNode = TreeNode{1, 1, nil, nil}
  var node5 TreeNode = TreeNode{6, 6, nil, nil}
  treeset.InsertTreeNode(node1, node2, node3, node4, node5)
  treeset.String()

Sequences
A sequence is a set of numbers that are grouped in a particular order. The number of elements in the stream can be infinite, and these sequences are called streams. A subsequence is a sequence that's created from another sequence. The relative positions of the elements in a subsequence will remain the same after deleting some of the elements in a sequence.

In the following sections, we will take a look at different sequences such as the Farey sequence, Fibonacci sequence, look-and-say, and Thue–Morse.

Farey sequence
A Farey sequence consists of reduced fractions with values between zero and one. The denominators of the fractions are less than or equal to m, and organized in ascending order. This sequence is called a Farey series. In the following code, reduced fractions are displayed:

Fibonacci sequence
The Fibonacci sequence consists of a list of numbers in which every number is the sum of the two preceding numbers. Pingala, in 200 BC, was the first to come up with Fibonacci numbers. The Fibonacci sequence is as follows:

Look-and-say
The look-and-say sequence is a sequence of integers:

The sequence is generated by counting the digits of the previous number in the group. John Conway initially coined the term look-and-say sequence.

The look-and-say sequence is shown in the following code. The look_say method takes a string as a parameter and returns a look-and-say sequence of integers:

The Thue–Morse sequence is a binary sequence starting at zero that appends the Boolean complement of the current sequence.

The Thue–Morse sequence is as follows:


The Thue–Morse sequence was applied by Eugene Prophet and used by Axel Thue in the study of combinatorics on words. The Thue–Morse sequence is used in the area of fractal curves, such as Koch snowflakes.

The following code snippet creates the Thue–Morse sequence. The ThueMorseSequence function takes a bytes.Buffer instance buffer and modifies the buffer to the Thue–Morse sequence by applying the complement operation on the bytes:


## Classic Algorithms

Canonicalize Data and create readable content 
Sorting - asc/ desc arrangement 
Searching - find element in the set
Recursing - calls itself with input elements 
Hashing - cryptographic hash technique - maps data with a subject size to a hash with a seettled size, single direction and you cannot alter 

### Sorting Algorithms

best soritng algorithm time is O(n log n)

Sorting algorithms classified by 
Computational Complexity
Memory Usage 
Stability
Types of Sorting: serial or parallell
Adaptability
Method of Sorting


Bubble - compares neighbouring elements and swaps them if they are in wrong order
O(n2) where n is the nof of elements sorted
smallest or greatest value bubbles up to the top of the collection
depending on whether you are sorting in asc or desc order


## Selectiom Sort 

Divides the input collection into two fragments 
subset of elements is sorted by swapping the smallest or leargets elemnt from the left of the lis to the right 
on2
ineffcient for large collections  perfoms worse than the insertion sort algo rithm.


## Insertion














```