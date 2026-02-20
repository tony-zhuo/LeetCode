# 543. Diameter of Binary Tree

## Problem

Given the `root` of a binary tree, return the length of the **diameter** of the tree.

The diameter of a binary tree is the **longest path** between any two nodes in a tree. This path may or may not pass through the `root`. The length of a path between two nodes is represented by the number of edges between them.

## Solutions

### Go

#### DFS (`diameterOfBinaryTree`)

Recursively compute the depth of each subtree. At every node, the diameter passing through that node is `left_depth + right_depth`. Track the global maximum across all nodes.

- Time: O(n)
- Space: O(h) where h is the height of the tree

### Python

#### DFS (`diameterOfBinaryTree`)

Same recursive depth approach — compute left and right depths, update the global maximum diameter at each node.

- Time: O(n)
- Space: O(h) where h is the height of the tree
