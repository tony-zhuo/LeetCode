class Solution:
    def scoreDifference(self, nums: list[int]) -> int:
        firstPlayer = 0
        secondPlayer = 0
        currentPlayer = 1
        for game, num in enumerate(nums):
            if num % 2 != 0:
                currentPlayer = 3 - currentPlayer
            if (game + 1) % 6 == 0:
                currentPlayer = 3 - currentPlayer
            if currentPlayer == 1:
                firstPlayer += num
            else:
                secondPlayer += num

        return firstPlayer - secondPlayer