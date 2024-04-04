---
title: "Data Structures and Algorithms"
date: 2022-11-13T16:56:59+13:00
draft: false
---

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

