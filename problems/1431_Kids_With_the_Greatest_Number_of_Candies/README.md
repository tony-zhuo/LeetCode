# 1431. Kids With the Greatest Number of Candies

## 題目描述

有 `n` 個有糖果的孩子。給你一個數組 `candies`，其中 `candies[i]` 代表第 `i` 個孩子擁有的糖果數量，和一個整數 `extraCandies` ，代表你擁有的額外糖果數量。

返回一個長度為 `n` 的布林數組 `result`，其中 `result[i]` 為 `true` 當且僅當把所有額外的糖果給第 `i` 個孩子之後，他將擁有所有孩子中最多的糖果。注意，允許有多個孩子同時擁有最多的糖果。

### 範例

**範例 1:**
```
輸入: candies = [2,3,5,1,3], extraCandies = 3
輸出: [true,true,true,false,true] 
解釋: 
孩子 1 有 2 個糖果，如果他得到所有額外的糖果（3個），那麼他總共有 5 個糖果，這是孩子們之間的最大數量。
孩子 2 有 3 個糖果，如果他得到所有額外的糖果（3個），那麼他總共有 6 個糖果，這是孩子們之間的最大數量。
孩子 3 有 5 個糖果，這已經是孩子們之間的最大數量。
孩子 4 有 1 個糖果，如果他得到所有額外的糖果（3個），那麼他總共有 4 個糖果，這不是孩子們之間的最大數量。
孩子 5 有 3 個糖果，如果他得到所有額外的糖果（3個），那麼他總共有 6 個糖果，這是孩子們之間的最大數量。
```

**範例 2:**
```
輸入: candies = [4,2,1,1,2], extraCandies = 1
輸出: [true,false,false,false,false]
解釋: 只有孩子 1 在得到額外糖果後能擁有最多的糖果。
```

**範例 3:**
```
輸入: candies = [12,1,12], extraCandies = 10
輸出: [true,false,true]
```

### 限制條件

- `n == candies.length`
- `2 <= n <= 100`
- `1 <= candies[i] <= 100`
- `1 <= extraCandies <= 50`

## 解法思路

這是一個簡單的陣列遍歷問題，主要思路如下：

1. **找出最大值**: 首先遍歷 `candies` 陣列，找出當前所有孩子中糖果數量的最大值。

2. **判斷每個孩子**: 對於每個孩子，檢查他的糖果數量加上額外糖果數量是否能達到或超過最大值。

3. **返回結果**: 根據判斷結果構建布林陣列並返回。

### 演算法步驟

```go
func kidsWithCandies(candies []int, extraCandies int) []bool {
    // 步驟 1: 找出最大值
    greatestNum := 0
    for _, candy := range candies {
        greatestNum = max(candy, greatestNum)
    }
    
    // 步驟 2: 初始化結果陣列
    res := make([]bool, len(candies))
    
    // 步驟 3: 檢查每個孩子
    for i, candy := range candies {
        if (candy + extraCandies) >= greatestNum {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    
    return res
}
```

### 關鍵洞察

- 每個孩子得到額外糖果後，只需要與**當前的最大值**比較
- 不需要考慮其他孩子也得到額外糖果的情況，因為題目假設額外糖果只給一個孩子
- 允許多個孩子同時擁有最多糖果（使用 `>=` 而不是 `>`）

## 複雜度分析

### 時間複雜度: O(n)
- 需要遍歷陣列兩次：
  - 第一次遍歷找出最大值：O(n)
  - 第二次遍歷判斷每個孩子：O(n)
- 總時間複雜度為 O(n)

### 空間複雜度: O(1)
- 除了輸出陣列 `result`，只使用了常數額外空間
- 變數 `greatestNum` 和循環變數佔用 O(1) 空間
- 如果不計算輸出空間，空間複雜度為 O(1)
