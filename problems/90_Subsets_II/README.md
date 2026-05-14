# 90. Subsets II

## Problem

Given an integer array `nums` that **may contain duplicates**, return all possible subsets (the power set).

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

### Examples

- `nums = [1,2,2]` → `[[],[1],[1,2],[1,2,2],[2],[2,2]]`
- `nums = [0]` → `[[],[0]]`

### Constraints

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`

## Solutions

### Go

#### Backtracking with sort + sibling skip (`subsetsWithDup`)

Sort `nums` first so duplicates sit next to each other. Then run the same backtracking decision tree used in [#78 Subsets](../78_Subsets/README.md) with one extra guard: at each level, skip an index `i > start` whose value equals `nums[i-1]`. The first occurrence of each duplicate value already explores every continuation, so re-using a later duplicate at the same depth would generate an identical subset.

- Time: O(N × 2^N) worst case — bounded by the unique-element case; duplicates only prune.
- Space: O(N) auxiliary recursion depth plus the sort's O(log N) stack (output excluded).
