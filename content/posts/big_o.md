---
title: "Big O Notation"
date: 2024-12-13T16:56:59+13:00
draft: false
---

Big O is a way to categorize your algorithms time or memory requirements based on input.  

Generalize growth of an algorithm.

O(n) - algorithm with grow linerily based on input.

Make decisions about what datastructures and algorithms to use.

Knowing how they perform can help greatly mking the best possible program.

As your input grows how fast does your computation or memory grow.

Growth is with respect to input.

1. Simplest trick for complexity look for loops. 

```
function sum_char_codes(n: string): number{
    let sum = 0;
    for (let i =0; i < n.length; i++){
        sum += n.charCodeAt(i);
    }
    return sum;
} 

for loop has to execute length of string 
for one more unit of string is one more loop.
- linear complexity 

```

```
function sum_char_codes(n: string): number{
    let sum = 0;
    for (let i =0; i < n.length; i++){
        sum += n.charCodeAt(i);
    }

    for (let i =0; i < n.length; i++){
        sum += n.charCodeAt(i);
    }
   return sum;
} 

2. Drop constants 
O(2N) -> O(N)

```

```
function sum_char_codes(n: string): number{
    let sum = 0;
    for (let i =0; i < n.length; i++){
       const charCode = n.charCodeAt(i);
       //E
       if(charCode === 69) {
        return sum;
       }
       sum += charCode;
    }
   return sum;
} 

3. worse case 
still O(n)

```

o(1) - constant time same set of operations instant

o(logn) - logn  binary search trees - half the amount of input you have to search but only look at one point at a time.

o(n) - linear

o(nlogn) - nlogn Quick Sort (go over everything but search half)

o(n^2) - n squared 

```
function sum_char_codes(n: string): number{
    let sum = 0;
    for (let i =0; i < n.length; i++){
         const charCode = n.charCodeAt(i);
         for (let j =0; j < n.length; j++){
            sum += charCode;
         }
    }
    return sum;
} 


```
o(n^3) - n cubed 

```
function sum_char_codes(n: string): number{
    let sum = 0;
    for (let i =0; i < n.length; i++){
         const charCode = n.charCodeAt(i);
         for (let j =0; j < n.length; j++){
             for (let k =0; k < n.length; k++){
                sum += charCode;
             }
         }
    }
    return sum;
} 

```

o(sqrt(n)) 

o(2^n) 
o(n!)

Big O upper bound of runtime there are others.


## Array Data Structure
const a = [];
Contiguous (unbreaking) memory space (contains a certain number of bytes) 

```
//NodeJs

const a = new ArrayBuffer(6); // bytelength = 6 
const a8 = new Unit8Array(a);
- unsigned 8 bit numbers (0-255)
a8[0] = 45; //set first byte (decimal)
a8[2] = 45; //set 2nd byte (decimal)
const a16 = new Unit16Array(a);
a16[2] = 0x4545 (heximdecimal) 

<2d 00 2d 45 45> byte length 6 (walk through array 8bit units, 16bit units)

```

array insert overwrites -> cannot not grow array without reallocations
delete at index -> 
insert at index -> (a + width*offset)

static array has no methods like push or pop





//up to trees