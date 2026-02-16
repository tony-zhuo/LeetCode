# String (StringBuilder)

## Overview
A mutable string builder implemented on a byte slice, providing common string manipulation operations. Since Go strings are immutable, this structure allows efficient in-place modifications.

## Operations & Time Complexity
| Operation | Time | Space |
|-----------|------|-------|
| Append | O(k)* | O(1) |
| InsertAt | O(n+k) | O(n+k) |
| DeleteRange | O(n) | O(1) |
| CharAt | O(1) | O(1) |
| Substring | O(k) | O(k) |
| IndexOf | O(n*m) | O(1) |
| Contains | O(n*m) | O(1) |
| Replace | O(n) | O(n) |
| ReplaceAll | O(n*m) | O(n) |
| Reverse | O(n) | O(1) |
| IsPalindrome | O(n) | O(1) |
| CountChar | O(n) | O(1) |
| Len | O(1) | O(1) |

*k = length of appended string, n = current length, m = target length

## Related LeetCode Problems
- [#3 Longest Substring Without Repeating Characters](../../problems/3_Longest_Substring_Without_Repeating_Characters/)
- [#5 Longest Palindromic Substring](../../problems/5_Longest_Palindromic_Substring/)
- [#242 Valid Anagram](../../problems/242_Valid_Anagram/)
- [#567 Permutation in String](../../problems/567_Permutation_in_String/)
