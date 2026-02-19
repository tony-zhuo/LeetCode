# 383. Ransom Note

## Description

Given two strings `ransomNote` and `magazine`, return `true` if `ransomNote` can be constructed by using the letters from `magazine` and `false` otherwise.

Each letter in `magazine` can only be used once in `ransomNote`.

## Examples

**Example 1:**
```
Input: ransomNote = "a", magazine = "b"
Output: false
```

**Example 2:**
```
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

**Example 3:**
```
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

## Constraints

- `1 <= ransomNote.length, magazine.length <= 10^5`
- `ransomNote` and `magazine` consist of lowercase English letters.

## Solution

Use a **character frequency array** of size 26:

1. Count each character in `magazine`.
2. Subtract each character in `ransomNote`.
3. If any count goes below zero, return `false`.

**Time complexity:** O(m + n) where m and n are the lengths of magazine and ransomNote
**Space complexity:** O(1) — fixed 26-element array
