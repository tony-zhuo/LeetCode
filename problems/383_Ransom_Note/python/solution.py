class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        cnt = [0] * 26

        for i in range(len(magazine)):
            cnt[ord(magazine[i])-ord('a')] += 1

        for i in range(len(ransomNote)):
            cnt[ord(ransomNote[i])-ord('a')] -= 1

        return all(x >= 0 for x in cnt)
