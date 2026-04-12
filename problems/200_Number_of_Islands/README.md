# 200. Number of Islands

## Problem

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Solutions

### Go

#### DFS (`numIslands`)

Iterate through every cell. When a `'1'` is found, increment the count and flood-fill the entire island to `'0'` using recursive DFS in all four directions so it won't be counted again.

- Time: O(m × n)
- Space: O(m × n) worst-case recursion depth
