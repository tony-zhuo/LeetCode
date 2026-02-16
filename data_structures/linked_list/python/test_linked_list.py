import pytest
from linked_list import SinglyLinkedList, DoublyLinkedList, arr2node


# ── arr2node ──────────────────────────────────────────────────

class TestArr2Node:
    def test_empty(self):
        assert arr2node([]) is None

    def test_single(self):
        node = arr2node([1])
        assert node.val == 1
        assert node.next is None

    def test_multiple(self):
        node = arr2node([1, 2, 3])
        vals = []
        while node:
            vals.append(node.val)
            node = node.next
        assert vals == [1, 2, 3]


# ── Singly Linked List ───────────────────────────────────────

class TestSinglyInsert:
    def test_insert_at_head(self):
        ll = SinglyLinkedList()
        ll.insert_at_head(3); ll.insert_at_head(2); ll.insert_at_head(1)
        assert ll.to_list() == [1, 2, 3]

    def test_insert_at_tail(self):
        ll = SinglyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.to_list() == [1, 2, 3]

    @pytest.mark.parametrize("initial,index,val,want", [
        ([2, 3], 0, 1, [1, 2, 3]),
        ([1, 3], 1, 2, [1, 2, 3]),
        ([1, 2], 2, 3, [1, 2, 3]),
    ])
    def test_insert_at(self, initial, index, val, want):
        ll = SinglyLinkedList()
        for v in initial:
            ll.insert_at_tail(v)
        ll.insert_at(index, val)
        assert ll.to_list() == want

    def test_insert_at_out_of_range(self):
        ll = SinglyLinkedList()
        with pytest.raises(IndexError):
            ll.insert_at(-1, 0)
        with pytest.raises(IndexError):
            ll.insert_at(1, 0)


class TestSinglyDelete:
    def test_delete_at_head(self):
        ll = SinglyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.delete_at_head() == 1
        assert ll.to_list() == [2, 3]

    def test_delete_at_tail(self):
        ll = SinglyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.delete_at_tail() == 3
        assert ll.to_list() == [1, 2]

    def test_delete_at(self):
        ll = SinglyLinkedList()
        for v in [1, 2, 3, 4]:
            ll.insert_at_tail(v)
        assert ll.delete_at(1) == 2
        assert ll.to_list() == [1, 3, 4]

    def test_delete_empty(self):
        ll = SinglyLinkedList()
        with pytest.raises(IndexError):
            ll.delete_at_head()
        with pytest.raises(IndexError):
            ll.delete_at_tail()


class TestSinglySearchAndReverse:
    def test_search_found(self):
        ll = SinglyLinkedList()
        for v in [10, 20, 30]:
            ll.insert_at_tail(v)
        assert ll.search(20) == 1

    def test_search_not_found(self):
        ll = SinglyLinkedList()
        ll.insert_at_tail(1)
        assert ll.search(99) == -1

    def test_reverse(self):
        ll = SinglyLinkedList()
        for v in [1, 2, 3]:
            ll.insert_at_tail(v)
        ll.reverse()
        assert ll.to_list() == [3, 2, 1]

    def test_len(self):
        ll = SinglyLinkedList()
        assert len(ll) == 0
        ll.insert_at_tail(1); ll.insert_at_tail(2)
        assert len(ll) == 2


# ── Doubly Linked List ────────────────────────────────────────

class TestDoublyInsert:
    def test_insert_at_head(self):
        ll = DoublyLinkedList()
        ll.insert_at_head(3); ll.insert_at_head(2); ll.insert_at_head(1)
        assert ll.to_list() == [1, 2, 3]

    def test_insert_at_tail(self):
        ll = DoublyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.to_list() == [1, 2, 3]


class TestDoublyDelete:
    def test_delete_at_head(self):
        ll = DoublyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.delete_at_head() == 1
        assert ll.to_list() == [2, 3]

    def test_delete_at_tail(self):
        ll = DoublyLinkedList()
        ll.insert_at_tail(1); ll.insert_at_tail(2); ll.insert_at_tail(3)
        assert ll.delete_at_tail() == 3
        assert ll.to_list() == [1, 2]

    def test_delete_empty(self):
        ll = DoublyLinkedList()
        with pytest.raises(IndexError):
            ll.delete_at_head()
        with pytest.raises(IndexError):
            ll.delete_at_tail()

    def test_delete_single(self):
        ll = DoublyLinkedList()
        ll.insert_at_tail(1)
        assert ll.delete_at_head() == 1
        assert ll.head is None
        assert ll.tail is None


class TestDoublySearchAndReverse:
    def test_search(self):
        ll = DoublyLinkedList()
        for v in [10, 20, 30]:
            ll.insert_at_tail(v)
        assert ll.search(20) == 1
        assert ll.search(99) == -1

    def test_reverse(self):
        ll = DoublyLinkedList()
        for v in [1, 2, 3]:
            ll.insert_at_tail(v)
        ll.reverse()
        assert ll.to_list() == [3, 2, 1]
        assert ll.head.val == 3
        assert ll.tail.val == 1

    def test_len(self):
        ll = DoublyLinkedList()
        assert len(ll) == 0
        ll.insert_at_tail(1); ll.insert_at_tail(2)
        assert len(ll) == 2
