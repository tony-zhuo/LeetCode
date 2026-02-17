# 226. Invert Binary Tree

## Description

Given the `root` of a binary tree, invert the tree, and return its root.

Inverting a binary tree means swapping every left child with its corresponding right child.

## Examples

**Example 1:**
```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

**Example 2:**
```
Input: root = [2,1,3]
Output: [2,3,1]
```

**Example 3:**
```
Input: root = []
Output: []
```

## Constraints

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution

Use **recursion** to swap children at each node:

1. If the node is `nil`, return.
2. Swap its left and right children.
3. Recursively invert the left and right subtrees.

**Time complexity:** O(n)
**Space complexity:** O(h) where h is the tree height (recursion stack)
