# Queue

## Overview
A queue is a linear data structure that follows the First-In-First-Out (FIFO) principle. This implementation includes both a standard slice-backed queue and a fixed-size circular queue (ring buffer).

## Operations & Time Complexity

### Standard Queue
| Operation | Time |
|-----------|------|
| Enqueue | O(1)* |
| Dequeue | O(n) |
| Peek | O(1) |
| IsEmpty | O(1) |
| Size | O(1) |

*Amortized O(1)

### Circular Queue
| Operation | Time |
|-----------|------|
| Enqueue | O(1) |
| Dequeue | O(1) |
| Peek | O(1) |
| IsEmpty | O(1) |
| IsFull | O(1) |
| Size | O(1) |

## Related LeetCode Problems
- No directly related solved problems yet
