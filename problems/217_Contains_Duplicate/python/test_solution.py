import pytest
from solution import Solution

CASES = [
    ([1, 2, 3, 1], True),
    ([1, 2, 3, 4], False),
    ([1, 1, 1, 3, 3, 4, 3, 2, 4, 2], True),
    ([1], False),
    ([1, 1], True),
]


class TestContainsDuplicate:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("nums,want", CASES)
    def test_counter(self, nums, want):
        assert self.sol.containsDuplicate_counter(nums[:]) == want

    @pytest.mark.parametrize("nums,want", CASES)
    def test_set(self, nums, want):
        assert self.sol.containsDuplicate_set(nums[:]) == want
