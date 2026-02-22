# 3848. Check Digitorial Permutation

## Problem

Given an integer `n`, compute the **digitorial** — the sum of the factorial of each digit. Return `true` if the digitorial is a **permutation** of `n` (i.e. they contain the same digits in any order).

For example, `n = 145`: `1! + 4! + 5! = 1 + 24 + 120 = 145`, which is a permutation of itself → `true`.
For `n = 514`: `5! + 1! + 4! = 120 + 1 + 24 = 145`, and `"145"` is a permutation of `"514"` → `true`.

## Solutions

### Go

#### Sort & Compare (`isDigitorialPermutation`)

Compute the digit factorial sum, convert both the sum and `n` to strings, sort the characters, and compare.

- Time: O(d log d) where d is the number of digits
- Space: O(d)

### Python

#### Sort & Compare (`isDigitorialPermutation`)

Same approach using `math.factorial` and `sorted()` on the string representations.

- Time: O(d log d) where d is the number of digits
- Space: O(d)
