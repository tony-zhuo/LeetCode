from solution import Solution

import pytest


class TestTwoSum:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,target,want", [
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
    ])
    def test_two_sum(self, nums, target, want):
        assert self.sol.twoSum(nums, target) == want
