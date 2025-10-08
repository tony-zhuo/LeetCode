# 252. Meeting Rooms

## Problem Description

Given an array of meeting time intervals where `intervals[i] = [starti, endi]`, determine if a person could attend all meetings.

## Examples

### Example 1:
```
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false
Explanation: The first meeting [0,30] conflicts with both the second meeting [5,10] and third meeting [15,20].
```

### Example 2:
```
Input: intervals = [[7,10],[2,4]]
Output: true
Explanation: The meetings don't overlap, so a person can attend all meetings.
```

## Constraints

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti < endi <= 10^6`

## Approach

1. **Sort by start time**: Sort all intervals by their start times
2. **Check overlaps**: Iterate through sorted intervals and check if any meeting starts before the previous one ends
3. **Return result**: If any overlap is found, return false; otherwise return true

## Time Complexity
- **Time**: O(n log n) - due to sorting
- **Space**: O(1) - only using constant extra space

## Algorithm Steps

1. Sort the intervals array by start time
2. For each pair of consecutive meetings, check if the current meeting starts before the previous meeting ends
3. If overlap found (`intervals[i][0] < intervals[i-1][1]`), return false
4. If no overlaps found after checking all pairs, return true
