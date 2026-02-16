import math


class SegmentTree:
    """Supports range sum queries with point updates."""
    def __init__(self, nums: list[int]):
        self._n = len(nums)
        self._tree = [0] * (4 * self._n)
        if self._n > 0:
            self._build(nums, 1, 0, self._n - 1)

    def _build(self, nums: list[int], node: int, start: int, end: int) -> None:
        if start == end:
            self._tree[node] = nums[start]
            return
        mid = (start + end) // 2
        self._build(nums, 2 * node, start, mid)
        self._build(nums, 2 * node + 1, mid + 1, end)
        self._tree[node] = self._tree[2 * node] + self._tree[2 * node + 1]

    def update(self, index: int, val: int) -> None:
        self._update(1, 0, self._n - 1, index, val)

    def _update(self, node: int, start: int, end: int, index: int, val: int) -> None:
        if start == end:
            self._tree[node] = val
            return
        mid = (start + end) // 2
        if index <= mid:
            self._update(2 * node, start, mid, index, val)
        else:
            self._update(2 * node + 1, mid + 1, end, index, val)
        self._tree[node] = self._tree[2 * node] + self._tree[2 * node + 1]

    def query(self, left: int, right: int) -> int:
        return self._query(1, 0, self._n - 1, left, right)

    def _query(self, node: int, start: int, end: int, left: int, right: int) -> int:
        if right < start or end < left:
            return 0
        if left <= start and end <= right:
            return self._tree[node]
        mid = (start + end) // 2
        left_sum = self._query(2 * node, start, mid, left, right)
        right_sum = self._query(2 * node + 1, mid + 1, end, left, right)
        return left_sum + right_sum


class RangeMinSegmentTree:
    """Supports range minimum queries with point updates."""
    def __init__(self, nums: list[int]):
        self._n = len(nums)
        self._tree = [math.inf] * (4 * self._n)
        if self._n > 0:
            self._build(nums, 1, 0, self._n - 1)

    def _build(self, nums: list[int], node: int, start: int, end: int) -> None:
        if start == end:
            self._tree[node] = nums[start]
            return
        mid = (start + end) // 2
        self._build(nums, 2 * node, start, mid)
        self._build(nums, 2 * node + 1, mid + 1, end)
        self._tree[node] = min(self._tree[2 * node], self._tree[2 * node + 1])

    def update(self, index: int, val: int) -> None:
        self._update(1, 0, self._n - 1, index, val)

    def _update(self, node: int, start: int, end: int, index: int, val: int) -> None:
        if start == end:
            self._tree[node] = val
            return
        mid = (start + end) // 2
        if index <= mid:
            self._update(2 * node, start, mid, index, val)
        else:
            self._update(2 * node + 1, mid + 1, end, index, val)
        self._tree[node] = min(self._tree[2 * node], self._tree[2 * node + 1])

    def query(self, left: int, right: int) -> int:
        return self._query(1, 0, self._n - 1, left, right)

    def _query(self, node: int, start: int, end: int, left: int, right: int) -> int:
        if right < start or end < left:
            return math.inf
        if left <= start and end <= right:
            return self._tree[node]
        mid = (start + end) // 2
        left_min = self._query(2 * node, start, mid, left, right)
        right_min = self._query(2 * node + 1, mid + 1, end, left, right)
        return min(left_min, right_min)
