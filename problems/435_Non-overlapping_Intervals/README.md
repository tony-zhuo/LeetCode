# 435. Non-overlapping Intervals

Given an array of intervals `intervals` where `intervals[i] = [starti, endi]`, return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

**Note** that intervals which only touch at a point are non-overlapping. For example, `[1, 2]` and` [2, 3]` are non-overlapping.

**Example 1:**

>Input: intervals = [[1, 2], [2, 3], [3, 4], [1, 3]]  
> Output: 1  
> Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.  

**Example 2:**

> Input: intervals = [[1, 2], [1, 2], [1, 2]]  
> Output: 2  
> Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.  

**Example 3:**

> Input: intervals = [[1, 2], [2, 3]]  
> Output: 0  
> Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

**Constraints:**

- `1 <= intervals.length <= 105`
- `intervals[i].length == 2`
- `-5 * 104 <= starti < endi <= 5 * 104`

## Solution Approach

This problem can be solved using a **greedy algorithm** with the following strategy:

1. **Sort intervals by end time**: Sort all intervals by their ending points in ascending order
2. **Select non-overlapping intervals**: Greedily select intervals that don't overlap with previously selected ones
3. **Calculate removals**: The answer is total intervals minus the maximum number of non-overlapping intervals we can keep

### Algorithm Steps:
1. Sort intervals by end time (`intervals[i][1]`)
2. Initialize count of non-overlapping intervals to 1 (first interval)
3. Keep track of the end time of the last selected interval
4. For each subsequent interval:
   - If its start time >= last end time, it's non-overlapping
   - Update count and end time
5. Return `total intervals - non-overlapping count`

### Why this works:
By always choosing the interval that ends earliest among all available choices, we leave the most room for future intervals, maximizing the number of intervals we can keep.

## Complexity Analysis

- **Time Complexity**: O(n log n) - dominated by sorting
- **Space Complexity**: O(1) - only using constant extra space (sorting is in-place)

## Key Insights

- This is a classic **interval scheduling maximization** problem
- The greedy approach of selecting intervals with earliest end times is optimal
- Non-overlapping means intervals can touch at endpoints (e.g., [1,2] and [2,3])
