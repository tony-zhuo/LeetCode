# 190. Reverse Bits

## Problem Description
Reverse bits of a given 32 bits unsigned integer.

**Note:**
- Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in **Example 2** below, the input represents the signed integer `-3` and the output represents the signed integer `-1073741825`.

## Examples

### Example 1:
```
Input: n = 00000010100101000001111010011100
Output:   964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.
```

### Example 2:
```
Input: n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.
```

## Solution Approach
The solution uses bit manipulation to reverse the 32-bit integer:

### Algorithm Steps:
1. Initialize `result = 0`
2. For each of the 32 bits:
   - Left shift `result` by 1 position to make room for the next bit
   - Extract the least significant bit from `n` using `(n & 1)`
   - Add this bit to `result` using OR operation `result |= (n & 1)`
   - Right shift `n` by 1 position to process the next bit
3. Return the final `result`

### Key Operations:
- `result <<= 1`: Shift result left to make room for next bit
- `n & 1`: Extract the rightmost bit of n
- `result |= (n & 1)`: Add the extracted bit to result
- `n >>= 1`: Move to the next bit in n

### Example Trace:
For input `n = 5` (binary: `00000000000000000000000000000101`):
```
Initial: result = 0, n = 5 (101)
Iteration 1: result = 0 → 0 → 1, n = 5 → 2
Iteration 2: result = 1 → 2 → 2, n = 2 → 1  
Iteration 3: result = 2 → 4 → 5, n = 1 → 0
...continue for remaining 29 iterations
Final result: 2684354560 (10100000000000000000000000000000)
```

## Complexity
- **Time Complexity**: O(1) - always process exactly 32 bits
- **Space Complexity**: O(1) - only use constant extra space

## Test Cases
Run tests with: `go test -v`

The test cases cover:
- LeetCode official examples
- Edge cases (all zeros, all ones)
- Boundary cases (single bit at start/end positions)
