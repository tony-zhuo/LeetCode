# 242. Valid Anagram

## Problem Description

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Examples

**Example 1:**
```
Input: s = "anagram", t = "nagaram"
Output: true
```

**Example 2:**
```
Input: s = "rat", t = "car"
Output: false
```

### Constraints

- `1 <= s.length, t.length <= 5 * 10^4`
- `s` and `t` consist of lowercase English letters.

## Solution

The solution uses a hash map to count the frequency of each character in both strings:

1. First check if the strings have different lengths - if so, they cannot be anagrams
2. Create a map to count character frequencies in the first string
3. Iterate through the second string, decrementing the count for each character
4. If any character count goes negative, the strings are not anagrams
5. If we complete the iteration without any negative counts, the strings are anagrams

### Time Complexity
- **Time:** O(n) where n is the length of the strings
- **Space:** O(1) since we're using at most 26 characters (lowercase English letters)

### Alternative Approaches
1. **Sorting:** Sort both strings and compare - O(n log n) time
2. **Array counting:** Use fixed-size array instead of map - same complexity but potentially faster
3. **Character frequency comparison:** Build frequency maps for both strings and compare
