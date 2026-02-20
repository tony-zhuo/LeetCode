# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def __init__(self):
        self.ans = 0

    def diameterOfBinaryTree(self, root: TreeNode | None) -> int:
        self.depth(root)
        return self.ans

    def depth(self, node: TreeNode | None) -> int:
        if not node:
            return 0

        left = self.depth(node.left)
        right = self.depth(node.right)
        self.ans = max(self.ans, left + right)
        return max(left, right) + 1