import pytest
from solution import Solution

CASES = [
    ([3, 2, 3], 3),
    ([2, 2, 1, 1, 1, 2, 2], 2),
    ([1], 1),
    ([6, 5, 5], 5),
]


class TestMajorityElement:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,want", CASES)
    def test_counter(self, nums, want):
        assert self.sol.majorityElement(nums[:]) == want

    @pytest.mark.parametrize("nums,want", CASES)
    def test_boyer_moore(self, nums, want):
        assert self.sol.majorityElement_BoyerMoore(nums[:]) == want
