# Heap / Priority Queue

## Overview
A heap is a complete binary tree stored in an array where each parent node satisfies the heap property (min-heap: parent <= children, max-heap: parent >= children). It is commonly used to implement priority queues.

## Operations & Time Complexity
| Operation | Time |
|-----------|------|
| Insert | O(log n) |
| ExtractMin/Max | O(log n) |
| Peek | O(1) |
| Heapify | O(n) |
| Size | O(1) |

## Implementation Notes
- Array-based: parent = (i-1)/2, children = 2i+1 and 2i+2
- Heapify uses bottom-up siftDown for O(n) construction

## Related LeetCode Problems
- [#253 Meeting Rooms II](../../problems/253_Meeting_Rooms_II/)
- [#347 Top K Frequent Elements](../../problems/347_Top_K_Frequent_Elements/)
