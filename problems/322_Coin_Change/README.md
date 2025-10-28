# 322. Coin Change

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

You may assume that you have an infinite number of each kind of coin.

**Example 1:**

> Input: coins = [1, 2, 5], amount = 11
> Output: 3
> Explanation: 11 = 5 + 5 + 1

**Example 2:**

> Input: coins = [2], amount = 3
> Output: -1

**Example 3:**

> Input: coins = [1], amount = 0
> Output: 0

**Constraints:**

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## Solution Approach

This is a classic **dynamic programming** problem that can be solved using the **unbounded knapsack** pattern.

### Algorithm Steps:

1. **Define DP state**: Let `dp[i]` represent the minimum number of coins needed to make amount `i`
2. **Base case**: `dp[0] = 0` (zero coins needed for amount 0)
3. **Initialization**: Set all other `dp[i]` to infinity (or `math.MaxInt32`)
4. **State transition**: For each amount `i` from 1 to `amount`:
   - Try each coin denomination
   - If `i >= coin`, then `dp[i] = min(dp[i], dp[i-coin] + 1)`
5. **Return result**: If `dp[amount]` is still infinity, return -1; otherwise return `dp[amount]`

### Why this works:

For any amount `i`, we can build it from a smaller amount `i - coin` by adding one more coin. By trying all possible coins and taking the minimum, we ensure we find the optimal solution.

### Execution Example:

Let's trace through **Example 1**: `coins = [1, 2, 5], amount = 11`

**Initialization:**
```
dp = [0, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞, ∞]
      0  1  2  3  4  5  6  7  8  9 10 11
```

**Step-by-step DP table construction:**

For amount = 1:
- Try coin 1: dp[1] = min(∞, dp[0] + 1) = 1
- Result: `dp[1] = 1` (one 1-coin)

For amount = 2:
- Try coin 1: dp[2] = min(∞, dp[1] + 1) = 2
- Try coin 2: dp[2] = min(2, dp[0] + 1) = 1
- Result: `dp[2] = 1` (one 2-coin)

For amount = 3:
- Try coin 1: dp[3] = min(∞, dp[2] + 1) = 2
- Try coin 2: dp[3] = min(2, dp[1] + 1) = 2
- Result: `dp[3] = 2` (1+2 or 2+1)

For amount = 5:
- Try coin 1: dp[5] = min(∞, dp[4] + 1) = 3
- Try coin 2: dp[5] = min(3, dp[3] + 1) = 3
- Try coin 5: dp[5] = min(3, dp[0] + 1) = 1
- Result: `dp[5] = 1` (one 5-coin)

...continuing this process...

For amount = 11:
- Try coin 1: dp[11] = min(∞, dp[10] + 1) = 3
- Try coin 2: dp[11] = min(3, dp[9] + 1) = 3
- Try coin 5: dp[11] = min(3, dp[6] + 1) = 3
- Result: `dp[11] = 3` (5+5+1)

**Final DP table:**
```
dp = [0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 2, 3]
      0  1  2  3  4  5  6  7  8  9 10 11
```

The answer is `dp[11] = 3`, which represents the combination 5 + 5 + 1.

## Complexity Analysis

- **Time Complexity**: O(amount × n) - where n is the number of coins. We compute dp[i] for all amounts from 1 to amount, and for each amount we try all coins.
- **Space Complexity**: O(amount) - for the dp array

## Key Insights

- This is an **unbounded knapsack** problem (unlimited use of each coin)
- The greedy approach (always using the largest coin) doesn't work for all cases
- Dynamic programming ensures we explore all possible combinations
- Similar problems: Climbing Stairs (70), Perfect Squares (279)
