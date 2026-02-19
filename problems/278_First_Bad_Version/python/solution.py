# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

# Placeholder so the module can be imported in tests.
def isBadVersion(version: int) -> bool:
    raise NotImplementedError


class Solution:
    def firstBadVersion(self, n: int) -> int:
        left, right = 1, n
        while left < right:
            mid = (left + right) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid + 1

        return left
