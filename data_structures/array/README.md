# Array (Dynamic Array)

## Overview
A dynamic array (similar to Go's slice or Java's ArrayList) that manages its own capacity. Elements are stored in a contiguous block of memory and the array automatically grows/shrinks as needed.

## Operations & Time Complexity
| Operation | Time | Space |
|-----------|------|-------|
| Get | O(1) | O(1) |
| Set | O(1) | O(1) |
| Append | O(1)* | O(1) |
| InsertAt | O(n) | O(1) |
| DeleteAt | O(n) | O(1) |
| Contains | O(n) | O(1) |
| IndexOf | O(n) | O(1) |
| Reverse | O(n) | O(1) |
| Size | O(1) | O(1) |

*Amortized O(1), doubles capacity when full

## Related LeetCode Problems
- [#1 Two Sum](../../problems/1_Two_Sum/)
- [#53 Maximum Subarray](../../problems/53_Maximum_Subarray/)
- [#217 Contains Duplicate](../../problems/217_Contains_Duplicate/)
