# 53. Maximum Subarray

## Problem

Given an integer array `nums`, find the subarray with the largest sum, and return its **sum**.

## Solutions

### Go

#### Kadane's Algorithm (`maxSubArray`)

Track the current subarray sum and a global maximum. At each element, decide whether to extend the current subarray or start fresh. Update the global maximum accordingly.

- Time: O(n)
- Space: O(1)

### Python

#### Kadane's Algorithm (`maxSubArray`)

Same approach — maintain a running sum, resetting when the current element alone is larger than the extended sum.

- Time: O(n)
- Space: O(1)
