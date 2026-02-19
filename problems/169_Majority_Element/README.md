# 169. Majority Element

## Problem

Given an array `nums` of size `n`, return the **majority element**.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

## Solutions

### Go

#### Boyer-Moore Voting (`majorityElement`)

Maintain a candidate and a counter. When the counter drops to 0, pick the current element as the new candidate. Increment on match, decrement on mismatch. The majority element is guaranteed to survive.

- Time: O(n)
- Space: O(1)

### Python

#### 1. Counter (`majorityElement`)

Use `collections.Counter` to count frequencies, then return the key with the highest count.

- Time: O(n)
- Space: O(n)

#### 2. Boyer-Moore Voting (`majorityElement_BoyerMoore`)

Same approach as the Go solution — single pass with a candidate and counter.

- Time: O(n)
- Space: O(1)
