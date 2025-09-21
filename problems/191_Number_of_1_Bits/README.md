# 191. Number of 1 Bits

## Problem Description

Write a function that takes the binary representation of a positive integer and returns the number of set bits it has (also known as the Hamming weight).

## Examples

### Example 1:
```
Input: n = 11
Output: 3
Explanation: The input binary string 1011 has a total of three set bits.
```

### Example 2:
```
Input: n = 128
Output: 1
Explanation: The input binary string 10000000 has a total of one set bit.
```

### Example 3:
```
Input: n = 2147483645
Output: 30
Explanation: The input binary string 01111111111111111111111111111101 has a total of thirty set bits.
```

## Constraints

- `1 <= n <= 2^31 - 1`

## Algorithm

This implementation uses **Brian Kernighan's Algorithm**, which is more efficient than the naive bit-by-bit checking approach.

### Key Insight
The operation `n & (n-1)` removes the rightmost set bit from `n`.

### How it works:
1. Start with a counter `cnt = 0`
2. While `n > 0`:
   - Increment the counter (`cnt++`)
   - Remove the rightmost set bit using `n = n & (n-1)`
3. Return the counter

### Example walkthrough with n = 11 (binary: 1011):
```
Initial: n = 11 (1011), cnt = 0
Step 1:  n = 11 & 10 = 10 (1010), cnt = 1
Step 2:  n = 10 & 9 = 8 (1000), cnt = 2  
Step 3:  n = 8 & 7 = 0 (0000), cnt = 3
Result: cnt = 3
```

## Complexity

- **Time Complexity**: O(k) where k is the number of set bits in n
  - In the worst case, O(log n) when all bits are set
- **Space Complexity**: O(1) - only using constant extra space

## Alternative Approach (commented out)

The code also includes a commented-out naive approach:
- Iterate through all 32 bits
- Check each bit using `(n & 1) == 1`
- Right shift by 1 position each iteration
- Time complexity: O(32) = O(1), but slower in practice when there are few set bits

## Why Brian Kernighan's Algorithm is Better

- **Fewer iterations**: Only loops for the number of set bits, not all bits
- **Efficient bit manipulation**: Directly removes set bits instead of checking each position
- **Optimal for sparse bit patterns**: Performs exceptionally well when there are few set bits
