import pytest
from solution import Solution

CASES = [
    ([-1, 0, 3, 5, 9, 12], 9, 4),
    ([-1, 0, 3, 5, 9, 12], 2, -1),
    ([5], 5, 0),
    ([5], -5, -1),
    ([2, 5], 5, 1),
    ([2, 5], 2, 0),
    ([], 0, -1),
    ([1, 2, 3, 4, 5, 6, 7, 8, 9], 1, 0),
    ([1, 2, 3, 4, 5, 6, 7, 8, 9], 9, 8),
]


class TestSearch:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,target,want", CASES)
    def test_search(self, nums, target, want):
        assert self.sol.search(nums, target) == want

    @pytest.mark.parametrize("nums,target,want", CASES)
    def test_two_pointer(self, nums, target, want):
        assert self.sol.search_two_pointer(nums, target) == want