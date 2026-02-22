from math import factorial


class Solution:
    def isDigitorialPermutation(self, n: int) -> bool:
        digiSum = sum(factorial(int(d)) for d in str(n))
        return sorted(str(digiSum)) == sorted(str(n))
