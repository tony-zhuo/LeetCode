# Data Structures Practice

From-scratch implementations of common data structures in Go, organized by category.

## Usage

```bash
# Run all data structure tests
make test_ds

# Run tests for a specific data structure
make test_ds_single

# Scaffold a new data structure directory
make new_ds
```

## Overview

### Basic Structures

| Data Structure | Directory | Key Operations | Related LeetCode |
|---------------|-----------|---------------|-----------------|
| Array | [array](array/) | Get, Set, Append, InsertAt, DeleteAt, Contains, Reverse | [#1](../problems/1_Two_Sum/), [#53](../problems/53_Maximum_Subarray/), [#217](../problems/217_Contains_Duplicate/) |
| String | [string](string/) | Append, InsertAt, DeleteRange, IndexOf, Replace, Reverse, IsPalindrome | [#3](../problems/3_Longest_Substring_Without_Repeating_Characters/), [#5](../problems/5_Longest_Palindromic_Substring/), [#242](../problems/242_Valid_Anagram/), [#567](../problems/567_Permutation_in_String/) |
| Linked List | [linked_list](linked_list/) | InsertAtHead/Tail, Delete, Search, Reverse | [#141](../problems/141_Linked_List_Cycle/), [#206](../problems/206_Reverse_Linked_List/) |
| Stack | [stack](stack/) | Push, Pop, Peek, IsEmpty, Size | — |
| Queue | [queue](queue/) | Enqueue, Dequeue, Peek + CircularQueue | — |
| Hash Table | [hash_table](hash_table/) | Put, Get, Delete, Contains, Keys | [#49](../problems/49_Group_Anagrams/), [#128](../problems/128_Longest_Consecutive_Sequence/), [#242](../problems/242_Valid_Anagram/), [#347](../problems/347_Top_K_Frequent_Elements/) |

### Tree Structures

| Data Structure | Directory | Key Operations | Related LeetCode |
|---------------|-----------|---------------|-----------------|
| Binary Tree | [binary_tree](binary_tree/) | Inorder/Preorder/Postorder/LevelOrder, Height, Count | [#100](../problems/100_Same_Tree/), [#102](../problems/102_Binary_Tree_Level_Order_Traversal/), [#104](../problems/104_Maximum_Depth_of_Binary_Tree/), [#105](../problems/105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal/), [#572](../problems/572_Subtree_of_Another_Tree/) |
| BST | [bst](bst/) | Insert, Delete, Search, Min, Max, InorderSuccessor | [#98](../problems/98_Validate_Binary_Search_Tree/), [#230](../problems/230_Kth_Smallest_Element_in_a_BST/) |
| Heap | [heap](heap/) | Insert, Extract, Peek, Heapify (MinHeap + MaxHeap) | [#253](../problems/253_Meeting_Rooms_II/), [#347](../problems/347_Top_K_Frequent_Elements/) |
| Trie | [trie](trie/) | Insert, Search, StartsWith, Delete | [#208](../problems/208_Implement_Trie_(Prefix_Tree)/), [#211](../problems/211_Design_Add_and_Search_Words_Data_Structure/) |

### Graph & Advanced

| Data Structure | Directory | Key Operations | Related LeetCode |
|---------------|-----------|---------------|-----------------|
| Graph | [graph](graph/) | AddVertex, AddEdge, BFS, DFS, ShortestPath | [#133](../problems/133_Clone_Graph/), [#207](../problems/207_Course_Schedule/), [#417](../problems/417_Pacific_Atlantic_Water_Flow/) |
| Union-Find | [union_find](union_find/) | Find, Union, Connected, Count | — |
| Segment Tree | [segment_tree](segment_tree/) | Build, Update, Query (range sum/min) | — |
