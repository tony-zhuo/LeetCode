import pytest
from solution import Solution

CASES = [
    ([4, 2], 6),
    ([1], -1),
    ([3, 1, 4], 2),
    ([1, 2, 3], 0),
    ([2, 2, 2, 2, 2, 2], 8),
    ([2, 2, 2, 2, 2, 3], 13),
]


class TestScoreDifference:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,want", CASES)
    def test_scoreDifference(self, nums, want):
        assert self.sol.scoreDifference(nums[:]) == want
