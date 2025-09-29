# 371. Sum of Two Integers

## Problem Description

Given two integers `a` and `b`, return the sum of the two integers without using the operators `+` and `-`.

### Examples

**Example 1:**
```
Input: a = 1, b = 2
Output: 3
```

**Example 2:**
```
Input: a = 2, b = 3
Output: 5
```

### Constraints

- `-1000 <= a, b <= 1000`

## Solution

This problem requires implementing addition using only bitwise operations. The key insight is that addition can be broken down into two parts:

1. **XOR (^)**: Gives the sum without carry
2. **AND (&) + Left Shift (<<)**: Calculates the carry

### Algorithm Steps

1. Use XOR to get the sum without considering carry: `a ^ b`
2. Use AND to find positions where carry occurs: `a & b`
3. Left shift the carry by 1 position: `carry << 1`
4. Repeat until there's no carry (carry becomes 0)

### Example Walkthrough

Adding 1 + 2:
```
  a = 1 (binary: 001)
  b = 2 (binary: 010)

  Step 1: a ^ b = 001 ^ 010 = 011 (3)
          carry = (a & b) << 1 = (001 & 010) << 1 = 000 << 1 = 000

  Since carry = 0, result = 3
```

Adding 3 + 5:
```
  a = 3 (binary: 011)
  b = 5 (binary: 101)

  Step 1: a ^ b = 011 ^ 101 = 110 (6)
          carry = (011 & 101) << 1 = 001 << 1 = 010 (2)

  Step 2: a = 6 (binary: 110), b = 2 (binary: 010)
          a ^ b = 110 ^ 010 = 100 (4)
          carry = (110 & 010) << 1 = 010 << 1 = 100 (4)

  Step 3: a = 4 (binary: 100), b = 4 (binary: 100)
          a ^ b = 100 ^ 100 = 000 (0)
          carry = (100 & 100) << 1 = 100 << 1 = 1000 (8)

  Step 4: a = 0 (binary: 000), b = 8 (binary: 1000)
          a ^ b = 000 ^ 1000 = 1000 (8)
          carry = (000 & 1000) << 1 = 0000 << 1 = 0000 (0)

  Since carry = 0, result = 8
```

### Time Complexity
- **Time:** O(1) - The loop runs at most 32 times (for 32-bit integers)
- **Space:** O(1) - Only using constant extra space

### Key Concepts
- **Bit Manipulation**: Using XOR for addition without carry
- **Carry Propagation**: Using AND and left shift to handle carries
- **Iterative Process**: Continue until no more carries remain
