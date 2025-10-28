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

## Complexity Analysis

- **Time Complexity**: O(amount × n) - where n is the number of coins. We compute dp[i] for all amounts from 1 to amount, and for each amount we try all coins.
- **Space Complexity**: O(amount) - for the dp array

## Key Insights

- This is an **unbounded knapsack** problem (unlimited use of each coin)
- The greedy approach (always using the largest coin) doesn't work for all cases
- Dynamic programming ensures we explore all possible combinations
- Similar problems: Climbing Stairs (70), Perfect Squares (279)
