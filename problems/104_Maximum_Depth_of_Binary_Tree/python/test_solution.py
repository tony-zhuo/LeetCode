from solution import TreeNode, Solution


class TestMaxDepth:
    def setup_method(self):
        self.sol = Solution()

    def test_example(self):
        #     3
        #    / \
        #   9  20
        #     /  \
        #    15   7
        root = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
        assert self.sol.maxDepth(root) == 3

    def test_empty(self):
        assert self.sol.maxDepth(None) == 0

    def test_single(self):
        assert self.sol.maxDepth(TreeNode(1)) == 1

    def test_left_only(self):
        root = TreeNode(1, TreeNode(2))
        assert self.sol.maxDepth(root) == 2
