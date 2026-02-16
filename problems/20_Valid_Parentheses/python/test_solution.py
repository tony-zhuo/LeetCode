import pytest
from solution import Solution


class TestIsValid:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("s,want", [
        ("()", True),
        ("()[]{}", True),
        ("(]", False),
        ("([)]", False),
        ("{[]}", True),
        ("", True),
        ("(", False),
        (")", False),
    ])
    def test_is_valid(self, s, want):
        assert self.sol.isValid(s) == want
