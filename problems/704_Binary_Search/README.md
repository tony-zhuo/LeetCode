# 704. Binary Search

## Problem Description

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

### Examples

**Example 1:**
```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
```

**Example 2:**
```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
```

### Constraints

- `1 <= nums.length <= 10^4`
- `-10^4 < nums[i], target < 10^4`
- All the integers in `nums` are unique.
- `nums` is sorted in ascending order.

## Solution

### Recursive

Recursively divide the array in half. Compare the target with the middle element and recurse into the appropriate half.

1. Find the middle index of the array
2. If the middle element equals target, return the index
3. If target is greater, recurse on the right half and adjust the returned index
4. Otherwise, recurse on the left half

### Two Pointer (Iterative)

Use two pointers (`left` and `right`) to narrow the search range iteratively.

1. Initialize `left = 0` and `right = len(nums) - 1`
2. While `left <= right`, compute the middle index
3. If the middle element equals target, return the index
4. If target is greater, move `left` to `mid + 1`
5. Otherwise, move `right` to `mid - 1`
6. If the loop exits, the target is not found — return `-1`

### Time Complexity
- **Time:** O(log n)
- **Space:** O(log n) for recursive, O(1) for iterative