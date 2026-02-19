# 70. Climbing Stairs

## Problem

You are climbing a staircase. It takes `n` steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Solutions

### 1. Hashmap DP (`climbStairs_hashmap`)

Use a hash map to store the number of ways for each step.

- Time: O(n)
- Space: O(n)

### 2. Space-Optimized DP (`climbStairs_dp`)

Since `dp[i]` only depends on `dp[i-1]` and `dp[i-2]`, we only need two variables to track the previous two values.

- Time: O(n)
- Space: O(1)

### 3. Matrix Exponentiation (`climbStairs_matrix`)

The Fibonacci recurrence can be expressed as a matrix multiplication:

```
| F(n+1) |   | 1  1 |^n   | 1 |
| F(n)   | = | 1  0 |   * | 0 |
```

By using fast matrix exponentiation (repeated squaring), we can compute the result in O(log n) time.

- Time: O(log n)
- Space: O(1)
