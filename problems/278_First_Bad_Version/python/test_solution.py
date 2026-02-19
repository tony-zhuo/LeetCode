import pytest
import solution as sol_module
from solution import Solution

CASES = [
    (5, 4),
    (1, 1),
    (10, 1),
    (10, 10),
    (100, 50),
    (2, 2),
]


class TestFirstBadVersion:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("n,bad", CASES)
    def test_firstBadVersion(self, n, bad):
        sol_module.isBadVersion = lambda version: version >= bad
        assert self.sol.firstBadVersion(n) == bad
