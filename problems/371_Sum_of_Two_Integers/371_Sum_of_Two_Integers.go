package problems

import "fmt"

func getSum(a int, b int) int {
	// 1
	// 32 16 8 4 2 1
	//  0  0 0 0 0 1

	// 2
	// 32 16 8 4 2 1
	//  0  0 0 0 1 0

	// 1 + 2
	//   0  0 0 0 0 1
	//+  0  0 0 0 1 0
	// --------------
	//.  0  0 0 0 1 1

	// 當 b 等於 0 時，表示沒有進位了，計算完成
	for b != 0 {
		// 計算進位
		carry := a & b
		fmt.Printf("carry: %b\n", carry)

		// 計算無進位的加法結果
		a = a ^ b
		fmt.Printf("carry: %b\n", a)

		// 進位要左移一位
		b = carry << 1
	}

	return a

}
