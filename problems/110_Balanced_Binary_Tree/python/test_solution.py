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
    ([3, 9, 20, None, None, 15, 7], True),
    ([1, 2, 2, 3, 3, None, None, 4, 4], False),
    ([], True),
    ([1], True),
    ([1, 2, None, 3], False),
    ([1, 2, 3, 4, 5, 6, None, 8], True),
]


class TestIsBalanced:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("root,want", CASES)
    def test_isBalanced(self, root, want):
        tree = list2tree(root)
        assert self.sol.isBalanced(tree) == want
