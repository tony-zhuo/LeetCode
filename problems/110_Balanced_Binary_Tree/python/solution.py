# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isBalanced(self, root: TreeNode | None) -> bool:
        return self.heigh(root) != -1

    def heigh(self, root: TreeNode | None) -> int:
        if root is None:
            return 0
        
        leftH = self.heigh(root.left)
        rightH = self.heigh(root.right)

        if leftH == -1 or rightH == -1 or abs(leftH-rightH) > 1:
            return -1
        
        return max(leftH, rightH) + 1