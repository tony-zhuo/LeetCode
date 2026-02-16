class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        match = {
            ')': '(',
            '}': '{',
            ']': '['
        }
        for char in s:
            if char in "({[":
                stack.append(char)
            else:
                if not stack or stack[-1] != match[char]:
                    return False
                stack.pop()

        return not stack