from collections import Counter


class Solution:
    def majorityElement(self, nums: list[int]) -> int:
        counter = Counter(nums)
        return max(counter, key=lambda k: counter.get(k))

    def majorityElement_BoyerMoore(self, nums: list[int]) -> int:
        candidate = None
        cnt = 0
        for num in nums:
            if cnt == 0:
                candidate = num

            if num == candidate:
                cnt += 1
            else:
                cnt -= 1

        return candidate

