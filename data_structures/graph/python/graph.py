from __future__ import annotations
from collections import deque


class Node:
    """Graph node used by problem solutions (adjacency via pointers)."""
    def __init__(self, val: int = 0, neighbors: list[Node] | None = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


class Graph:
    """Graph using adjacency list representation."""
    def __init__(self, directed: bool = False):
        self._adjacency: dict[int, list[int]] = {}
        self._directed = directed

    def add_vertex(self, v: int) -> None:
        if v not in self._adjacency:
            self._adjacency[v] = []

    def add_edge(self, from_v: int, to_v: int) -> None:
        self.add_vertex(from_v)
        self.add_vertex(to_v)
        self._adjacency[from_v].append(to_v)
        if not self._directed:
            self._adjacency[to_v].append(from_v)

    def remove_edge(self, from_v: int, to_v: int) -> None:
        if from_v in self._adjacency:
            try:
                self._adjacency[from_v].remove(to_v)
            except ValueError:
                pass
        if not self._directed and to_v in self._adjacency:
            try:
                self._adjacency[to_v].remove(from_v)
            except ValueError:
                pass

    def has_edge(self, from_v: int, to_v: int) -> bool:
        return from_v in self._adjacency and to_v in self._adjacency[from_v]

    def neighbors(self, v: int) -> list[int]:
        return self._adjacency.get(v, [])

    def vertices(self) -> list[int]:
        return sorted(self._adjacency.keys())

    def bfs(self, start: int) -> list[int]:
        if start not in self._adjacency:
            return []
        visited: set[int] = {start}
        order: list[int] = []
        queue: deque[int] = deque([start])
        while queue:
            current = queue.popleft()
            order.append(current)
            for neighbor in sorted(self._adjacency[current]):
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)
        return order

    def dfs(self, start: int) -> list[int]:
        if start not in self._adjacency:
            return []
        visited: set[int] = set()
        order: list[int] = []
        self._dfs_helper(start, visited, order)
        return order

    def _dfs_helper(self, v: int, visited: set[int], order: list[int]) -> None:
        visited.add(v)
        order.append(v)
        for neighbor in sorted(self._adjacency[v]):
            if neighbor not in visited:
                self._dfs_helper(neighbor, visited, order)

    def shortest_path(self, start: int, end: int) -> tuple[list[int] | None, int]:
        if start not in self._adjacency or end not in self._adjacency:
            return None, -1
        if start == end:
            return [start], 0
        visited: set[int] = {start}
        parent: dict[int, int] = {}
        queue: deque[int] = deque([start])
        while queue:
            current = queue.popleft()
            for neighbor in self._adjacency[current]:
                if neighbor not in visited:
                    visited.add(neighbor)
                    parent[neighbor] = current
                    if neighbor == end:
                        path = self._reconstruct_path(parent, start, end)
                        return path, len(path) - 1
                    queue.append(neighbor)
        return None, -1

    def _reconstruct_path(self, parent: dict[int, int], start: int, end: int) -> list[int]:
        path = []
        current = end
        while current != start:
            path.append(current)
            current = parent[current]
        path.append(start)
        path.reverse()
        return path
