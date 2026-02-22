import pytest
from solution import Solution

CASES = [
    ("1100", "1010", "1111"),
    ("0000", "0000", "0000"),
    ("1111", "1111", "0000"),
    ("10", "10", "11"),
    ("110", "110", "101"),
    ("1", "0", "1"),
]


class TestMaximumXor:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("s,t,want", CASES)
    def test_maximumXor(self, s, t, want):
        assert self.sol.maximumXor(s, t) == want
