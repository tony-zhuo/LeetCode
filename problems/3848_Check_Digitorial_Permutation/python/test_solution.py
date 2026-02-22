import pytest
from solution import Solution

CASES = [
    (1, True),
    (2, True),
    (145, True),
    (514, True),
    (3, False),
    (10, False),
]


class TestIsDigitorialPermutation:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("n,want", CASES)
    def test_isDigitorialPermutation(self, n, want):
        assert self.sol.isDigitorialPermutation(n) == want
