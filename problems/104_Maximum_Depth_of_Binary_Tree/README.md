# 104. Maximum Depth of Binary Tree

## Problem

Given the `root` of a binary tree, return its **maximum depth**.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Solutions

### Go

#### DFS (`maxDepth`)

Recursively compute the depth of left and right subtrees. Return the larger of the two plus one. Base case: nil node returns 0.

- Time: O(n)
- Space: O(h) where h is the height of the tree

### Python

#### DFS (`maxDepth`)

Same recursive approach — return `1 + max(left_depth, right_depth)` at each node.

- Time: O(n)
- Space: O(h) where h is the height of the tree
