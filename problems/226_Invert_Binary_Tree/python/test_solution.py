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


def tree2list(root: TreeNode | None) -> list[int | None]:
    if root is None:
        return []
    result = []
    queue = [root]
    while queue:
        node = queue.pop(0)
        if node is None:
            result.append(None)
        else:
            result.append(node.val)
            queue.append(node.left)
            queue.append(node.right)
    # trim trailing Nones
    while result and result[-1] is None:
        result.pop()
    return result


class TestInvertTree:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("root,want", [
        ([4, 2, 7, 1, 3, 6, 9], [4, 7, 2, 9, 6, 3, 1]),
        ([2, 1, 3], [2, 3, 1]),
        ([], []),
        ([1], [1]),
        ([1, 2], [1, None, 2]),
    ])
    def test_invert_tree(self, root, want):
        tree = list2tree(root)
        got = self.sol.invertTree(tree)
        assert tree2list(got) == want
