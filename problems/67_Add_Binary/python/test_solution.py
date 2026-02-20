import pytest
from solution import Solution

CASES = [
    ("11", "1", "100"),
    ("1010", "1011", "10101"),
    ("0", "0", "0"),
    ("1111", "1", "10000"),
    ("1", "111", "1000"),
]


class TestAddBinary:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("a,b,want", CASES)
    def test_addBinary(self, a, b, want):
        assert self.sol.addBinary(a, b) == want
