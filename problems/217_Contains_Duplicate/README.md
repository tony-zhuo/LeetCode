# 217. Contains Duplicate

## Problem

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

## Solutions

### Go

#### Hash Set (`containsDuplicate`)

Iterate through the array, storing each number in a map. If a number is already present, return true immediately.

- Time: O(n)
- Space: O(n)

### Python

#### 1. Counter (`containsDuplicate_counter`)

Use `collections.Counter` to count frequencies, then check if any count exceeds 1.

- Time: O(n)
- Space: O(n)

#### 2. Set (`containsDuplicate_set`)

Compare the length of the array with the length of its set. If they differ, duplicates exist.

- Time: O(n)
- Space: O(n)
