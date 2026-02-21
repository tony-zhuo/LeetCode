# 57. Insert Interval

## Problem

You are given an array of non-overlapping intervals `intervals` where `intervals[i] = [starti, endi]` represent the start and end of the `ith` interval and `intervals` is sorted in ascending order by `starti`. You are also given an interval `newInterval = [start, end]`.

Insert `newInterval` into `intervals` such that `intervals` is still sorted in ascending order by `starti` and `intervals` still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return `intervals` after the insertion.

## Solutions

### Go

#### Linear Scan (`insert`)

Iterate through intervals. Append non-overlapping intervals on the left directly. For overlapping intervals, expand `newInterval` by taking the min start and max end. Once `newInterval` is fully past the current interval, append it and continue with the rest.

- Time: O(n)
- Space: O(n)
