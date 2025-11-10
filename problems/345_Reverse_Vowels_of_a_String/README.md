# 345. Reverse Vowels of a String

## Problem Description

Given a string `s`, reverse only all the vowels in the string and return it.

The vowels are `'a'`, `'e'`, `'i'`, `'o'`, and `'u'`, and they can appear in both lower and upper cases, more than once.

## Examples

**Example 1:**
```
Input: s = "hello"
Output: "holle"
```

**Example 2:**
```
Input: s = "leetcode"
Output: "leotcede"
```

**Example 3:**
```
Input: s = "aA"
Output: "Aa"
```

## Constraints

- `1 <= s.length <= 3 * 10^5`
- `s` consist of printable ASCII characters.

## Solution Approach

This problem uses the **two-pointer technique**:

1. Convert string to byte array for in-place manipulation
2. Use left and right pointers starting from both ends
3. Move pointers inward until both point to vowels
4. Swap the vowels and continue until pointers meet

## Algorithm Steps

1. Initialize left pointer at start (0) and right pointer at end (len-1)
2. While left < right:
   - If left character is not a vowel, move left pointer right
   - If right character is not a vowel, move right pointer left  
   - If both characters are vowels, swap them and move both pointers
3. Return the modified string

## Time Complexity
- **O(n)** where n is the length of the string
- Each character is visited at most once

## Space Complexity
- **O(n)** for converting string to byte array in Go
- The algorithm itself uses O(1) extra space for pointers
