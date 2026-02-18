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


def find_node(root: TreeNode | None, val: int) -> TreeNode | None:
    if root is None:
        return None
    if root.val == val:
        return root
    left = find_node(root.left, val)
    if left:
        return left
    return find_node(root.right, val)


CASES = [
    ([6, 2, 8, 0, 4, 7, 9, None, None, 3, 5], 2, 8, 6),
    ([6, 2, 8, 0, 4, 7, 9, None, None, 3, 5], 2, 4, 2),
    ([2, 1], 2, 1, 2),
    ([6, 2, 8, 0, 4, 7, 9], 7, 9, 8),
    ([6, 2, 8], 2, 2, 2),
]


class TestLowestCommonAncestor:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("root,p,q,want", CASES)
    def test_lowestCommonAncestor(self, root, p, q, want):
        tree = list2tree(root)
        p_node = find_node(tree, p)
        q_node = find_node(tree, q)
        got = self.sol.lowestCommonAncestor(tree, p_node, q_node)
        assert got is not None and got.val == want