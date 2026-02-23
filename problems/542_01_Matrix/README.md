# 542. 01 Matrix

## Problem

Given an `m x n` binary matrix `mat`, return the distance of the nearest `0` for each cell.

The distance between two adjacent cells is `1`.

## Solutions

### Go

#### BFS (`updateMatrix`)

Start BFS from all `0` cells simultaneously (multi-source BFS). Initialize a result matrix with `0` for zero-cells and `-1` for one-cells. Process the queue level by level, setting each unvisited neighbor's distance to `current + 1`.

- Time: O(m × n)
- Space: O(m × n)
