import pytest
from solution import Solution, TreeNode


def list2tree(arr: list[int | None]) -> TreeNode | None:
    if not arr or arr[0] is None:
        return None
    root = TreeNode(arr[0])
    queue = [root]
    i = 1
    while i < len(arr):
        node = queue.pop(0)
        if i < len(arr) and arr[i] is not None:
            node.left = TreeNode(arr[i])
            queue.append(node.left)
        i += 1
        if i < len(arr) and arr[i] is not None:
            node.right = TreeNode(arr[i])
            queue.append(node.right)
        i += 1
    return root


CASES = [
    ([1, 2, 3, 4, 5], 3),
    ([1, 2], 1),
    ([1], 0),
    ([1, 2, None, 3, 4, 5, None, None, None, 6], 4),
]


class TestDiameterOfBinaryTree:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("root,want", CASES)
    def test_diameterOfBinaryTree(self, root, want):
        tree = list2tree(root)
        assert self.sol.diameterOfBinaryTree(tree) == want
