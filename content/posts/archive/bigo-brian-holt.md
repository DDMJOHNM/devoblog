---
title: "Big O - Brian Holt Course"
date: 2024-12-13T16:56:59+13:00
draft: false
---

BO important details of an algorithm

3x2+x+1 o(n2)
look for the biggest term
carries most of the weight of the algorithm

Complexity O(n) - goes through all the items in the loops 
eg for loop v size of input array
n is length of the array

no loops constant 

big o ignores coeficient just looking for loops

```
function find(needle, haystack){
    for (var i=0; i < haystack.length; i++) {
        if (haystack === needle) return true;
    }   
}

```
still o(n)
worst case scenario - last item in array
best case scenario - first item in the array
bigO doesnt care 


pair of values 
nested loops
o(n) --> Linear 
o(n2)  ----> n squared (Quadratic Time)
o(n3)  ----> n cubed -NO 
```
makeTuples(input){
    var answer = [];
      for (var i=0; i < input.length; i++) {
        for (var j=0; j < input.length; j++) {
            answer.push(input[i], input[j]) 
        } 
      } 
      return answer;
}

```
Algorithms - designed different problems 

Trade offs better v worse complexity  

```
function getMiddleOfArray(array){
    return array[Math.floor(array.length/2)]
}
```
Constant Time 
o(1) - doesnt matter how long the array is still takes the same amount of time to do 

graph 
(y) how long it takes the fucntion to run 
(x) number of items in the array

Logrithmic time (1000 v 100 items not that much more time)
algorithm to divide into chunks

more inputs will it take longer not just for loops 

Question --> it depends i need more information 
ie how big is the system 

make more maintainable code v making faster

science of trade offs

- one tool is big O

Complete picture of what the requirements are 

(helps to choose correct algoithm)
end user 
constraints 
devices - where what 
how big is the data set
how important to the user
what are the user doing

time (computational complexity) - how long it takes to run  
space complexity (space)

==================================================================

## Trade offs  
(problem constraits)
discuss trade offs (money v customer experience)
based on constraints 

code smells - might be a bad idea 
could this be done better (what is better in the particular case)

multiple choices - no rules 
some times upu cant ignore the incidentials 
eg for loop with 10000 line of code in it 

sometime bigo bad to use 
job happens once a day took 30 mins - dont optimise, doesnt matter 

readibility and maintainability most important things in code
(code is communication) future self future devs
only go for n code if constraints actually called for

simple code less bugs

human time is more valuable teh computer time 

don't prematurely optimise

99% time features are built into languages or existing module  

Load Testing - relative to customer base 

## Bubble sort 

one number is biogger than another in the array swap 
go through and nothing swaps done
end of array then start again
- larger item bubble to top 
- look at less and less of the array 

av case on2
every item will see every item in the array
worst case 
reverse sorted list 
best case n
sorted list 

- spacial complexity constant

outer while loop 
inner for loop

- destructive operates on array itself

### Example 
```
function bubbleSort(nums) {
  // code goes here
  // let i = 0;
  // do {
  //   i = i + 1
  //   for (j = 0; j < nums.length; j++) {
  //     if (nums[i] < nums[j]) {
  //         nums[i] = nums[j]
  //     } else {
  //       break;
  //     }
  //   } 

  // } while (i <  nums.length - 1);

  let swapped = false;
  
  do {
    swapped = false;
    for ( i=0; i < nums.length; i++) {
      if(nums[i] > nums[i+1]){
        const temp = nums[i];
        nums[i] = nums[i+1];
        nums[i+1] = temp;
        swapped = true; 
      }
    }
  
  } while (swapped)  
   return nums; 
}

// unit tests
// do not modify the below code
test.skip("bubble sort", function () {
  const nums = [10, 5, 3, 8, 2, 6, 4, 7, 9, 1];
  const sortedNums = bubbleSort(nums);
  expect(sortedNums).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
});


```

## Insertion Sort

array of length one always sorted
index one insert in a sorted way

[3,2,5,4,1] 
is larger than 3 no move to the right
Begining of list insert 2 at index 0
is index 4 larger than 5 no move 5 to the right 
is index 4 larger than 3 yes
insert after 3 and index 2
[2,3,4,5,1]

best case nothing has to shift
average case mostly sorted
worse case sorted reverse 

on2 with better coeficients

**  most or totally sorted lists

spatial complexity constant time

Stable

Example

```

function insertionSort(nums) {
  // code goes here
  //first item always sorted
  for (let i = 0; i < nums.length; i++) {//forwards    
    let numberToInsert = nums[i]
    let j;
    //everything before i is the sorted elements  of the array  
    //greater than the number to insert and greater than 0
    for (j=i-1; nums[j] > numberToInsert && j >=0; j--){
      //if number in the right place assigns to itself 
      nums[j+1] = nums[j] //moves forward 1
    } 
    //j works backwards through sorted part of array to insert the next thing
    nums[j+1] = numberToInsert; //inserts number into exact space
  } 
  return nums;

}

// unit tests
// do not modify the below code
test("insertion sort", function () {
  const nums = [10, 5, 3, 8, 2, 6, 4, 7, 9, 1];
  insertionSort(nums);
  expect(nums).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
});


```


