from __future__ import annotations
from collections import deque


class TreeNode:
    def __init__(self, val: int = 0, left: TreeNode | None = None, right: TreeNode | None = None):
        self.val = val
        self.left = left
        self.right = right


def inorder_traversal(root: TreeNode | None) -> list[int]:
    result: list[int] = []
    def _inorder(node: TreeNode | None) -> None:
        if node is None:
            return
        _inorder(node.left)
        result.append(node.val)
        _inorder(node.right)
    _inorder(root)
    return result


def preorder_traversal(root: TreeNode | None) -> list[int]:
    result: list[int] = []
    def _preorder(node: TreeNode | None) -> None:
        if node is None:
            return
        result.append(node.val)
        _preorder(node.left)
        _preorder(node.right)
    _preorder(root)
    return result


def postorder_traversal(root: TreeNode | None) -> list[int]:
    result: list[int] = []
    def _postorder(node: TreeNode | None) -> None:
        if node is None:
            return
        _postorder(node.left)
        _postorder(node.right)
        result.append(node.val)
    _postorder(root)
    return result


def level_order_traversal(root: TreeNode | None) -> list[list[int]]:
    if root is None:
        return []
    result: list[list[int]] = []
    queue: deque[TreeNode] = deque([root])
    while queue:
        level_size = len(queue)
        level: list[int] = []
        for _ in range(level_size):
            node = queue.popleft()
            level.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        result.append(level)
    return result


def height(root: TreeNode | None) -> int:
    if root is None:
        return 0
    return 1 + max(height(root.left), height(root.right))


def count_nodes(root: TreeNode | None) -> int:
    if root is None:
        return 0
    return 1 + count_nodes(root.left) + count_nodes(root.right)


def slice2binary_tree(arr: list[int | None]) -> TreeNode | None:
    if not arr or arr[0] is None:
        return None
    root = TreeNode(arr[0])
    queue: deque[TreeNode] = deque([root])
    i = 1
    while i < len(arr):
        current = queue.popleft()
        if i < len(arr) and arr[i] is not None:
            current.left = TreeNode(arr[i])
            queue.append(current.left)
        i += 1
        if i < len(arr) and arr[i] is not None:
            current.right = TreeNode(arr[i])
            queue.append(current.right)
        i += 1
    return root
