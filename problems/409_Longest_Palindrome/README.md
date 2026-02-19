# 409. Longest Palindrome

## Problem

Given a string `s` which consists of lowercase or uppercase letters, return the length of the **longest palindrome** that can be built with those letters.

Letters are **case sensitive** — `"Aa"` is not considered a palindrome.

## Solutions

### Go

#### 1. Map (`longestPalindrome_map`)

Count character frequencies with a hash map. Sum up even portions of each count; if any odd count exists, add 1 for the center.

- Time: O(n)
- Space: O(k) where k is the number of distinct characters

#### 2. Array (`longestPalindrome_array`)

Same logic but uses a fixed `[128]int` array indexed by ASCII value instead of a map.

- Time: O(n)
- Space: O(1)

### Python

#### 1. Array (`longestPalindrome`)

Uses a `[128]` frequency array, same approach as the Go array solution.

- Time: O(n)
- Space: O(1)

#### 2. Counter (`longestPalindrome_counter`)

Uses `collections.Counter` for counting, then applies the same even-pair logic.

- Time: O(n)
- Space: O(k)
