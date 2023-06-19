# Chapter 08

Stacks and queues are not entirely new: they are simply arrays with restrictions.

Stacks and queues are elegant tools for handling temporary data; from operating system architecture to printing jobs, they are ideal as temporary containers.

Stacks and queues allow you to handle data in order, and then get rid of it once you don't need it anymore.

## Stacks

Stacks store data as arrays do - it's simply a list of elements. They have three constraints, however:

1. Data can only be inserted at the end of a stack.
2. Data can only be read from the end of a stack.
3. Data can only be removed from the end of a stack

Most computer science literature refers to the end of the stack as its top, and the beginning of the stack as its bottom. Inserting a new value into a stack is also called _pushing_ onto the stack; removing elements from the top of the stack is called _popping_ from the stack.

A common acronym used to describe stack operations is LIFO: Last In, First Out.

```ruby
stack = []

stack.push(1)
stack.push(2)
stack.push(3)

stack # => [1, 2, 3]

stack.pop

stack # => [1, 2]
```

## Stacks in Action

Although a stack is not typically used to store data on a long term basis, it can be a great tool to handle temporary data as part of various algorithms.

Here is an algorithm to check if a string has the correct number of matching parentheses and other bracing symbols:

We’d prepare an empty stack, and then we read each character from left to right, following these rules:

1. If we find any character that isn’t a type of brace (parenthesis, square bracket, or curly brace), we ignore it and move on.

2. If we find an opening brace, we push it onto the stack. Having it on the stack means we’re waiting to close that particular brace.

3. If we find a closing brace, we inspect the top element in the stack. We then analyze:
    - If there are no elements in the stack, that means we’ve found a closing brace without a corresponding opening brace beforehand. This is Syntax Error Type #2.
    - If there is data in the stack, but the closing brace is not a corresponding match for the top element of the stack, that means we’ve encountered Syntax Error Type #3.
    - If the closing character is a corresponding match for the element at the top of the stack, that means we’ve successfully closed that opening brace. We pop the top element from the stack, since we no longer need to keep track of it.
    
4. If we make it to the end of the line and there’s still something left on the stack, that means there’s an opening brace without a corresponding closing brace, which is Syntax Error Type #1.

Stacks are ideal for processing any data that should be handled in reverse order to how it was received (LIFO). For example, the 'undo' function in a word processor or function calls in a networked application.

## Queues

Like stacks, queues are arrays with three restrictions:

1. Data can only be inserted at the end of the queue.
2. Data can only be read from the front of the queue.
3. Data can obly be removed from the front of the queue.

## Queues in Action

Queues are very comon in many applications, ranging from printing jobs to background workers in web applications.

Queues are also the perfect tool for handling asyncronous requests - they ensure that the requests are processed in the same order they were received. They are also commonly used to model real-world scenarios where events need to occur in a certain order, such as airplanes waiting for takeoff and patients waiting for their doctors.

