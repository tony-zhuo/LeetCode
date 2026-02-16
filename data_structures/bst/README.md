# Binary Search Tree (BST)

## Overview
A binary search tree is a binary tree where for each node, all values in the left subtree are less than the node's value, and all values in the right subtree are greater. This property enables efficient search, insertion, and deletion.

## Operations & Time Complexity
| Operation | Average | Worst |
|-----------|---------|-------|
| Insert | O(log n) | O(n) |
| Delete | O(log n) | O(n) |
| Search | O(log n) | O(n) |
| Min/Max | O(log n) | O(n) |
| Inorder Traversal | O(n) | O(n) |

Worst case occurs with a skewed tree (essentially a linked list).

## Implementation Notes
- Delete uses in-order successor replacement for nodes with two children
- Duplicate values are ignored on insert

## Related LeetCode Problems
- [#98 Validate Binary Search Tree](../../problems/98_Validate_Binary_Search_Tree/)
- [#230 Kth Smallest Element in a BST](../../problems/230_Kth_Smallest_Element_in_a_BST/)
