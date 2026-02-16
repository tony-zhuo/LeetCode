from solution import ListNode, Solution


def _to_list(head: ListNode | None) -> list[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


def _from_list(vals: list[int]) -> ListNode | None:
    head = None
    for v in reversed(vals):
        head = ListNode(v, head)
    return head


class TestReverseList:
    def setup_method(self):
        self.sol = Solution()

    def test_reverse(self):
        head = _from_list([1, 2, 3])
        result = self.sol.reverseList(head)
        assert _to_list(result) == [3, 2, 1]

    def test_empty(self):
        assert self.sol.reverseList(None) is None

    def test_single(self):
        head = ListNode(1)
        result = self.sol.reverseList(head)
        assert _to_list(result) == [1]
