# 110. Balanced Binary Tree

## Description

Given a binary tree, determine if it is **height-balanced**.

A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Examples

**Example 1:**
```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

**Example 2:**
```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Example 3:**
```
Input: root = []
Output: true
```

## Constraints

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-10^4 <= Node.val <= 10^4`

## Solution

Use **bottom-up recursion** that returns early when imbalance is detected:

1. Recursively compute the height of left and right subtrees.
2. If either subtree is unbalanced (returns -1), propagate -1 up.
3. If the height difference exceeds 1, return -1.
4. Otherwise return `max(left, right) + 1`.

The tree is balanced if the final result is not -1.

**Time complexity:** O(n)
**Space complexity:** O(h) where h is the tree height (recursion stack)
