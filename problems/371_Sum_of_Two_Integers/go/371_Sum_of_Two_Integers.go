package problems

func getSum(a int, b int) int {
	// 當 b 等於 0 時，表示沒有進位了，計算完成
	for b != 0 {
		// 計算進位
		carry := a & b

		// 計算無進位的加法結果
		a = a ^ b

		// 進位要左移一位
		b = carry << 1
	}

	return a

}
