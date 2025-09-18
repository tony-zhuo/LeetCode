# 54. Spiral Matrix

## Problem Description
Given an `m x n` matrix, return all elements of the matrix in spiral order.

## Examples

### Example 1:
```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

### Example 2:
```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

## Solution Approach
The solution uses a boundary-based approach to traverse the matrix in spiral order:

1. Define four boundaries: top, bottom, left, right
2. Traverse in four directions cyclically:
   - Left to right along the top row
   - Top to bottom along the right column  
   - Right to left along the bottom row (if valid)
   - Bottom to top along the left column (if valid)
3. Shrink boundaries after each direction
4. Continue until all elements are visited

### Algorithm Steps:
1. Initialize boundaries: `top=0, bottom=n-1, left=0, right=m-1`
2. **Right direction**: Traverse from left to right along top row, then increment top
3. **Down direction**: Traverse from top to bottom along right column, then decrement right
4. **Left direction**: If still valid, traverse from right to left along bottom row, then decrement bottom
5. **Up direction**: If still valid, traverse from bottom to top along left column, then increment left
6. Repeat until `left > right` or `top > bottom`

### Key Points:
- Check boundaries before traversing left and up directions to avoid duplicate elements
- Use `if top <= bottom` before left traversal
- Use `if left <= right` before up traversal

## Complexity
- **Time Complexity**: O(m × n) - visit each element exactly once
- **Space Complexity**: O(1) - only use constant extra space (excluding result array)

## Test Cases
Run tests with: `go test -v`
