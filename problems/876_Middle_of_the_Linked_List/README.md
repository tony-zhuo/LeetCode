# 876. Middle of the Linked List

## Problem

Given the `head` of a singly linked list, return the **middle node** of the linked list.

If there are two middle nodes, return the **second middle** node.

## Solutions

### Go

#### 1. Two Pass (`middleNode_two_pass`)

First pass counts the total number of nodes. Second pass advances to the `count / 2` position.

- Time: O(n)
- Space: O(1)

#### 2. Slow & Fast Pointers (`middleNode_slow_fast`)

Use two pointers: slow moves one step at a time, fast moves two steps. When fast reaches the end, slow is at the middle.

- Time: O(n)
- Space: O(1)

### Python

#### Slow & Fast Pointers (`middleNode`)

Same two-pointer approach — slow advances one node, fast advances two. When fast can no longer move, slow points to the middle.

- Time: O(n)
- Space: O(1)
