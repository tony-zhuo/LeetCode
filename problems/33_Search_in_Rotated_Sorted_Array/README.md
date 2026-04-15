# 33. Search in Rotated Sorted Array

## Problem

There is an integer array `nums` sorted in ascending order (with distinct values), possibly rotated at an unknown pivot index `k` (`0 <= k < nums.length`) so that the resulting array becomes `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`.

Given the rotated array `nums` and an integer `target`, return the index of `target` if it is in `nums`, or `-1` if it is not.

You must write an algorithm with `O(log n)` runtime complexity.

## Solutions

### Go

#### Modified Binary Search (`search`)

Standard binary search, but at each step decide which half is sorted by comparing `nums[mid]` with `nums[left]`:

- If `nums[mid] >= nums[left]`, the left half `[left, mid]` is sorted. If `target` lies within `[nums[left], nums[mid])`, search left; otherwise search right.
- Otherwise the right half `[mid, right]` is sorted. If `target` lies within `(nums[mid], nums[right]]`, search right; otherwise search left.

This narrows the search to the half guaranteed to contain `target` (if present) in `O(log n)` time.

- Time: O(log n)
- Space: O(1)