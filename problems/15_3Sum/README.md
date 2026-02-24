# 15. 3Sum

## Problem

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

## Solutions

### Go

#### Sorting + Two Pointers (`threeSum`)

Sort the array first. For each element `nums[i]`, use two pointers (`left`, `right`) to find pairs that sum to `-nums[i]`. Skip duplicates for both the outer loop and inner pointers to avoid duplicate triplets.

- Time: O(n^2)
- Space: O(1) (excluding output)
