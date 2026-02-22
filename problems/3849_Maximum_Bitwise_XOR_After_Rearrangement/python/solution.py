class Solution:
    # Multiple solutions: name each method as {methodName}_{approach}
    # e.g. twoSum_bruteforce, twoSum_hashmap
    def maximumXor(self, s: str, t: str) -> str:
        # t can rearrange
        # s cant rearrange
        # t and s XOR to max int
        ones = t.count('1')
        zeros = len(t) - ones
        result = []

        for num in s:
            if num == '1':
                if zeros > 0:
                    zeros -= 1
                    result.append('1')
                else:
                    ones -= 1
                    result.append('0')
            else:
                if ones > 0:
                    ones -= 1
                    result.append('1')
                else:
                    zeros -= 1
                    result.append('0')


        return ''.join(result)
