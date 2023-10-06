# Chapter 14: Dealing with Space Constraints

When analyzing the efficiency of various algorithms, we’ve focused exclusively on their _time complexity_ - that is, how fast they run. There are situations, however, where we need to measure algorithm efficiency by another measure known as_ space complexity_, which is how much memory an algorithm consumes.

Space complexity becomes an important factor when memory is limited. This can happen when programming for small hardware devices that only contain a relatively small amount of memory or when dealing with very large amounts of data, which can quickly fill even large memory containers.

In a perfect world, we’d always use algorithms that are both quick and consume a small amount of memory. However, there are times when we’re faced with choosing between the speedy algorithm or the memory

## Big O Notation As Applied to Space Complexity

Computer scientists use Big O Notation to describe space complexity just as they do to describe time complexity. 

We’ve used Big O Notation to describe an algorithm’s speed as follows: For N elements of data, an algorithm takes a relative number of operational steps. For example, O(N) means that for N data elements, the algorithm takes N steps. And O(N<sup>2</sup>) means that for N data elements, the algorithm takes N<sup>2</sup> steps.

Big O can similarly be used to describe how much space an algorithm takes up: _For N elements of data, an algorithm consumes a relative number of additional data elements in memory_. 

Let’s say that we were writing a JavaScript function that would accept an array of strings, and return an array of those strings in ALL CAPS:

```js
function makeUpperCase(array) {
  let newArray = [];
  for(let i = 0; i < array.length; i++) {
    newArray[i] = array[i].toUpperCase();
  }
  return newArray;
}
```

By the time this function is complete, we have two arrays floating around in our computer’s memory. We have `array`, which contains `["amy", "bob", "cindy","derek"]`, and we have `newArray`, which contains `["AMY", "BOB", "CINDY", "DEREK"]`.

When we analyze this function in terms of space complexity, we can see that it accepts an array with N elements, and creates a brand new array that also contains N elements. Because of this, we’d say that the `makeUpperCase()` function has a space efficiency of `O(N)`.

![title](images/71.png)

This graph is identical to the way we’ve depicted O(N) in the graphs from previous chapters, with the exception that the vertical axis now represents memory rather than speed.

Let’s present an alternative `makeUpperCase() `function that is more memory efficient (assuming we can modify the original input array):

```js
function makeUpperCase(array) {
  for(var i = 0; i < array.length; i++) {
    array[i] = array[i].toUpperCase();
  }
    return array;
  }
```

In this second version, we do not create any new variables or new arrays. Instead, we modify each string within the original array in place. 

Since this function does not consume any memory in addition to the original array, we’d describe the space complexity of this function as being O(1). Remember that by time complexity, O(1) represents that the speed of an algorithm is constant no matter how large the data. Similarly, by space complexity, O(1) means that the memory consumed by an algorithm is constant no matter how large the data.

It’s important to reiterate that in this book, we judge the space complexity based on additional memory consumed - known as auxiliary space - meaning that we don’t count the original data.

Let’s now compare the two versions of `makeUpperCase()` in both time and space complexity:

| Version | Time Complexity | Space complexity |
| --- | --- | --- |
| Vesion 01 | O(N) | O(N) |
| Version 02 | O(N) | O(1) |

This is a pretty strong case that Version #2 is preferable to Version #1.

## Trade-offs between time and space

In Speeding Up Your Code with Big O, we wrote a JavaScript function that checked whether an array contained any duplicate values. Our first version looked like this:
```js
function hasDuplicateValue(array) {
  for(let i = 0; i < array.length; i++) {
    for(let j = 0; j < array.length; j++) {
      if(i !== j && array[i] == array[j]) {
        return true;
      }
    }
  }
  return false;
}
```

It uses nested for loops, and we pointed out that it has a time complexity of O(N<sup>2</sup>).

We then created a more efficient version, which looked like this:

```js
function hasDuplicateValue(array) {
  let existingNumbers = [];
  for(let i = 0; i < array.length; i++) {
    if(existingNumbers[array[i]] === undefined) {
      existingNumbers[array[i]] = 1;
    } else {
     return true;
    }
  }
  return false;
}
```

We declared victory with this second version, pointing to its time complexity of O(N) compared with the first version’s O(N2). Indeed, from the perspective of time alone, the second version is much faster.

However, when we take space into account, we find that this second version has a disadvantage compared with the first version. The first version does not consume any additional memory beyond the original array, and therefore has a space complexity of O(1). On the other hand, this second version creates a brand new array that is the same size as the original array, and therefore
has a space complexity of O(N).

Let’s look at the complete contrast between the two versions of `hasDuplicateValue()`:

| Version | Time Complexity | Space complexity |
| --- | --- | --- |
| Vesion 01 | O(N<sup>2</sup>) | O(1) |
| Version 02 | O(N) | O(N) |

How do we decide which to choose? The answer, of course, is that it depends on the situation.

## Final Toughts

The analysis of data structures and algorithms can dramatically affect your code - in speed, memory, and even elegance.
What you can take away from this book is a framework for making educated technology decisions. Now, computing contains many details, and while something like Big O Notation may suggest that one approach is better than another, other factors can also come in to play. The way memory is organized within your hardware and the way that your computer language of choice implements things under the hood can also affect how efficient your code may be. Sometimes what you think is the most efficient choice may not be due to various external reasons.

Because of this, it is always best to test your optimizations with benchmarking tools. There are many excellent software applications out there that can measure the speed and memory consumption of your code. The knowledge in this book will point you in the right direction, and the benchmarking tools will confirm whether you’ve made the right choices.