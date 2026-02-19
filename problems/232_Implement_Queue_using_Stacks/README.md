# 232. Implement Queue using Stacks

## Description

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

Implement the `MyQueue` class:

- `void push(int x)` — Pushes element x to the back of the queue.
- `int pop()` — Removes the element from the front of the queue and returns it.
- `int peek()` — Returns the element at the front of the queue.
- `boolean empty()` — Returns `true` if the queue is empty, `false` otherwise.

## Examples

**Example 1:**
```
Input:
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output:
[null, null, null, 1, 1, false]

Explanation:
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2]
myQueue.peek();  // return 1
myQueue.pop();   // return 1, queue is [2]
myQueue.empty(); // return false
```

## Constraints

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `peek`, and `empty`.
- All the calls to `pop` and `peek` are valid.

## Solution

Use **two stacks** — an input stack and an output stack:

1. `push`: always push onto the input stack.
2. `pop`/`peek`: if the output stack is empty, transfer all elements from the input stack (reversing LIFO → FIFO order). Then pop/peek from the output stack.
3. `empty`: both stacks must be empty.

**Time complexity:** O(1) amortized per operation
**Space complexity:** O(n)
