# 11: Learning to Write in Recursive

Through deliberate practice and taking note of various recursive patterns, I
discovered some techniques that helped me to learn to “write in recursive”
more easily.

Over the course of tackling various recursive problems, I began to find that
there are various “categories” of problems. Once I learned an effective tech-
nique for a certain category, when I found another problem that belonged to
the same category, I was able to apply the same technique to solve it.

## Recursive Category: Repeatedly Execute

For problems of this category, the last line of code in the function was a simple, single call to the function again. This line does one thing: it makes the next recursive call.

For instance:

```js
function countdown(number) {
  console.log(number);
  if(number === 0) { // number being 0 is the base case
    return;
  } else {
    countdown(number - 1);
  }
}
```

```ruby
def find_directories(directory)
  Dir.foreach(directory) do |filename|
    if File.directory?("#{directory}/#{filename}") && filename != "." && filename != ".."
      puts "#{directory}/#{filename}"

      # Recursively call this function on the subdirectory:
      find_directories("#{directory}/#{filename}")
    end
  end
end
```

### Recursive Trick: Passing Extra Parameters

```python
def double_array(array, index=0):
  # Base case: when the index goes past the end of the array
  if (index >= len(array)):
    return
  array[index] *= 2
  double_array(array, index + 1)
```

Once we have the index as a function argument, we now have a way of incrementing and tracking the index as we make each successive recursive call. In each successive call, we pass in the array again as the first argument, but we also pass along an incremented index. This allows us to keep track of an index just as we would in a classical loop.

All we updated here was setting a default argument of index=0. This way, the first time we call the function, we don’t have to pass in the index parameter. However, we still get to use the index parameter for all successive calls.

The “trick” of using extra function parameters is a common technique in writing recursive functions, and a handy one.

## Recursive Category: Calculations

In this second general category, the problems perform a calculation based on a subproblem. 

One area in which recursion shines is where we need to act on _a problem that has an arbitrary number of levels of depth_. A second area in which recursion shines is _where it is able to make a calculation based on a subproblem of the problem at hand_.

A subproblem is a version of the very same problem applied to a smaller input.

For example, the factorial of 6 is the result of 6 * the factorial of 5; the factorial of 5 is the result of 5 * the factorial of 4, and so on. This is a recursive structure.

```ruby
def factorial(number)
  if number == 1
    return 1
  else
    return number * factorial(number - 1)
  end
end
```

### Two Approaches to Calculations

We’ve seen that when writing a function that makes a calculation, there are two potential approaches: we can try to build the solution from the “bottom up,” or we can attack the problem going “top down” by making the calculation based on the problem’s subproblem. The truth is that both approaches can be achieved through recursion.

#### The Bottom Up Approach

While we previously saw the bottom-up approach using a classic loop, we can also use recursion to implement the bottom-up strategy.

```ruby
def factorial(n, i=1, product=1)
  return product if i > n
  return factorial(n, i + 1, product * i)
end
```

While we can use recursion in this way to achieve the bottom-up approach, it’s not particularly elegant and does not add much value over using a classic loop.

When going bottom up, we’re employing the same strategy for making the calculation whether we’re using a loop or recursion. The computational approach is the same.

But to go top down, we need recursion. And because recursion is the only way to achieve a top-down strategy, it’s one of the key factors that makes recursion a powerful tool.

#### The Top-Down Approach in Recursion: A New Way of Thinking

This brings us to the central point of this chapter: recursion shines when implementing a top-down approach because going top down offers a new mental strategy for tackling a problem. That is, a recursive top-down approach allows one to think about a problem in a completely different way.

Specifically, when we go top down, we get to mentally “kick the problem down the road.” We can free our mind from some of the nitty-gritty details we normally have to think about when going bottom up.

##### The Top-Down Thought Process

I found that when tackling a top-down problem, it helps to think the following three thoughts:

1. Imagine the function you’re writing has already been implemented by someone else.
2. Identify the subproblem of the problem.
3. See what happens when you call the function on the subproblem and go from there.

##### Example: Array Sum

Write a function called sum that sums up all the numbers in a given array.

The first thing we’ll do is imagine that the sum function has already been
implemented.

Next, let’s identify the subproblem. In our case, we can say that the subproblem is the array, `[2, 3, 4, 5]` —that is, all the numbers from the array save the first one.

Finally, let’s see what happens when we apply the sum function to our sub problem. If the sum function “already works,” and the subproblem is `[2, 3, 4, 5]`, what happens when we call sum(`[2, 3, 4, 5]`)? Well, we get the sum of 2 + 3 + 4 + 5, which is 14. To solve our problem of finding the sum of `[1, 2, 3, 4, 5]` then, we can just add the first number, 1, to the result of sum(`[2, 3, 4, 5]`).

The last thing we need to do is handle the base case. That is, if each subproblem recursively calls its own subproblem, we’ll eventually reach the subproblem of sum(`[5]`). This function will eventually try to add the 5 to the remainder of the array, but there are no other elements in the array. To deal with this, we can add the base case.

```ruby
def sum(array)
  # Base case: only one element in the array:
  return array[0] if array.length == 1
  return array[0] + sum(array[1, array.length - 1])
end
```

##### Example: String Reversal

```ruby
def reverse(string)
  # Base case: string with just one character
  return string[0] if string.length == 1
  return reverse(string[1, string.length - 1]) + string[0]
end
```