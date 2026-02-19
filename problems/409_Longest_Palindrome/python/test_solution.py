import pytest
from solution import Solution

CASES = [
    ("abccccdd", 7),
    ("a", 1),
    ("aa", 2),
    ("aaa", 3),
    ("abcba", 5),
    ("abcde", 1),
    ("AaBb", 1),
    ("ccc", 3),
    ("aaabbbccc", 7),
]


class TestLongestPalindrome:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("s,want", CASES)
    def test_array(self, s, want):
        assert self.sol.longestPalindrome(s) == want

    @pytest.mark.parametrize("s,want", CASES)
    def test_counter(self, s, want):
        assert self.sol.longestPalindrome_counter(s) == want
