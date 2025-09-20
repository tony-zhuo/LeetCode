# 73. Set Matrix Zeroes

## Problem Description

Given an `m x n` integer matrix `matrix`, if an element is `0`, set its entire row and column to `0`'s.

You must do it **in place**.

## Examples

### Example 1:
```
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

### Example 2:
```
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
```

## Constraints

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `-2^31 <= matrix[i][j] <= 2^31 - 1`

## Algorithm

The solution uses a two-pass approach:

1. **First pass**: Identify which rows and columns contain zeros
   - Use two boolean arrays: `rows_zero` and `columns_zero`
   - Scan the entire matrix and mark corresponding positions in these arrays

2. **Second pass**: Set zeros based on the marked positions
   - For each row marked as containing zero, set all elements in that row to 0
   - For each column marked as containing zero, set all elements in that column to 0

## Complexity

- **Time Complexity**: O(m × n) where m is the number of rows and n is the number of columns
- **Space Complexity**: O(m + n) for the two boolean arrays

## Follow-up

A more space-efficient approach would be to use the first row and first column of the matrix itself as markers, achieving O(1) extra space complexity.
