# 78. Subsets

## Problem

Given an integer array `nums` of **unique** elements, return all possible subsets (the power set).

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

### Examples

- `nums = [1,2,3]` → `[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]`
- `nums = [0]` → `[[],[0]]`

### Constraints

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`
- All the numbers of `nums` are **unique**.

## Solutions

### Go

#### Backtracking (`subsets`)

Walk the decision tree once with a single mutable `path`. At every recursive call, snapshot the current `path` into `res` (the empty path produces `[]`, and every subsequent state is a distinct subset). Then for each index `i` from `start` to the end, append `nums[i]`, recurse with `i + 1` (so each element is used at most once and earlier indices are never revisited), and pop on the way back.

- Time: O(N × 2^N) — there are 2^N subsets and copying each takes up to O(N).
- Space: O(N) auxiliary recursion depth (output excluded).
