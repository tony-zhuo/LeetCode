# 3. Longest Substring Without Repeating Characters

## Problem

Given a string `s`, find the length of the longest substring without repeating characters.

## Solutions

### Go

#### Sliding Window with Set (`lengthOfLongestSubstring_2On`)

Use two pointers and a hash set. Expand the right pointer; when a duplicate is found, shrink from the left by removing characters until the duplicate is gone.

- Time: O(2n) worst case — each character is visited at most twice
- Space: O(min(n, m)) where m is the character set size

#### Sliding Window with Index Map (`lengthOfLongestSubstring_On`)

Use two pointers and a hash map storing each character's last seen index. When a duplicate is found, jump the left pointer directly past the previous occurrence instead of shrinking one by one.

- Time: O(n)
- Space: O(min(n, m)) where m is the character set size
