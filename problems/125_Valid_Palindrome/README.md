# 125. Valid Palindrome

## Description

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

## Examples

**Example 1:**
```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

**Example 2:**
```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```

**Example 3:**
```
Input: s = " "
Output: true
Explanation: After removing non-alphanumeric characters, s is an empty string "".
```

## Constraints

- `1 <= s.length <= 2 * 10^5`
- `s` consists only of printable ASCII characters.

## Solution

Use **two pointers** from both ends:

1. Skip non-alphanumeric characters from both sides.
2. Compare lowercase versions of the characters at left and right pointers.
3. If any mismatch, return `false`. If pointers meet, return `true`.

**Time complexity:** O(n)
**Space complexity:** O(1)
