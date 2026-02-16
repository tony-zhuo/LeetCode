# 20. Valid Parentheses

## Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Examples

**Example 1:**
```
Input: s = "()"
Output: true
```

**Example 2:**
```
Input: s = "()[]{}"
Output: true
```

**Example 3:**
```
Input: s = "(]"
Output: false
```

**Example 4:**
```
Input: s = "([])"
Output: true
```

## Constraints

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`

## Solution

Use a **stack** to track expected closing brackets:

1. For each opening bracket, push the corresponding closing bracket onto the stack.
2. For each closing bracket, pop from the stack and check if it matches.
3. If the stack is empty when encountering a closing bracket, or the popped value doesn't match, return `false`.
4. After processing all characters, return `true` only if the stack is empty.

**Time complexity:** O(n)
**Space complexity:** O(n)
