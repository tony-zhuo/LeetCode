# Segment Tree

## Overview
A segment tree is a binary tree used for storing information about intervals or segments. It allows efficient range queries (sum, minimum, maximum) and point updates.

## Operations & Time Complexity
| Operation | Time | Space |
|-----------|------|-------|
| Build | O(n) | O(4n) |
| Query | O(log n) | O(log n) |
| Update | O(log n) | O(log n) |

## Implementation Notes
- Array-based: root at index 1, left child = 2*i, right child = 2*i+1
- Two variants: range sum and range minimum
- Tree array size: 4 * n (safe upper bound)

## Related LeetCode Problems
- No directly related solved problems yet
