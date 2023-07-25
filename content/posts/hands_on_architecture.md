---
title: "Hands On Arcitecture"
date: 2023-04-06T16:56:59+13:00
draft: false
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


