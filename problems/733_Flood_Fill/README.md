# 733. Flood Fill

## Problem Description

An image is represented by an `m x n` integer grid `image` where `image[i][j]` represents the pixel value of the image.

You are also given three integers `sr`, `sc`, and `color`. You should perform a **flood fill** on the image starting from the pixel `image[sr][sc]`.

To perform a flood fill, consider the starting pixel, plus any pixels connected **4-directionally** to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with `color`.

Return the modified image after performing the flood fill.

### Examples

**Example 1:**
```
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center pixel (sr=1, sc=1) with value 1, all connected 1s are changed to 2.
The bottom-right 1 is not connected to the starting pixel (blocked by 0s).
```

**Example 2:**
```
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
Output: [[0,0,0],[0,0,0]]
Explanation: The starting pixel already has the target color, so no changes are made.
```

### Constraints

- `m == image.length`
- `n == image[i].length`
- `1 <= m, n <= 50`
- `0 <= image[i][j], color < 2^16`
- `0 <= sr < m`
- `0 <= sc < n`

## Solution

### DFS (Recursive)

Use depth-first search to flood fill from the starting pixel:

1. Record the original color at `(sr, sc)`
2. If the original color equals the new color, return immediately (avoids infinite recursion)
3. Set the current pixel to the new color
4. Recursively fill the 4 adjacent pixels (up, down, left, right) that match the original color

### Time Complexity
- **Time:** O(m * n) — each pixel is visited at most once
- **Space:** O(m * n) — recursion stack in the worst case
