from collections import Counter


class Solution:
    def longestPalindrome(self, s: str) -> int:
        freq = [0]*128
        for c in s:
            freq[ord(c)] += 1

        res = 0
        hasOdd = False

        for c in freq:
            res += c //2 * 2
            if c%2 == 1:
                hasOdd = True

        if hasOdd:
            res += 1

        return res

    def longestPalindrome_counter(self, s: str) -> int:
        counts = Counter(s)

        res = 0
        hasOdd = False

        for c in counts.values():
            res += c //2 * 2
            if c%2 == 1:
                hasOdd = True

        if hasOdd:
            res += 1

        return res