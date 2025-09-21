# 338. Counting Bits

## Problem Description

Given an integer `n`, return an array `ans` of length `n + 1` such that for each `i` (`0 <= i <= n`), `ans[i]` is the **number of 1's** in the binary representation of `i`.

## Examples

### Example 1:
```
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
```

### Example 2:
```
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101
```

## Constraints

- `0 <= n <= 10^5`

## Algorithm

This implementation uses a **brute force approach** with **Brian Kernighan's Algorithm** for counting bits in each number.

### Approach:
1. Create a result array of size `n+1`
2. For each number `i` from 1 to `n`:
   - Use Brian Kernighan's algorithm to count set bits
   - Store the count in `res[i]`

### Brian Kernighan's Bit Counting:
For each number, the algorithm uses `tmp & (tmp - 1)` to remove the rightmost set bit:
```go
cnt := 0
tmp := i
for tmp > 0 {
    cnt++
    tmp = tmp & (tmp - 1)  // Remove rightmost set bit
}
```

### Example walkthrough with n = 5:
```
i=0: 0 (0b0)     → 0 bits
i=1: 1 (0b1)     → 1 bit
i=2: 2 (0b10)    → 1 bit  
i=3: 3 (0b11)    → 2 bits
i=4: 4 (0b100)   → 1 bit
i=5: 5 (0b101)   → 2 bits
Result: [0,1,1,2,1,2]
```

## Complexity

- **Time Complexity**: O(n × k) where k is the average number of set bits
  - For each number i, Brian Kernighan's algorithm runs in O(number of set bits)
  - In worst case: O(n × log n)
- **Space Complexity**: O(1) extra space (not counting the output array)

## Pattern Recognition

The result shows interesting patterns:
- **Powers of 2** (1, 2, 4, 8, 16...): Always have exactly 1 bit
- **Consecutive ranges**: Each power of 2 starts a new pattern
- **Symmetry**: `ans[i] = ans[i - power_of_2] + 1` for certain patterns

## Optimized Approaches (Alternative)

While this implementation works correctly, there are more efficient dynamic programming solutions:

1. **DP with right shift**: `ans[i] = ans[i >> 1] + (i & 1)`
2. **DP with power of 2**: `ans[i] = ans[i - power] + 1`

These achieve O(n) time complexity but the current brute force approach is easier to understand and implement.
