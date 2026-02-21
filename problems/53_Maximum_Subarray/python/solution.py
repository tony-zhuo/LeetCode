class Solution:
    def maxSubArray(self, nums: list[int]) -> int:
        currentSum = nums[0]
        globalMax = nums[0]

        for i in range(1, len(nums)):
            currentSum = max(nums[i], currentSum + nums[i])
            globalMax = max(currentSum, globalMax)

        return globalMax
