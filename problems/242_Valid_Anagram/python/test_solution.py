import pytest
from solution import Solution

CASES = [
    ("anagram", "nagaram", True),
    ("rat", "car", False),
    ("", "", True),
    ("a", "a", True),
    ("a", "b", False),
    ("ab", "a", False),
    ("aacc", "ccac", False),
]


class TestIsAnagram:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("s,t,want", CASES)
    def test_hashmap(self, s, t, want):
        assert self.sol.isAnagram_hashmap(s, t) == want

    @pytest.mark.parametrize("s,t,want", CASES)
    def test_array(self, s, t, want):
        assert self.sol.isAnagram_array(s, t) == want

    @pytest.mark.parametrize("s,t,want", CASES)
    def test_sort(self, s, t, want):
        assert self.sol.isAnagram_sort(s, t) == want
