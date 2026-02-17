from typing import List


class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        targetNum = image[sr][sc]

        if targetNum == color:
            return image

        return self.fill(image, sr, sc, color, targetNum)


    def fill(self, image: List[List[int]], sr: int, sc: int, color: int, targetNum: int) -> List[List[int]]:
        if sr < 0 or sr >= len(image) or sc < 0 or sc >= len(image[sr]):
            return image

        if targetNum == image[sr][sc]:
            image[sr][sc] = color
            self.fill(image, sr+1, sc, color, targetNum)
            self.fill(image, sr-1, sc, color, targetNum)
            self.fill(image, sr, sc+1, color, targetNum)
            self.fill(image, sr, sc-1, color, targetNum)

        return image