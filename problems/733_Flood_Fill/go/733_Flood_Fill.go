package problems

// Multiple solutions: name each func as {funcName}_{approach}
// e.g. twoSum_bruteforce, twoSum_hashmap
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	targetNum := image[sr][sc]
	if targetNum == color {
		return image
	}

	return fill(image, sr, sc, color, targetNum)
}

func fill(image [][]int, sr int, sc int, color int, targetNum int) [][]int {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) {
		return image
	}

	if targetNum == image[sr][sc] {
		image[sr][sc] = color
		fill(image, sr+1, sc, color, targetNum)
		fill(image, sr-1, sc, color, targetNum)
		fill(image, sr, sc+1, color, targetNum)
		fill(image, sr, sc-1, color, targetNum)
	}

	return image
}
