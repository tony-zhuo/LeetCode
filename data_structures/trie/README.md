# Trie (Prefix Tree)

## Overview
A trie is a tree-like data structure used to store strings, where each node represents a character. It enables efficient prefix-based search and autocomplete operations.

## Operations & Time Complexity
| Operation | Time | Space |
|-----------|------|-------|
| Insert | O(m) | O(m) |
| Search | O(m) | O(1) |
| StartsWith | O(m) | O(1) |
| Delete | O(m) | O(1) |

m = length of the word/prefix

## Implementation Notes
- Uses a fixed-size array of 26 children (lowercase English letters only)
- Each node tracks whether it marks the end of a complete word

## Related LeetCode Problems
- [#208 Implement Trie (Prefix Tree)](../../problems/208_Implement_Trie_(Prefix_Tree)/)
- [#211 Design Add and Search Words Data Structure](../../problems/211_Design_Add_and_Search_Words_Data_Structure/)
