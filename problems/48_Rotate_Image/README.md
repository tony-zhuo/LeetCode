# 48. Rotate Image

## Problem Description
You are given an `n x n` 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image **in-place**, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

## Examples

### Example 1:
```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

### Example 2:
```
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

## Solution Approach
The solution uses a layer-by-layer rotation approach:

1. Process the matrix from outer layer to inner layer
2. For each layer, rotate 4 elements at a time in a cycle
3. Use a temporary variable to store one element while rotating the other 3

### Algorithm Steps:
1. Define left and right boundaries for the current layer
2. For each position in the current layer:
   - Save the top-left element
   - Move bottom-left → top-left
   - Move bottom-right → bottom-left  
   - Move top-right → bottom-right
   - Move saved top-left → top-right
3. Move to the next inner layer

## Complexity
- **Time Complexity**: O(n²) - visit each element once
- **Space Complexity**: O(1) - only use constant extra space

## Test Cases
Run tests with: `go test -v`

The test cases verify the rotation works correctly for different matrix sizes.
