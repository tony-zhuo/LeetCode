# 278. First Bad Version

## Description

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

## Examples

**Example 1:**
```
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
So, 4 is the first bad version.
```

**Example 2:**
```
Input: n = 1, bad = 1
Output: 1
```

## Constraints

- `1 <= bad <= n <= 2^31 - 1`

## Solution

Use **binary search** to narrow down the first bad version:

1. Maintain `left = 1` and `right = n`.
2. Compute `mid = (left + right) / 2`.
3. If `isBadVersion(mid)` is true, the first bad version is at `mid` or earlier — set `right = mid`.
4. Otherwise, set `left = mid + 1`.
5. When `left == right`, that is the first bad version.

**Time complexity:** O(log n)
**Space complexity:** O(1)
