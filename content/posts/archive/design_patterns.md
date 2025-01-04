---
title: "Creational Design Patterns"
date: 2022-11-13T16:56:59+13:00
draft: true
---

- object creation mechanisms
- code using the object does not need to know about how the object is created.

## Factory Method
object used to created other onjects

```
func NewReservation(vertical, reservationDate string) Reservation {
    switch(vertical) {
        case "flight":
            return FlightReservationImpl{reservationDate,}
        case "hotel":
            return HotelReservationImpl{reservationDate,}
        default:
            return nil
    }   
}
```

## Builder
not striaght forward
- business rules to validate some paramters
- caching
- ideompotency and thread safety
- multiple constructors arguments

create different flavours of an object 

## Abstract Factory
Objects that need to be created together
two entities eg flight invoice to a hotel resrvation 
factories of factories construct.

## Singleton
Restricts creation of objects to a single one.
eg single coordinator object across multiple pieces of code.
intoduces global state so considered an antipattern

# Structural Design Patterns
Delineate relationships between objects and simplifiy design
receipes for various situations

## Adaptor
you have a new requirement and the components almosts meets that requirement.
Adapter class proxies what is required and delgates the work to the incompatible class.

```
Object Adaptor: Here, the adaptor class has an instance of the Adaptee class and the adaptor method delegates the work to the wrapped instance.
Class Adaptor: Here, the adaptor is a mix-in (a class with multiple inheritance) and inherits from both places:
The interface that is expected
The Adaptee interface
```

## Bridge

is a relationships between intefaces get mixed up with the implementations details
the abstraction and the implmentation cannot be extended

decouple the abstarction form the implmentation so both can vary independently
Segregating the the interface heirarchy and implementation hierrachy into two differnet trees.

The abstraction here is a struct not an interface, since in Go you can't have abstract structs/interfaces where you can store a reference to the seller implementation.

## Composite 

