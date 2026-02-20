class Solution:
    def addBinary(self, a: str, b: str) -> str:
        carry = 0
        result = ""
        i = len(a) - 1
        j = len(b) - 1

        while i >= 0 or j >= 0 or carry > 0:
            total = carry
            if i >= 0:
                total += int(a[i])
                i -= 1

            if j >= 0:
                total += int(b[j])
                j -= 1

            result = str(total % 2) + result
            carry = total // 2

        return result