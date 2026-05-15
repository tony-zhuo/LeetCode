# 46. Permutations

## Problem

Given an array `nums` of **distinct** integers, return all the possible permutations. You can return the answer in **any order**.

### Examples

- `nums = [1,2,3]` → `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`
- `nums = [0,1]` → `[[0,1],[1,0]]`
- `nums = [1]` → `[[1]]`

### Constraints

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are **unique**.

## Solutions

### Go

#### Backtracking with `used` set (`permute`)

Walk the decision tree by building `path` one element at a time. A boolean `used` array tracks which indices are already in the current path, so each level iterates over every index and skips those already chosen. When `len(path) == len(nums)`, snapshot the path into `res`; otherwise pick an unused index, recurse, then undo (`pop` from `path`, clear `used[i]`) before trying the next sibling.

- Time: O(N × N!) — there are N! permutations and copying each into the result is O(N).
- Space: O(N) auxiliary for the recursion stack, `path`, and `used` (output excluded).
