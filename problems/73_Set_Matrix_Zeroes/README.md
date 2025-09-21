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

The solution uses an optimized O(1) space approach by using the first row and column as markers:

1. **Check if first row/column should be zeroed**: 
   - Record if the first row contains any zeros (`is_first_rows_zero`)
   - Record if the first column contains any zeros (`is_first_columns_zero`)

2. **Use first row and column as markers**:
   - Scan the matrix starting from `matrix[1][1]`
   - If `matrix[i][j] == 0`, mark `matrix[0][j] = 0` and `matrix[i][0] = 0`

3. **Set zeros based on markers**:
   - For each cell `matrix[i][j]` (starting from `[1][1]`), if either `matrix[0][j] == 0` or `matrix[i][0] == 0`, set `matrix[i][j] = 0`

4. **Handle first row and column**:
   - If `is_first_rows_zero` is true, set entire first row to zeros
   - If `is_first_columns_zero` is true, set entire first column to zeros

## Complexity

- **Time Complexity**: O(m × n) where m is the number of rows and n is the number of columns
- **Space Complexity**: O(1) - only using constant extra space for flags

## Key Insights

- The first row and column serve dual purposes: they contain original data and act as markers
- We must handle the first row/column separately since they're used as markers
- This approach achieves the optimal space complexity without additional data structures
