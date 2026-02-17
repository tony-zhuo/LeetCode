from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1

        n = len(nums)//2

        if target == nums[n]:
            return n
        elif target > nums[n]:
            res = self.search(nums[n+1:], target)
            if res == -1:
                return -1
            return n+1+res

        else:
            return self.search(nums[:n], target)

    def search_two_pointer(self, nums: List[int], target: int) -> int:
            if not nums:
                return -1

            left, right = 0, len(nums) - 1
            while left <= right:
                n = (left + right) // 2
                if nums[n] == target:
                    return n
                elif target > nums[n]:
                    left = n+1
                else:
                    right = n-1

            return -1