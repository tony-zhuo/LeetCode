# 122. Best Time to Buy and Sell Stock II

## Problem Description

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

On each day, you may decide to buy and/or sell the stock. You can only hold **at most one** share of the stock at any time. However, you can buy it then immediately sell it on the **same day**.

Find and return the **maximum profit** you can achieve.

### Examples

**Example 1:**
```
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
```

**Example 2:**
```
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.
```

**Example 3:**
```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
```

### Constraints

- `1 <= prices.length <= 3 * 10^4`
- `0 <= prices[i] <= 10^4`

## Solution

The key insight is that we can make as many transactions as we want. This means we can capture **every profitable price increase** by buying before any price rise and selling immediately after.

### Algorithm

The solution uses a greedy approach:
1. Iterate through the prices array
2. For each day, if the price is higher than the previous day, add the difference to profit
3. This is equivalent to buying yesterday and selling today for every profitable opportunity

### Example Walkthrough

For `prices = [7,1,5,3,6,4]`:
- Day 0→1: 7→1 (decrease, no action)
- Day 1→2: 1→5 (increase by 4, profit += 4)
- Day 2→3: 5→3 (decrease, no action)
- Day 3→4: 3→6 (increase by 3, profit += 3)
- Day 4→5: 6→4 (decrease, no action)

Total profit = 4 + 3 = 7

### Alternative Implementation

The current implementation uses a two-pointer approach:
- `left` pointer tracks the buying price
- `right` pointer scans forward
- When `prices[right] > prices[left]`, we capture the profit and reset

Both approaches achieve the same result of capturing all profitable increases.

### Time Complexity
- **Time:** O(n) where n is the length of prices array
- **Space:** O(1) - only using constant extra space

### Key Differences from Stock I
- **Stock I (121)**: Only one transaction allowed → find maximum single profit
- **Stock II (122)**: Unlimited transactions → capture all profitable increases

### Greedy Strategy
Since we can make unlimited transactions, we can be greedy and take profit from every price increase. This optimal strategy works because:
1. We can buy and sell on the same day
2. Multiple small profits can be better than waiting for one large profit
3. Missing any profitable increase is suboptimal
