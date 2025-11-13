# 151. Reverse Words in a String

Given an input string `s`, reverse the order of the words.

A **word** is defined as a sequence of non-space characters. The words in `s` will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

**Note** that `s` may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

**Example 1:**

> Input: s = "the sky is blue"
> Output: "blue is sky the"

**Example 2:**

> Input: s = "  hello world  "
> Output: "world hello"
> Explanation: Your reversed string should not contain leading or trailing spaces.

**Example 3:**

> Input: s = "a good   example"
> Output: "example good a"
> Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

**Constraints:**

- `1 <= s.length <= 10^4`
- `s` contains English letters (upper-case and lower-case), digits, and spaces `' '`
- There is at least one word in `s`

**Follow-up:** If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

## 解題思路

此問題可以使用 **字串分割與陣列反轉** 的方法有效解決。

### 演算法步驟:

1. **分割字串**: 使用 `strings.Split(s, " ")` 將字串以空格分割成單字陣列
2. **反向迭代**: 從陣列最後一個元素往第一個元素遍歷
3. **過濾空字串**: 跳過空字串(由多個連續空格產生)
4. **收集有效單字**: 將非空字串加入結果陣列
5. **單一空格合併**: 使用 `strings.Join(result, " ")` 將單字以單一空格連接

### 為什麼這樣做有效:

- **以空格分割** 自然地分離單字,並為多個連續空格產生空字串
- **反向迭代** 以相反順序處理單字
- **過濾空字串** 一次性移除所有前導、尾隨和多餘空格
- **單一空格合併** 確保輸出的正確間距

### 執行過程範例:

讓我們追蹤 **Example 3**: `s = "a good   example"`

**步驟 1: 以空格分割**
```
sArr = ["a", "good", "", "", "example"]
        0     1      2   3      4
```

**步驟 2: 反向迭代並過濾**

迭代 i=4 (sArr[4] = "example"):
- "example" != "" ✓ → 加入結果
- result = ["example"]

迭代 i=3 (sArr[3] = ""):
- "" == "" ✗ → 跳過

迭代 i=2 (sArr[2] = ""):
- "" == "" ✗ → 跳過

迭代 i=1 (sArr[1] = "good"):
- "good" != "" ✓ → 加入結果
- result = ["example", "good"]

迭代 i=0 (sArr[0] = "a"):
- "a" != "" ✓ → 加入結果
- result = ["example", "good", "a"]

**步驟 3: 以單一空格合併**
```
輸出: "example good a"
```

### 其他解法:

1. **兩次遍歷法**: 先移除空格,再分割並反轉
2. **原地解法** (若字串可變): 先反轉整個字串,再反轉每個單字
3. **手動解析**: 使用雙指標提取單字而不分割

## 複雜度分析

- **時間複雜度**: O(n) - 其中 n 為字串長度
  - 分割: O(n)
  - 反向迭代: O(n)
  - 合併: O(n)
  - 整體: O(n)

- **空間複雜度**: O(n) - 用於儲存分割陣列和結果陣列
  - `sArr`: O(n) 空間用於分割後的單字
  - `result`: O(k) 空間,其中 k 為單字數量
  - 整體: O(n)

**注意:** 在 Go 中,字串是不可變的,因此無法達成 O(1) 空間的進階挑戰。具有可變字串的語言(如 C++)可以使用原地字元反轉達成 O(1) 額外空間。

## 關鍵洞察

- **strings.Split** 很方便,但會為連續分隔符號產生空字串
- **迭代時過濾** 比多次遍歷更有效率
- **反向迭代** 消除了反轉結果陣列的需求
- 此演算法自然地處理所有邊界情況:前導/尾隨空格、多個空格
- 類似題目: Reverse String (344)、Reverse String II (541)、Rotate Array (189)
