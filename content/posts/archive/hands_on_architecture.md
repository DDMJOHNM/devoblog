---
title: "Hands On Arcitecture"
date: 2023-04-06T16:56:59+13:00
draft: true
---

# Hands on Architecture 

## Role Of the Architect
- Define a blueprint

## Requirements Clarification 
- Top level functional and non functional requirements (system qualities)  

## True North   
- Engineering Principles for the system
- High level design (Decomposition of system into components)
- Quality Attributes 
- Product Velocity (cicd)
- a/b testing - every feature has a flag 

## Technology Selection 

## Leadership 
(Steering and asking tough questions at design meetings)

## Coaching and Mentoring Developers 
- Outside of their normal deliverables

## Target State vs Current State
- mvp 
- next bite sized chunks after that 

## Software 

## Architecture vs Design 
- High Level Structure 
- Low Level Implementatiion Details 

## What does Architecture look like?
- High Cohesion (Component Performs Single Task)
- Low Coupling (low dependancy between themselves) 
(Component can be extended or completely replaceable)

### See Robert Martin Concentric Circles
- Enterprise Business Rules
- Application Business Rules
- Inteface Adapters
- Frameworks and Drivers 

Different Architectures will bring out different numbers and sets of circles.

## Main Architectural Paradigms that are commonly used.

- Package Based
- Layering/N-tier/3-tier
- Async / message-bus / actor model / Communicating Sequential Processes(CSP)
- Object Oriented
- MVC/Separated Presentation
- Microservices/Service Oriented Architecture (SOA)

## Microservices 

A simple standalone application 

Three distinct layers:
- Front end 
- Back End 
- Data Store

As product features start to multiply 
- Code complexity increases 
- Operational work increases (Dev Ops)
- Product Hits Scalability Limits
- Developers write huge amount of tests as quality gates

Applications in this state are called 'Monoliths' 

Each microservice is autonomous, self contained and individually deployable and scalable.

Each Feature can be a microservice.

## Challenges for Microservices

- Efficientcy.
- Programming complexity. (Compartmentalisation key to managing)  
- Synchronisation Primatives - mutexes and Semaphores (Deadlocks)
- Priority Inversion - High priority process wait on Low Priority Process
- Starvation 


## Golang to solve these challenges

- Code becomes hard to read and poorly documented. 

- Contracts between components cannot be easily inferred.

- Builds are slow. The development cycles of code-compile-test grow in difficulty,with inefficiency in modeling concurrent systems, as writing efficient code with synchronization primitives is tough.

- Manual memory management often leads to bugs.

- There are uncontrolled dependencies.

- There is a variety of programming styles due to multiple ways of doing
something, leading to difficulty in code reviews, among other things.


## Concurrency

With shared memory its easy to get deadlocks or corruption if a process crashes or misbehaves.

CSP promotes messaging using channels queues with a logical inteface of send() and recv()   

First class channels. (First Class Go Primatives)
Proceedures called go routines.

## Garbage Collection 
In a concurrent system garbage collection is a must have feature 

## Object Orientation

Go way is composition over inheritance
For Polymorphic Behavior Go uses interfaces and duck typing.

## Pacakging Code 
encapsulation—packaging code into a higher level of abstractions
- Object orientation in go
- Code layout of packages, dependencies

### Contracts key aspect of module design  
Formalized Documentation documenetation of an interaction with a software component.
(Durable, Versioned, Service Level Agreements)

## Object Orientation in Go
- Class is a blueprint or template for objects
- Functions invoke behavior
- Encapsulation - implies exposing a contract (inteface) for the behavior of objects and hides implementation details. (Private attributes and methods)

- Inheritance (is-a relationship) can lead to a hieraechy of classes that can be hard to debug.

```
This ability of an interface method to behave differently based on the actual object is called polymorphism (derived classes <- Super class)

An alternative to inheritance is to delegate behavior, also called composition has-a relationship

Classes implement an interface—which is the contract the base class offers.
Functionality reuse happens through having references to objects, rather than
deriving from classes.

Golang favours composition over inheritance.
```
---
* THe Struct - equivelent container for encapsulation.
Methods hava a receiver clause (Method Set)
* pointer receiver - pass by refernce 
- non - pointer receiver pass by reference 
- slices and maps act as a reference pass them as a value will allow mutation of the objects.

* Visibility - struct field with a lowercase letter is private.
  (Public struct field code outside package can reference.) 
* The Interface - Interface construct key to polymorhism in go 
Objects that implement all the interface methods
automatically implement the interface.

* Embedding - is a mechanism to allow the ability to borrow pieces from different classes.
Let's call Base , struct embedded into a Derived struct.
---

This may feel like inheritance, but embedding does not provide polymorphism. Embedding
differs from subclassing in an important way: when a type is embedded, the methods of
that type are available as methods of the outer type; however, for invocation of the
embedded struct methods, the receiver of the method must be the inner (embedded)
type, not the outer one
(Can also be done on interfaces)
multiple inheritences -deadly diamond of death

## Code Layout 
bin
pkg
<package>
vendor
Makefile
scripts
Main driver - main files that drive components and control the life cycle of top level objects 
tests
src 
github.com

use makefile to stage vendor code before checking into parent repository
takes a snap shot of the dependancies and cehcks them into a vendor folder

## framework refactor code into a framework package
auth, logger, config helper classes 

## Testing 
code business logic is iolated from depenadncies such as external services. 
go mock
test packages independently
structuring tests 
-table driven test (DRY)
-sub tests can be run in parallel

good packaging is necessary because it its eanables code changes to happen 
faster with less risk, also leads to fewer bugs in production


# Scaling Algorithms

- Different Algorithms have different characteristics in terms of time and space complexity
- Time complexity defines the amount of time taken by an algorithm to run as a function of the inout size.
- Space complexity gives the amount of space taken for an alogrithm to run for a specific length of the input.

```

## Bubble sort

func bubbleSort(array []int) {
  swapped:= true;

  for swapped {
    swapped = false
    for i:= 0; i < len(array) - 1; i++ {
      if array[i + 1] < array[i] {
        array[i + 1], array[i ] = array[i], array[i + 1]
        swapped = true
      }
    }
  }
}


```

we will need n * n comparisons; that is, n2 operations. 
Assuming each operation takes a unit time, 
this algorithm is thus said to have a worst case complexity of O(n2), or quadratic complexity.

### Quicksort 
is another example of solving the problem. It is a type of the divide-and-conquer strategy of algorithm design, and is based on the idea of choosing one element of the array as a pivot and partitioning the array around this so that the elements to the left of the pivot are less than the value, while those on the right are larger than the pivot. A Go implementation is shown here:

```
func quickSort(array []int) []int {
  if len(array) <= 1 {
    return array
  }
  left, right:= 0, len(array) - 1

  // Pick a pivot randomly and move it to the end
  pivot:= rand.Int() % len(array)
  a[pivot], a[right] = a[right], a[pivot]

  // Partition
  for i:= range array {
    if array[i] < array[right] {
      array[i], array[left] = array[left], array[i]
      left++
    }
  }

  // Put the pivot in place
  array[left], array[right] = array[right], array[left]

  // Recurse
  quickSort(array[:left])
  quickSort(array[left + 1:])

  return array
}
```
nlogn. Thus, for quicksort, the time complexity is O(nlogn).
There is a linear scan of the data.
The input is divided into two parts and the code recourses on it.

# Distrubuted Algorithms
even when most efficient alogrithms cannot provide the answer in a reasonable time.
sometime dataset is so large it cant fit into the memory or disk of a single machine.
spolit tasks across machines involves alot of tricky low level details.

- Automatic Parallelization
- Communication and Coordination
- Distribution 
- Optimization for network and disk access
- Googles map reduce library

# Scaling data structures
- The way in which we store the data for an alogrithm has a huge impact of the scalability of the particular algorithm.
- Various datastructures provide different complexities for different operations.

# Profiling Data Structures 
The alogrithm scalability choices often manifest themselves in the choice of datastructures.
This table gives the time and space complexity of common data structures and their operations.

Array
LinkedList
Skip List
Hash Table
Binary Search Tree
Red black Tree

To summarize, based on what operation you want to do, how often, with a data structure, you can choose the right one for a job.

- consider how space allocation plays out with an increase in data.

Another aspect to consider in data structure is how space allocation plays out with an increase in data. For example, arrays are fixed length, and thus the data structure scalability is limited to the size at the time of allocation. In contrast, linked lists need no upfront total capacity allocation, but they do not offer the O(1) access time characteristics of arrays. That said, there are some hybrid data structures such as array lists that offer the best of both worlds.

Also, for a large amount of data, it is useful to think about how efficiently we can store the information. For example, if we have to store a Boolean flag for all users on a website, then there are two options:

Boolean array: one byte per flag/user
Bitset: a bit for each user/flag

The first option is slightly easier to code, but 50 M users will need 47 MB for the first option verses about 6 MB for the second. If there are multiple such flags for different use cases, you can imagine that bit sets will allow us to store more data in RAM, leading to better performance.

# Probabilistic data structures
sometimes scability problem is the sheer number of elements that the data structure has to handle.
top k frequent elements in a set where the number of values is very large 

Simple Solution: hash table to count frequencies and an min heap of elements ordered by the frequencies of elements 

Alternative solution: count min sketch consists of a two dimensional matrix  
Each element is mapped to one position at each of the d rows using d different/independent hash functions.

Whenever a new value is inserted, the corresponding counters are incremented and the total taken. We can then use this as an good estimate for frequency in the min-heap. If the frequency count is greater than the frequency of the element at the heap top, then the top of the heap is popped and the new element is inserted.

Reservoir sampling is an algorithm/data structure that enables these types of queries. First, we create a reservoir (array) of a size equal to the sample size required. In our case, say, 10. The first 10 elements are ingested as-in in the set.

# Scaling Data 

