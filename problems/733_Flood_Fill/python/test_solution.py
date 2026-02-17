import copy
import pytest
from solution import Solution

CASES = [
    (
        [[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2,
        [[2, 2, 2], [2, 2, 0], [2, 0, 1]],
    ),
    (
        [[0, 0, 0], [0, 0, 0]], 0, 0, 0,
        [[0, 0, 0], [0, 0, 0]],
    ),
    (
        [[0, 0, 0], [0, 0, 0]], 0, 0, 2,
        [[2, 2, 2], [2, 2, 2]],
    ),
    (
        [[1]], 0, 0, 2,
        [[2]],
    ),
    (
        [[0, 1, 0], [1, 1, 1], [0, 1, 0]], 1, 1, 3,
        [[0, 3, 0], [3, 3, 3], [0, 3, 0]],
    ),
    (
        [[1, 2, 1], [2, 2, 2], [1, 2, 1]], 1, 1, 5,
        [[1, 5, 1], [5, 5, 5], [1, 5, 1]],
    ),
]


class TestFloodFill:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("image,sr,sc,color,want", CASES)
    def test_floodFill(self, image, sr, sc, color, want):
        assert self.sol.floodFill(copy.deepcopy(image), sr, sc, color) == want
