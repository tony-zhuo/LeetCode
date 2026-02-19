import pytest
from solution import Solution, ListNode


def make_linked_list(arr: list[int], pos: int) -> ListNode | None:
    """Build a linked list from arr. If pos >= 0, connect tail to node at index pos."""
    if not arr:
        return None
    nodes = [ListNode(v) for v in arr]
    for i in range(len(nodes) - 1):
        nodes[i].next = nodes[i + 1]
    if pos >= 0:
        nodes[-1].next = nodes[pos]
    return nodes[0]


CASES = [
    ([3, 2, 0, -4], 1, True),
    ([1, 2], 0, True),
    ([1], -1, False),
    ([], -1, False),
    ([1, 2, 3, 4, 5], -1, False),
    ([1, 2, 3, 4, 5], 4, True),
    ([1], 0, True),
]


class TestHasCycle:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("arr,pos,want", CASES)
    def test_hasCycle(self, arr, pos, want):
        head = make_linked_list(arr, pos)
        assert self.sol.hasCycle(head) == want
