from solution import Solution


class TestClimbStairs:
    def setup_method(self):
        self.sol = Solution()

    def test_two_steps(self):
        assert self.sol.climbStairs(2) == 2

    def test_three_steps(self):
        assert self.sol.climbStairs(3) == 3

    def test_one_step(self):
        assert self.sol.climbStairs(1) == 1

    def test_five_steps(self):
        assert self.sol.climbStairs(5) == 8
