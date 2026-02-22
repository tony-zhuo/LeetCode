# 3849. Maximum Bitwise XOR After Rearrangement

## Problem

Given two binary strings `s` and `t` of equal length, you can rearrange the characters of `t` (but not `s`). Return the binary string representing the maximum possible `s XOR t`.

## Solutions

### Go

#### Greedy (`maximumXor`)

Count the ones and zeros in `t`. Process `s` from left (most significant bit) to right. For each bit, greedily assign the opposite bit from `t` to maximize the XOR result. If no opposite bits remain, use the same bit (producing 0).

- Time: O(n)
- Space: O(n)

### Python

#### Greedy (`maximumXor`)

Same greedy approach — count ones/zeros in `t`, then iterate through `s` assigning opposite bits first to maximize higher-order positions.

- Time: O(n)
- Space: O(n)
