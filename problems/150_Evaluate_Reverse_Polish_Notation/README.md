# 150. Evaluate Reverse Polish Notation

## Problem

You are given an array of strings `tokens` that represents an arithmetic expression in [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:
- The valid operators are `+`, `-`, `*`, and `/`.
- Each operand may be an integer or another expression.
- The division between two integers always truncates toward zero.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in reverse polish notation.

## Solutions

### Go

#### Stack (`evalRPN`)

Iterate through tokens. Push numbers onto the stack. When encountering an operator, pop the top two elements, apply the operator, and push the result back. The final element on the stack is the answer.

- Time: O(n)
- Space: O(n)