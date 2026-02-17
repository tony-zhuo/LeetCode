import pytest
from solution import Solution


class TestIsPalindrome:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("s,want", [
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True),
        ("", True),
        ("a", True),
        ("ab", False),
        ("aba", True),
        ("0P", False),
    ])
    def test_is_palindrome(self, s, want):
        assert self.sol.isPalindrome(s) == want
