from binary_tree import (
    TreeNode, inorder_traversal, preorder_traversal, postorder_traversal,
    level_order_traversal, height, count_nodes, slice2binary_tree,
)


def _build_sample_tree() -> TreeNode:
    """
        1
       / \\
      2   3
     / \\
    4   5
    """
    return TreeNode(1,
        TreeNode(2, TreeNode(4), TreeNode(5)),
        TreeNode(3),
    )


class TestInorder:
    def test_sample(self):
        assert inorder_traversal(_build_sample_tree()) == [4, 2, 5, 1, 3]

    def test_none(self):
        assert inorder_traversal(None) == []

    def test_single(self):
        assert inorder_traversal(TreeNode(1)) == [1]


class TestPreorder:
    def test_sample(self):
        assert preorder_traversal(_build_sample_tree()) == [1, 2, 4, 5, 3]

    def test_none(self):
        assert preorder_traversal(None) == []


class TestPostorder:
    def test_sample(self):
        assert postorder_traversal(_build_sample_tree()) == [4, 5, 2, 3, 1]

    def test_none(self):
        assert postorder_traversal(None) == []


class TestLevelOrder:
    def test_sample(self):
        assert level_order_traversal(_build_sample_tree()) == [[1], [2, 3], [4, 5]]

    def test_none(self):
        assert level_order_traversal(None) == []

    def test_single(self):
        assert level_order_traversal(TreeNode(1)) == [[1]]


class TestHeight:
    def test_sample(self):
        assert height(_build_sample_tree()) == 3

    def test_none(self):
        assert height(None) == 0

    def test_single(self):
        assert height(TreeNode(1)) == 1

    def test_left_skewed(self):
        root = TreeNode(1, TreeNode(2, TreeNode(3)))
        assert height(root) == 3


class TestCountNodes:
    def test_sample(self):
        assert count_nodes(_build_sample_tree()) == 5

    def test_none(self):
        assert count_nodes(None) == 0

    def test_single(self):
        assert count_nodes(TreeNode(1)) == 1


class TestSlice2BinaryTree:
    def test_complete_tree(self):
        root = slice2binary_tree([1, 2, 3, 4, 5])
        assert inorder_traversal(root) == [4, 2, 5, 1, 3]

    def test_with_nones(self):
        root = slice2binary_tree([1, None, 3])
        assert root.left is None
        assert root.right.val == 3

    def test_empty(self):
        assert slice2binary_tree([]) is None

    def test_none_root(self):
        assert slice2binary_tree([None]) is None
