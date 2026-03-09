# 207. Course Schedule

## Problem

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course `bi` first if you want to take course `ai`.

Return `true` if you can finish all courses. Otherwise, return `false`.

## Solutions

### Go

#### DFS Cycle Detection (`CanFinish`)

Build an adjacency list from prerequisites. For each course, run DFS with a three-state visited map: unvisited (0), in current path (1), and confirmed no cycle (-1). If a node in the current path is revisited, a cycle exists and it's impossible to finish all courses.

- Time: O(V + E)
- Space: O(V + E)
