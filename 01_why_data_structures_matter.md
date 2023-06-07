# Chapter 01: Why Data Structures Matter

Data is a broad term that refers to all types of information, down to the most basic numbers and strings.

The organization of data does not just matter for organization's sake, but can significantly impact _how fast your code runs_. Depending on how you choose to organize your data, your program may run faster or slower by orders of magnitude.

## The Array: the Foundational Data Structure

The array is one of the most basic data structures in computer science; it's simply a list of data elements.

The _index_ of an array is the number which identifies where a piece of data lives inside the array. We begin counting the index at 0. (Arrays are zero-indexed)

To understand the performance of a data structure, like an array, we need to analyze the common ways that our code might interact with that data structure.

Most data structures are used in four basic ways (we interact with them in four basic ways), which we refer to as _operations_:

1. Read: looking something up from a particular spot within a data structure.

2. Search: looking for a particular value within a data structure.

3. Insert: adding another value to a data structure.

4. Delete: removing a value from a data structure.

_In the DSA (data structures and algorithms) context, when we measure how fast an operation takes, we do not refer to how fast the operation takes in terms of pure time, but instead in how many steps it takes_

Why?

Measuring the speed of an operation in terms of time is flaky, since it will always change depending on the hardware that it is run on. However, we can measure the speed of an operation in terms of how many steps it takes.

Measuring the speed of an operation is also known as measuring its _time complexity_. Throughout this book, we will use the terms _speed_, _time complexity_, _efficiency_ and _performance_ interchangeably. They all refer to the number of steps that a given operation takes.

### Reading

Reading from an array actually takes just one step. This is because the computer has the ability to jump to any particular index in the array and peek inside.

A computer's memory can be viewed as a giant collection of cells. 

When a program declares an array, it allocates a contiguous set of empty cells for use in the program.

Every cell in a computer's memory has a specific address, represented by a simple number. Each cell's memory address is one number greater than the previous cell.

When the computer reads a value at a particular index of an array, it can jump straight to that index in one step because of the combination of the following facts:

1. A computer can jump to any memory address in one step.

2. Recorded in each array is the memory address which it begins at.

3. Every array begins at index 0.

Reading from an array is a very efficient operation, since it takes just one step. An operation with just one step is naturally the fastest type of operation.

### Searching 

To search for a value within an array, the computer starts at index 0, checks the value, and if it does not find what it's looking for, moves on to the next index. It does this until it finds the value it's seeking.

This basic search operation in which the computer checks each cell one at a time is known as _linear search_.

For N cells in an array, linear search will take a maximum of N steps. In this context N is just a variable that can be replaced by any number.

Searching is less efficient than reading, since searching can take many steps, while reading always takes just one step no matter how large the array.

### Insertion

The efficiency of inserting a new piece of data inside an array depends on where inside the array you'd like to insert it. 

Inserting an element at the end of the array just takes one step; Inserting a new piece of data at the beginning or the _middle_ of the array is a different story. In these cases, we need to _shift_ many pieces of data to make room for what we're inserting, leading to additional steps.

The worst case scenario for insertion into an array -that is, the scenario in which insertion takes the most steps- is where we insert data at the beginning of the array. This is because when inserting into the beginning of the array we have to move all the other values one cell to the right.

We say that insertion in the worst case scenario can take up to N + 1 steps for an array containing N elements. (N shifts, 1 insertion)

### Deletion

The actual deletion itself is really just one step, but we need to follow up with additional steps of shifting data to the left to close the gap caused by the deletion.

Like insertion, the worst case scenario of deleting an element is deleting the very first element of the array.


## Sets: How a Single Rule Can Affect Efficiency

A set is a data structure that does not allow duplicate values to be contained within it. There are many types of sets, but for this discussion we'll talk about an _array-based set_. The only difference between this set and a classic array is that the set never allows duplicate values to be inserted into it. 

Sets are useful when you need to ensure that you don't have duplicate data.

In any case, a set is an array with one simple constraint of not allowing duplicates. Yet, this constraint causes the set to have a different efficiency for one of the four primary operations.

Reading from a set is exactly the same as reading from an array.

Searching a set is also the same as searching an array.

Insertion, however, is where arrays and sets diverge. With a set, the computer first needs to determine _that this value does not already exist in this set_. So, every insertion requires a search.

In a worst case scenario, where we are inserting a value at the beginning of a set, the computer needs to search N cells to ensure that the set does not already contain that value; that gives a total of 2N + 1 steps.

Sets are important when you need to ensure that there is no duplicate data. But when you don't have such a need, an array may preferable.

## Wrapping Up

Analyzing the number of steps that an operation takes is the heart of understanding the performance of data structures. Choosing the right data structure for your program can spell the difference between bearing a heavy load versus collapsing under it. 

