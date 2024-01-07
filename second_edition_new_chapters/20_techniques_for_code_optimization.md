# 20: Techniques for Code Optimization

## Prerequisites for code optimization 

### General Algorithm Before Actual Optimization

1. Determine the Big O category of your current algorithm.
2. Determine the best _possible_ Big O for the problem at hand. (For example, O(log N) when searching an array) This set our goal.
3. If the best-imaginable Big O is faster than your current Big O, you can now try to optimize your code, with the goal of bringing your algorithm as close to the best-imaginable Big O as possible.

It’s important to stress that **it’s not always possible to achieve the best-imaginable Big O**. (For instance, it's impossible to beat linear time when searching an unordered array)

It is a success to improve speeds by one degree of magnitude or one category, even if we don't reach the best possible Big O we set as a goal. For example, for an initial speed of O(N<sup>2</sup>), it is a success to reach O(N).

## Techniques and strategies for Code Optimization

### Lookup tables: The Power of Hashes

#### Example: Processing authors and books from two different data structures (from O(N<sup>2</sup>) or O(N * M) to O(N))

#### Example: Two-sum Problem

### Pattern Detection

Study the problem
Generate a table with results using different outputs (generally in progression, like 1, 2, 3...)
Detect the pattern
Use the pattern to implement a faster solution

#### Example: Coin Game

#### Example: Sum Swap Problem (pattern detection + lookup tables)

### Greedy Algorithms
