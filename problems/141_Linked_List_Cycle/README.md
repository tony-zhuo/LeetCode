# 141. Linked List Cycle

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

**Example 1:**

> Input: head = [3,2,0,-4], pos = 1  
> Output: true  
> Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

**Example 2:**

> Input: head = [1,2], pos = 0  
> Output: true  
> Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

**Example 3:**

> Input: head = [1], pos = -1  
> Output: false  
> Explanation: There is no cycle in the linked list.

**Constraints:**

- The number of the nodes in the list is in the range `[0, 10^4]`.
- `-10^5 <= Node.val <= 10^5`
- `pos` is `-1` or a **valid index** in the linked-list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Algorithm

This solution uses Floyd's Cycle Detection Algorithm (also known as the "tortoise and hare" algorithm):

1. Use two pointers: `slow` moves one step at a time, `fast` moves two steps at a time
2. If there's a cycle, the fast pointer will eventually meet the slow pointer
3. If there's no cycle, the fast pointer will reach the end (null)

**Time Complexity:** O(n) where n is the number of nodes  
**Space Complexity:** O(1) constant space
