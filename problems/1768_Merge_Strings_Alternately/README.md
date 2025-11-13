# 1768. Merge Strings Alternately

You are given two strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

**Example 1:**

> Input: word1 = "abc", word2 = "pqr"
> Output: "apbqcr"
> Explanation: The merged string will be merged as so:
> - word1:  a   b   c
> - word2:    p   q   r
> - merged: a p b q c r

**Example 2:**

> Input: word1 = "ab", word2 = "pqrs"
> Output: "apbqrs"
> Explanation: Notice that as word2 is longer, "rs" is appended to the end.
> - word1:  a   b
> - word2:    p   q   r   s
> - merged: a p b q   r   s

**Example 3:**

> Input: word1 = "abcd", word2 = "pq"
> Output: "apbqcd"
> Explanation: Notice that as word1 is longer, "cd" is appended to the end.
> - word1:  a   b   c   d
> - word2:    p   q
> - merged: a p b q c   d

**Constraints:**

- `1 <= word1.length, word2.length <= 100`
- `word1` and `word2` consist of lowercase English letters.

## Solution Approach

This problem can be solved using a **two-pointer** technique with efficient string building.

### Algorithm Steps:

1. **Initialize two pointers**: `word1Pointer` and `word2Pointer` both starting at 0
2. **Use strings.Builder**: For efficient string concatenation in Go
3. **Iterate while either pointer is valid**: Continue as long as either pointer hasn't reached the end of its string
4. **Alternately append characters**:
   - If `word1Pointer < len(word1)`, append `word1[word1Pointer]` and increment pointer
   - If `word2Pointer < len(word2)`, append `word2[word2Pointer]` and increment pointer
5. **Return the built string**

### Why this works:

The loop condition `word1Pointer < len(word1) || word2Pointer < len(word2)` ensures we process all characters from both strings. By checking each pointer individually before appending, we naturally handle cases where one string is longer than the other.

### Execution Example:

Let's trace through **Example 2**: `word1 = "ab", word2 = "pqrs"`

**Initial state:**
```
word1Pointer = 0, word2Pointer = 0
result = ""
```

**Iteration 1:**
- word1Pointer (0) < len(word1) (2) ✓ → append 'a', word1Pointer = 1
- word2Pointer (0) < len(word2) (4) ✓ → append 'p', word2Pointer = 1
- result = "ap"

**Iteration 2:**
- word1Pointer (1) < len(word1) (2) ✓ → append 'b', word1Pointer = 2
- word2Pointer (1) < len(word2) (4) ✓ → append 'q', word2Pointer = 2
- result = "apbq"

**Iteration 3:**
- word1Pointer (2) < len(word1) (2) ✗ → skip
- word2Pointer (2) < len(word2) (4) ✓ → append 'r', word2Pointer = 3
- result = "apbqr"

**Iteration 4:**
- word1Pointer (2) < len(word1) (2) ✗ → skip
- word2Pointer (3) < len(word2) (4) ✓ → append 's', word2Pointer = 4
- result = "apbqrs"

**Loop ends** (both pointers reached their limits)

## Complexity Analysis

- **Time Complexity**: O(n + m) - where n is the length of word1 and m is the length of word2. We visit each character exactly once.
- **Space Complexity**: O(n + m) - for the result string stored in strings.Builder.

## Key Insights

- **Two-pointer technique** is ideal for processing two sequences simultaneously
- **strings.Builder** in Go is more efficient than string concatenation with `+` operator
- The condition `||` (OR) in the loop ensures we handle strings of different lengths elegantly
- No need for special handling of remaining characters - the loop naturally processes them
- Similar problems: Merge Two Sorted Lists (21), Interleaving String (97)
