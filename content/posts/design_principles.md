---
title: "Design Principles"
date: 2022-11-13T16:56:59+13:00
draft: false
---
Low Level Design 

Responsibility Assignment

Dependancy Management 

# Solid 

## Single Responsibility Principle 
One class should have one and only one resposibility

eg build a hierarchy of base and derived classes 
eg reservation interface
client code does not care about the type of reservation extraneous coupling is not introduced

```
type Reservation interface {
    GetReservationDate() string
    CalculateCancellationFee() float64  ***** 
    Cancel()
    GetCustomerDetails() []Customer
    GetSellerDetails() Seller
}
```

##  Open/Closed Principle 
You should be able to extend a classes behavior without modifying it

It should be possible to extend or override class behavior without having to modify code.

model trips with a repository of reservations 
calculateCancellationFee 
for each type of reservation. 

```
type Trip struct {
    reservations []Reservation
}

func (t *Trip) CalculateCancellationFee() float64 { //implements reservation interface
    total:= 0.0

    for _, r:= range(t.reservations) {
        total += r.CalculateCancellationFee()
    }

    return total
}

func (t *Trip) AddReservation (r Reservation) {
    t.reservations = append(t.reservations, r)
}
```

## Liskov Substitution Principle
Derived types must be substitutable for their base types.

derived classes must be usablethrough the base class interface without 
the need for the client to know the specific derived class.
s
The interface should be able to suffice for all structs that implement that interface.

```
The client who wants to compute the cancellation fee for a trip does not care about the specific type of reservations in the trip.
 The trip's code does not know how each reservation computes the cancellation fee.
```


## Integration Segregation Principle 
Many client specific interfaces are better than one general purpose interface
avoid derived classes implementing unrelated methods.

base reservation interface then have specific types of reservations with 
specialised methods.

## Dependancy Inversion Principle 
Depend on abstractions and not concretions

```
Every package should have interfaces that advertise functionality without the
implementation specifics.

when a package needs a dependancy it should take it as a parameters

Higher level modules should depend on intefaces and not concrete implementations.


```

