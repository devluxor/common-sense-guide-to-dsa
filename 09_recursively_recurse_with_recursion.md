# Chapter 09: Recursively Recurse with Recursion

Recursion is official name for when a function calls itself. While infinite function calls are generally useless - and even dangerous - recursion is a powerful tool that can be harnessed. And when we harness the power of recursion, we can solve particularly tricky problems.

## Recurse instead of Loop

We can use recursion instead of looping through an array;

```js
function countdown(number) {
  console.log(number);
  countdown(number - 1)
}
```

In almost any case in which you can use a loop, you can also use recursion. Now, just because you can use recursion does not mean that you should use it. Recursion is a tool that allows for writing elegant code. In that example and in other simple iterations, the recursive approach is not better than a classic `for` loop. However, there are cases in which recursion really shines.

## The Base Case

We need a clause that will stop the recursion to go infinitely: the case in which the function will not recurse is known as the _base case_:

```js
function countdown(number) {
  console.log(number);
  if(number === 0) { // Base Case
    return;
  } else {
    countdown(number - 1);
  }
}
```

## Reading Recursive Code

It takes time and practice to get used to recursion, and you will ultimately learn two sets of skills: Reading recursive code, and writing recursive code. Reading recursive code is somewhat easier.

This is the process you need to know for reading recursive code:

1. Identify what the base case is.
2. Walk through the function assuming it’s dealing with the base case.
3. Then, walk through the function assuming it’s dealing with the case immediately before the base case.
4. Progress through your analysis by moving up the cases one at a time.

This is a recursive function that returns a number's factorial:

```ruby
def factorial(number)
  if number == 1 # Base case
    return 1 
  else
    return number * factorial(number - 1)
  end
end
```

1!                 = 1 <-- Base case
2! = 2 * 1         = 2
3! = 3 * 2 * 1     = 6
4! = 4 * 3 * 2 * 1 = 24

Becomes

factorial(4) = 4 * factorial(3)  ▲
factorial(3) = 3 * factorial(2)  |
factorial(2) = 2 * factorial(1)  |
factorial(1) = 1                 | <-- Base case

## Recursion in the Eyes of the Computer

The computer calls `factorial(4)`, and, before the method is complete, it calls `factorial(3)`; then, before `factorial(3)` is complete, it calls `factorial(2)`, and then, before it completes `factorial(2)`, it calls `factorial(1)`. Technically, while the computer runs `factorial(1)`, it’s still in the middle of `factorial(2)`, which in turn is running within `factorial(3)`.

The computer uses a stack to keep track of which functions it's in the middle of calling. This stack is called the _call stack_.

The computer begins by calling factorial(4). Before the method completes executing, however, factorial(3) gets called. In order to track that the computer is still in the middle of factorial(4), the computer pushes that info onto a call stack, and does the same until it reaches the base case, resulting in a call stack that looks like this:
```
[factorial(1)] <--- Base case
[factorial(2)]
[factorial(3)]
[factorial(4)]
```

When the computer reaches this case, it can resolve it, returning the value of `1`, to then popping the last function in the call stack:

```
                  ------> [factorial(1)]
[factorial(2)]
[factorial(3)]
[factorial(4)]
```

Even after the computer completes `factorial(1)`, it knows that it’s not finished with everything it needs to do since there’s data in the call stack, which indicates that it’s still in the middle of running other methods that it needs to complete. So, once factorial(1) has been popped out and it has returned the `1`, the next function in the stack can be resolved as well, returning the factorial of 2 (2 * 1; the 1 that got from the 1 returning from `factorial(2 - 1)`). Then, the function is popped as the last one:

```
                  ------> [factorial(2)]
[factorial(3)]
[factorial(4)]
```

And this process continues:

```
                  ------> [factorial(3)]
[factorial(4)]
```

```
                  ------> [factorial(4)]
```

At this point the stack is empty, so the computer knows it's done executing all of the methods, and the recursion is complete.

Looking back at the example, we can see that the order in which the computer calculates the factorial of 4 is as follows:

1. `factorial(4)` is called first.
2. `factorial(3)` is called second.
3. `factorial(2)` is called third.
4. `factorial(1)` is called fourth.
5. `factorial(1)` is _resolved_ first.
6. `factorial(2)` is resolved based on the result from `factorial(1)`
7. `factorial(3)` is resolved based on the result from `factorial(2)`
8. `factorial(4)` is resolved based on the result from `factorial(3)`

Interestingly, in the case of infinite recursion (such as the very first example in our chapter), the program keeps on pushing the same method over and over again onto the call stack, until there’s no more room in the computer’s memory - and this causes an error known as _stack overflow_.

## Recursion in Action

Recursion is a natural fit in any situation where you find yourself having to repeat an algorithm within the same algorithm. In these cases, recursion can make for more readable code.

Imagine the example of having to traverse through a filesystem; you don't want the script to only deal with the files inside the one directory, but act on all the files within the _subdirectories_ of the directory, and the subdirectories of the subdirectories, etc.

This is a recursive method that can go arbitrarily deep into a directory to get all files' names:

```ruby
def find_directories(directory)
  Dir.foreach(directory) do |filename|
    if File.directory?("#{directory}/#{filename}") && filename != "." && filename != ".."
      puts "#{directory}/#{filename}"
      find_directories("#{directory}/#{filename}")
    end
  end
end
```

This diagram illustrates this algorithm, and specifies the order in which the script traverses the subdirectories:

![title](images/26.png)

Recursion is often a great choice for an algorithm in which the algorithm itself does not know on the outset how many levels deep into something it needs to dig.

