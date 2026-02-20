# 67. Add Binary

## Problem

Given two binary strings `a` and `b`, return their sum as a binary string.

## Solutions

### Go

#### Simulation (`addBinary`)

Traverse both strings from right to left, summing digits with a carry. Build the result string by prepending each computed bit.

- Time: O(max(m, n))
- Space: O(max(m, n))

### Python

#### Simulation (`addBinary`)

Same approach — iterate from the end of both strings, accumulate carry, and prepend each bit to the result.

- Time: O(max(m, n))
- Space: O(max(m, n))
