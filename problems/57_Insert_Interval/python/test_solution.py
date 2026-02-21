import pytest
from solution import Solution

CASES = [
    ([[1, 3], [6, 9]], [2, 5], [[1, 5], [6, 9]]),
    ([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8], [[1, 2], [3, 10], [12, 16]]),
    ([], [5, 7], [[5, 7]]),
    ([[3, 5], [6, 9]], [1, 2], [[1, 2], [3, 5], [6, 9]]),
    ([[1, 2], [3, 5]], [6, 8], [[1, 2], [3, 5], [6, 8]]),
]


class TestInsert:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("intervals,newInterval,want", CASES)
    def test_insert(self, intervals, newInterval, want):
        assert self.sol.insert([i[:] for i in intervals], newInterval[:]) == want
