# 1071. Greatest Common Divisor of Strings

## 題目描述

For two strings `s` and `t`, we say "`t` divides `s`" if and only if `s = t + t + ... + t` (i.e., `t` is concatenated with itself one or more times).

Given two strings `str1` and `str2`, return the largest string `x` such that `x` divides both `str1` and `str2`.

### 範例

**Example 1:**
```
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
```

**Example 2:**
```
Input: str1 = "ABABAB", str2 = "ABAB"  
Output: "AB"
```

**Example 3:**
```
Input: str1 = "LEET", str2 = "CODE"
Output: ""
```

### 限制條件

- `1 <= str1.length, str2.length <= 1000`
- `str1` and `str2` consist of English uppercase letters.

## 解法思路

這道題目要找出兩個字串的最大公共除數字串。關鍵洞察是：

1. **必要條件檢查**: 如果兩個字串有公共除數，那麼 `str1 + str2` 必須等於 `str2 + str1`
   - 例如: `"ABCABC" + "ABC" = "ABCABCABC"` 和 `"ABC" + "ABCABC" = "ABCABCABC"`
   - 反例: `"LEET" + "CODE" = "LEETCODE"` 但 `"CODE" + "LEET" = "CODELEET"`

2. **長度關係**: 如果存在公共除數字串，其長度必須是兩個原字串長度的最大公約數 (GCD)
   - 這是因為除數字串必須能整除兩個字串

3. **算法步驟**:
   - 檢查 `str1 + str2 == str2 + str1`，如果不相等則沒有公共除數
   - 計算兩個字串長度的 GCD
   - 返回前 GCD 長度的子字串

## 程式碼實現

```go
func gcdOfStrings(str1 string, str2 string) string {
    len1, len2 := len(str1), len(str2)
    if str1+str2 != str2+str1 {
        return ""
    }
    
    gcdLen := gcd(len1, len2)
    return str1[:gcdLen]
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}
```

## 複雜度分析

### 時間複雜度: O(m + n)
- `m` 和 `n` 分別是 `str1` 和 `str2` 的長度
- 字串連接和比較: O(m + n)
- GCD 計算: O(log(min(m, n)))
- 總體時間複雜度為 O(m + n)

### 空間複雜度: O(m + n)
- 需要額外空間來存儲連接後的字串進行比較
- 遞迴調用 GCD 的棧空間: O(log(min(m, n)))
- 總體空間複雜度為 O(m + n)

## 關鍵洞察

1. **字串除法的本質**: 如果字串 X 能除字串 Y，那麼 Y 必須是 X 的若干次重複
2. **交換律測試**: 利用字串連接的交換律來快速判斷是否存在公共除數
3. **數學關聯**: 字串的最大公共除數長度等於字串長度的最大公約數

這個解法巧妙地將字串問題轉化為數學問題，通過簡單的檢查和 GCD 計算就能得到答案。
