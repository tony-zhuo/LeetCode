import pytest
from solution import ListNode, Solution


def from_list(vals: list[int]) -> ListNode | None:
    head = None
    for v in reversed(vals):
        head = ListNode(v, head)
    return head


def to_list(head: ListNode | None) -> list[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


CASES = [
    ([1, 2, 3, 4, 5], [3, 4, 5]),
    ([1, 2, 3, 4, 5, 6], [4, 5, 6]),
    ([1], [1]),
    ([1, 2], [2]),
]


class TestMiddleNode:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("head,want", CASES)
    def test_middleNode(self, head, want):
        node = from_list(head)
        assert to_list(self.sol.middleNode(node)) == want
