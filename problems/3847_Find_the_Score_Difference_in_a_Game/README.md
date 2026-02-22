# 3847. Find the Score Difference in a Game

## Problem

Two players play a game with an array `nums`. Player 1 starts. For each round, the current player adds the number to their score. The active player switches when:

1. The current number is **odd**
2. Every **6th** round

Return the difference `player1_score - player2_score`.

## Solutions

### Go

#### Simulation (`scoreDifference`)

Iterate through nums, tracking the current player. Toggle on odd numbers and every 6th round using `current = 3 - current`. Accumulate scores and return the difference.

- Time: O(n)
- Space: O(1)

### Python

#### Simulation (`scoreDifference`)

Same toggle approach with `enumerate` to track the round number.

- Time: O(n)
- Space: O(1)
