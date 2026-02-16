# Union-Find (Disjoint Set)

## Overview
Union-Find is a data structure that tracks a set of elements partitioned into disjoint (non-overlapping) subsets. It supports efficient union and find operations, commonly used for connectivity problems.

## Operations & Time Complexity
| Operation | Time |
|-----------|------|
| Find | O(α(n)) ≈ O(1) |
| Union | O(α(n)) ≈ O(1) |
| Connected | O(α(n)) ≈ O(1) |
| Count | O(1) |

α(n) = inverse Ackermann function, practically constant

## Implementation Notes
- Path compression in Find: directly connects nodes to root
- Union by rank: attaches shorter tree under taller tree
- Combined optimizations give amortized nearly O(1) per operation

## Related LeetCode Problems
- No directly related solved problems yet
