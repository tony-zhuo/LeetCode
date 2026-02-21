from collections import Counter


class Solution:
    def containsDuplicate_counter(self, nums: list[int]) -> bool:
        c = Counter(nums)
        return any(x > 1 for x in c.values())

    def containsDuplicate_set(self, nums: list[int]) -> bool:
        return len(nums) != len(set(nums))
