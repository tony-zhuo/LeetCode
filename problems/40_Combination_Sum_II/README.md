# 40. Combination Sum II

## Problem

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used **once** in the combination.

**Note:** The solution set must not contain duplicate combinations.

### Examples

- `candidates = [10,1,2,7,6,1,5], target = 8` → `[[1,1,6],[1,2,5],[1,7],[2,6]]`
- `candidates = [2,5,2,1,2], target = 5` → `[[1,2,2],[5]]`

### Constraints

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## Solutions

### Go

#### Backtracking with sort + sibling dedup (`combinationSum2`)

Sort `candidates` so duplicate values are adjacent and siblings are visited in ascending order. At each recursion level, skip any candidate equal to its left neighbor *within the same level* (`i > start && candidates[i] == candidates[i-1]`) — this prevents picking the same value as a sibling, which would yield a duplicate combination, while still allowing duplicates to be reused at deeper levels (a different position in the path). Because each number may only be used once, recurse with `i + 1` rather than `i`. The sorted order also lets us break out of the loop as soon as `candidates[i] > num` (remaining target).

- Time: O(2^N) worst case — every element is either chosen or skipped — with pruning bringing the practical cost well below that.
- Space: O(N) auxiliary recursion depth and `path` length (output excluded).
