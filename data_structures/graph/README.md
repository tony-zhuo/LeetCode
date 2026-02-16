# Graph

## Overview
A graph is a non-linear data structure consisting of vertices (nodes) and edges. This implementation uses an adjacency list representation and supports both directed and undirected graphs.

## Operations & Time Complexity
| Operation | Time |
|-----------|------|
| AddVertex | O(1) |
| AddEdge | O(1) |
| RemoveEdge | O(E) |
| HasEdge | O(E) |
| BFS | O(V + E) |
| DFS | O(V + E) |
| ShortestPath | O(V + E) |

V = number of vertices, E = number of edges per vertex

## Implementation Notes
- Adjacency list using `map[int][]int`
- ShortestPath uses BFS (unweighted graphs only)
- Supports both directed and undirected graphs

## Related LeetCode Problems
- [#133 Clone Graph](../../problems/133_Clone_Graph/)
- [#207 Course Schedule](../../problems/207_Course_Schedule/)
- [#417 Pacific Atlantic Water Flow](../../problems/417_Pacific_Atlantic_Water_Flow/)
