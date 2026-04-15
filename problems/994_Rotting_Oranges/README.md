# 994. Rotting Oranges

## Problem

You are given an `m x n` grid where each cell can have one of three values:

- `0` representing an empty cell,
- `1` representing a fresh orange, or
- `2` representing a rotten orange.

Every minute, any fresh orange that is **4-directionally adjacent** to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return `-1`.

## Solutions

### Go

#### Multi-source BFS (`orangesRotting`)

Push every initially rotten orange into a queue and track the count of fresh oranges. Process the queue level by level: each level represents one minute. For each rotten orange popped, rot its 4-directional fresh neighbors, enqueue them, and decrement the fresh count. After BFS, if any fresh oranges remain they are unreachable, so return `-1`; otherwise return the elapsed minutes.

- Time: O(m * n)
- Space: O(m * n)