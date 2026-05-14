# 39. Combination Sum

## Problem

Given an array of **distinct** integers `candidates` and a target integer `target`, return a list of all unique combinations of `candidates` where the chosen numbers sum to `target`. You may return the combinations in **any order**.

The **same** number may be chosen from `candidates` an **unlimited number of times**. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to `target` is fewer than 150 for the given input.

### Examples

- `candidates = [2,3,6,7], target = 7` → `[[2,2,3],[7]]`
- `candidates = [2,3,5], target = 8` → `[[2,2,2,2],[2,3,3],[3,5]]`
- `candidates = [2], target = 1` → `[]`

### Constraints

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- All elements of `candidates` are **distinct**.
- `1 <= target <= 40`

## Solutions

### Go

#### Backtracking with sort + prune (`combinationSum`)

Sort `candidates` so siblings are visited in ascending order; this lets us prune the entire tail of the loop as soon as a candidate exceeds the remaining target. At each level, recurse with `i` (not `i + 1`) so the same number can be reused, while only advancing forward through the array prevents permutations of an already-seen combination.

- Time: O(N^(T/M)) worst case, where `T` is `target` and `M` is the smallest candidate — bounded by the height of the search tree; the sort's O(N log N) is dominated.
- Space: O(T/M) auxiliary recursion depth and `path` length (output excluded).
