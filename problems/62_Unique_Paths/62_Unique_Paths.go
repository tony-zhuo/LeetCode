package problems

// uniquePaths calculates the number of unique paths from the top-left to the bottom-right of an m x n grid.
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

/*
## LeetCode 第 62 题：不同路径
### 问题描述
一个机器人位于一个 m × n 网格的左上角（起始点在下图中标记为 "Start"）。机器人每次只能向下或者向右移动一步。
机器人试图达到网格的右下角（在下图中标记为 "Finish"）。问总共有多少条不同的路径？

### 解题思路
这道题可以使用动态规划来解决，主要有以下几种思路：

#### 1. 动态规划（二维数组）
我们可以定义一个二维数组 `dp[i][j]` 表示从起点 (0,0) 到达位置 (i,j) 的不同路径数量。
**状态转移方程**：
- 由于机器人只能向下或向右移动，所以到达位置 (i,j) 的路径只能从 (i-1,j) 或 (i,j-1) 过来
- 因此 `dp[i][j] = dp[i-1][j] + dp[i,j-1]`

**边界条件**：
- 对于第一行，机器人只能向右移动，所以 `dp[0][j] = 1`
- 对于第一列，机器人只能向下移动，所以 `dp[i][0] = 1`

#### 2. 动态规划（一维数组优化）
我们可以观察到，计算 `dp[i][j]` 时只需要用到 `dp[i-1][j]` 和 `dp[i][j-1]`，所以可以使用一维数组优化空间复杂度。

#### 3. 数学组合解法
这个问题本质上可以看作是在 m+n-2 步中选择 m-1 步向下移动（或选择 n-1 步向右移动）的组合数，即 C(m+n-2, m-1) 或 C(m+n-2, n-1)。

### 代码实现
以下是动态规划的 Go 语言实现：

func uniquePaths(m int, n int) int {
    // 创建一个二维dp数组
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    // 初始化第一行和第一列
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }

    // 填充dp数组
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}

### 空间优化版本
以下是使用一维数组优化空间复杂度的实现：

func uniquePaths(m int, n int) int {
    // 创建一个一维dp数组
    dp := make([]int, n)

    // 初始化
    for j := 0; j < n; j++ {
        dp[j] = 1
    }

    // 动态规划过程
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j] + dp[j-1]
        }
    }

    return dp[n-1]
}

### 数学组合解法
func uniquePaths(m int, n int) int {
    // 计算组合数 C(m+n-2, n-1)
    // 为避免溢出，使用较小的数进行组合
    small := min(m-1, n-1)
    large := m + n - 2 - small

    // 计算 C(large+small, small)
    res := 1
    for i := 1; i <= small; i++ {
        res = res * (large + i) / i
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/
