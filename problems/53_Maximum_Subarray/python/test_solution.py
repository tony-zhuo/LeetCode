import pytest
from solution import Solution

CASES = [
    ([-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
    ([1], 1),
    ([5, 4, -1, 7, 8], 23),
    ([-3, -2, -5, -1], -1),
    ([-1, 2], 2),
]


class TestMaxSubArray:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,want", CASES)
    def test_maxSubArray(self, nums, want):
        assert self.sol.maxSubArray(nums[:]) == want
