import pytest
from solution import Solution

CASES = [
    ("a", "b", False),
    ("aa", "ab", False),
    ("aa", "aab", True),
    ("", "abc", True),
    ("", "", True),
    ("abc", "abc", True),
    ("aab", "aba", True),
    ("abcd", "abc", False),
]


class TestCanConstruct:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("ransomNote,magazine,want", CASES)
    def test_canConstruct(self, ransomNote, magazine, want):
        assert self.sol.canConstruct(ransomNote, magazine) == want
