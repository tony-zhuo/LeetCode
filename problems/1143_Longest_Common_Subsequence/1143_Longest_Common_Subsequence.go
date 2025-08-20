package problems

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

/*
### 1. **最優子結構（Optimal Substructure）**
對於 LCS 問題，大問題的最優解可以由子問題的最優解構成：
- 如果我們知道了 `text1[0...i-1]` 和 `text2[0...j-1]` 的 LCS 長度
- 那麼 `text1[0...i]` 和 `text2[0...j]` 的 LCS 可以基於這個結果計算出來

**具體來說**：
- 如果 `text1[i] == text2[j]`，那麼 `LCS(text1[0...i], text2[0...j]) = LCS(text1[0...i-1], text2[0...j-1]) + 1`
- 如果 `text1[i] != text2[j]`，那麼 `LCS(text1[0...i], text2[0...j]) = max(LCS(text1[0...i-1], text2[0...j]), LCS(text1[0...i], text2[0...j-1]))`

### 2. **重疊子問題（Overlapping Subproblems）**
在計算 LCS 時，同樣的子問題會被重複計算多次：

例如計算 LCS("ABC", "AC"):
- 需要計算 LCS("AB", "AC")
- 需要計算 LCS("ABC", "A")
- 而計算這兩個時，都需要計算 LCS("AB", "A")

這就產生了重疊的子問題，如果用遞歸解決會導致大量重複計算。
## 為什麼不能用貪心算法？
LCS 問題不能用貪心算法解決，因為：

text1 = "ABCD"
text2 = "ACBD"

貪心策略可能選擇：A -> C -> D，得到長度 3
但還有其他可能：A -> B -> D，也是長度 3

這說明可能有多個最優解，需要考慮所有可能性

## 遞歸關係的推導

func lcs(text1 string, text2 string, i, j int) int {
    // 基本情況
    if i == 0 || j == 0 {
        return 0
    }

    // 遞歸關係
    if text1[i-1] == text2[j-1] {
        return 1 + lcs(text1, text2, i-1, j-1)
    } else {
        return max(lcs(text1, text2, i-1, j),
                  lcs(text1, text2, i, j-1))
    }
}

這個遞歸解法展示了問題的結構，但會導致指數級的時間複雜度。
## 動態規劃的優勢
通過使用 DP 表格存儲子問題的解：
1. **避免重複計算**：每個子問題只計算一次
2. **時間複雜度優化**：從指數級降到 O(m×n)
3. **空間換時間**：使用 O(m×n) 的空間獲得高效的時間複雜度

## 總結
LCS 問題適合用 DP 解決是因為：
- ✅ 具有最優子結構
- ✅ 存在重疊子問題
- ✅ 可以定義清晰的狀態轉移方程
- ✅ 有明確的邊界條件

*/
