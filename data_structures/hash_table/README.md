# Hash Table

## Overview
A hash table is a data structure that maps keys to values using a hash function. This implementation uses separate chaining (linked lists) for collision resolution.

## Operations & Time Complexity
| Operation | Average | Worst |
|-----------|---------|-------|
| Put | O(1) | O(n) |
| Get | O(1) | O(n) |
| Delete | O(1) | O(n) |
| Contains | O(1) | O(n) |
| Keys | O(n) | O(n) |

## Implementation Notes
- Uses separate chaining for collision resolution
- Hash function: sum of byte values modulo capacity
- String keys, integer values

## Related LeetCode Problems
- [#49 Group Anagrams](../../problems/49_Group_Anagrams/)
- [#128 Longest Consecutive Sequence](../../problems/128_Longest_Consecutive_Sequence/)
- [#242 Valid Anagram](../../problems/242_Valid_Anagram/)
- [#347 Top K Frequent Elements](../../problems/347_Top_K_Frequent_Elements/)
