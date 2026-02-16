package problems

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i := 0; i < len(flowerbed); i++ { // 從 0 開始，到 len-1
		if flowerbed[i] == 0 {
			prevEmpty := (i == 0) || (flowerbed[i-1] == 0)
			nextEmpty := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			if prevEmpty && nextEmpty {
				flowerbed[i] = 1
				count++
				if count >= n {
					return true
				}
			}
		}
	}

	return count >= n
}
