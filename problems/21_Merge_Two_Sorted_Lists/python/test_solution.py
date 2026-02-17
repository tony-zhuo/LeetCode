import pytest
from solution import Solution, ListNode


def arr2node(arr: list[int]) -> ListNode | None:
    dummy = ListNode()
    cur = dummy
    for val in arr:
        cur.next = ListNode(val)
        cur = cur.next
    return dummy.next


def node2arr(head: ListNode | None) -> list[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


class TestMergeTwoLists:
    def setup_method(self):
        self.sol = Solution()

    @pytest.mark.parametrize("list1,list2,want", [
        ([1, 2, 4], [1, 3, 4], [1, 1, 2, 3, 4, 4]),
        ([], [], []),
        ([], [0], [0]),
        ([1, 2, 3], [], [1, 2, 3]),
        ([1], [2], [1, 2]),
        ([1, 3, 5], [2, 4, 6], [1, 2, 3, 4, 5, 6]),
        ([1, 1, 1], [1, 1], [1, 1, 1, 1, 1]),
    ])
    def test_merge_two_lists(self, list1, list2, want):
        l1 = arr2node(list1)
        l2 = arr2node(list2)
        got = self.sol.mergeTwoLists(l1, l2)
        assert node2arr(got) == want
