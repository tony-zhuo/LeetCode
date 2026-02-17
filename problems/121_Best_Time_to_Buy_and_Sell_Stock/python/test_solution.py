import pytest
from solution import Solution


class TestMaxProfit:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("prices,want", [
        ([7, 1, 5, 3, 6, 4], 5),
        ([7, 6, 4, 3, 1], 0),
        ([1, 2], 1),
        ([2, 1], 0),
        ([1], 0),
        ([3, 3, 3], 0),
        ([2, 4, 1, 7], 6),
    ])
    def test_max_profit(self, prices, want):
        assert self.sol.maxProfit(prices) == want
