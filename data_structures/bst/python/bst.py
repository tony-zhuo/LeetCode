from __future__ import annotations


class BSTNode:
    def __init__(self, val: int, left: BSTNode | None = None, right: BSTNode | None = None):
        self.val = val
        self.left = left
        self.right = right


class BST:
    def __init__(self):
        self.root: BSTNode | None = None
        self._size = 0

    def insert(self, val: int) -> None:
        self.root, inserted = self._insert(self.root, val)
        if inserted:
            self._size += 1

    def _insert(self, node: BSTNode | None, val: int) -> tuple[BSTNode, bool]:
        if node is None:
            return BSTNode(val), True
        if val < node.val:
            node.left, inserted = self._insert(node.left, val)
            return node, inserted
        elif val > node.val:
            node.right, inserted = self._insert(node.right, val)
            return node, inserted
        return node, False  # duplicate

    def delete(self, val: int) -> bool:
        self.root, deleted = self._delete(self.root, val)
        if deleted:
            self._size -= 1
        return deleted

    def _delete(self, node: BSTNode | None, val: int) -> tuple[BSTNode | None, bool]:
        if node is None:
            return None, False
        if val < node.val:
            node.left, deleted = self._delete(node.left, val)
            return node, deleted
        elif val > node.val:
            node.right, deleted = self._delete(node.right, val)
            return node, deleted
        else:
            # Found the node to delete
            if node.left is None and node.right is None:
                return None, True
            if node.left is None:
                return node.right, True
            if node.right is None:
                return node.left, True
            # Two children: use in-order successor
            successor = self._find_min(node.right)
            node.val = successor.val
            node.right, _ = self._delete(node.right, successor.val)
            return node, True

    def search(self, val: int) -> bool:
        return self._search(self.root, val)

    def _search(self, node: BSTNode | None, val: int) -> bool:
        if node is None:
            return False
        if val < node.val:
            return self._search(node.left, val)
        elif val > node.val:
            return self._search(node.right, val)
        return True

    def min(self) -> int:
        if self.root is None:
            raise ValueError("tree is empty")
        return self._find_min(self.root).val

    def max(self) -> int:
        if self.root is None:
            raise ValueError("tree is empty")
        return self._find_max(self.root).val

    def _find_min(self, node: BSTNode) -> BSTNode:
        while node.left is not None:
            node = node.left
        return node

    def _find_max(self, node: BSTNode) -> BSTNode:
        while node.right is not None:
            node = node.right
        return node

    def inorder_traversal(self) -> list[int]:
        result: list[int] = []
        self._inorder(self.root, result)
        return result

    def _inorder(self, node: BSTNode | None, result: list[int]) -> None:
        if node is None:
            return
        self._inorder(node.left, result)
        result.append(node.val)
        self._inorder(node.right, result)

    def size(self) -> int:
        return self._size
