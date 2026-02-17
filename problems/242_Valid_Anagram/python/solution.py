class Solution:
    def isAnagram_hashmap(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        strMap = {}

        for sItem in s:
            strMap[sItem] = strMap.get(sItem, 0) + 1

        for tItem in t:
            v = strMap.get(tItem, 0)
            if v == 0:
                return False
            strMap[tItem] -= 1

        return True

    def isAnagram_array(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        count = [0] * 26

        for i in range(len(s)):
            count[ord(s[i]) - ord('a')] += 1
            count[ord(t[i]) - ord('a')] -= 1

        return all(c == 0 for c in count)

    def isAnagram_sort(self, s: str, t: str) -> bool:
        return sorted(s) == sorted(t)