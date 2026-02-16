from __future__ import annotations


class ListNode:
    """Simple linked list node used by problem solutions."""
    def __init__(self, val: int = 0, next: ListNode | None = None):
        self.val = val
        self.next = next


def arr2node(values: list[int]) -> ListNode | None:
    if not values:
        return None
    head = ListNode(values[0])
    curr = head
    for v in values[1:]:
        curr.next = ListNode(v)
        curr = curr.next
    return head


# ── Singly Linked List ───────────────────────────────────────

class SinglyNode:
    def __init__(self, val: int = 0, next: SinglyNode | None = None):
        self.val = val
        self.next = next


class SinglyLinkedList:
    def __init__(self):
        self.head: SinglyNode | None = None
        self._size = 0

    def insert_at_head(self, val: int) -> None:
        self.head = SinglyNode(val, self.head)
        self._size += 1

    def insert_at_tail(self, val: int) -> None:
        node = SinglyNode(val)
        if self.head is None:
            self.head = node
            self._size += 1
            return
        curr = self.head
        while curr.next is not None:
            curr = curr.next
        curr.next = node
        self._size += 1

    def insert_at(self, index: int, val: int) -> None:
        if index < 0 or index > self._size:
            raise IndexError("index out of range")
        if index == 0:
            self.insert_at_head(val)
            return
        curr = self.head
        for _ in range(index - 1):
            curr = curr.next
        curr.next = SinglyNode(val, curr.next)
        self._size += 1

    def delete_at_head(self) -> int:
        if self.head is None:
            raise IndexError("list is empty")
        val = self.head.val
        self.head = self.head.next
        self._size -= 1
        return val

    def delete_at_tail(self) -> int:
        if self.head is None:
            raise IndexError("list is empty")
        if self.head.next is None:
            val = self.head.val
            self.head = None
            self._size -= 1
            return val
        curr = self.head
        while curr.next.next is not None:
            curr = curr.next
        val = curr.next.val
        curr.next = None
        self._size -= 1
        return val

    def delete_at(self, index: int) -> int:
        if self.head is None:
            raise IndexError("list is empty")
        if index < 0 or index >= self._size:
            raise IndexError("index out of range")
        if index == 0:
            return self.delete_at_head()
        curr = self.head
        for _ in range(index - 1):
            curr = curr.next
        val = curr.next.val
        curr.next = curr.next.next
        self._size -= 1
        return val

    def search(self, val: int) -> int:
        curr = self.head
        i = 0
        while curr is not None:
            if curr.val == val:
                return i
            curr = curr.next
            i += 1
        return -1

    def reverse(self) -> None:
        prev = None
        curr = self.head
        while curr is not None:
            nxt = curr.next
            curr.next = prev
            prev = curr
            curr = nxt
        self.head = prev

    def __len__(self) -> int:
        return self._size

    def to_list(self) -> list[int]:
        result = []
        curr = self.head
        while curr is not None:
            result.append(curr.val)
            curr = curr.next
        return result


# ── Doubly Linked List ────────────────────────────────────────

class DoublyNode:
    def __init__(self, val: int = 0, prev: DoublyNode | None = None, next: DoublyNode | None = None):
        self.val = val
        self.prev = prev
        self.next = next


class DoublyLinkedList:
    def __init__(self):
        self.head: DoublyNode | None = None
        self.tail: DoublyNode | None = None
        self._size = 0

    def insert_at_head(self, val: int) -> None:
        node = DoublyNode(val, next=self.head)
        if self.head is not None:
            self.head.prev = node
        else:
            self.tail = node
        self.head = node
        self._size += 1

    def insert_at_tail(self, val: int) -> None:
        node = DoublyNode(val, prev=self.tail)
        if self.tail is not None:
            self.tail.next = node
        else:
            self.head = node
        self.tail = node
        self._size += 1

    def delete_at_head(self) -> int:
        if self.head is None:
            raise IndexError("list is empty")
        val = self.head.val
        self.head = self.head.next
        if self.head is not None:
            self.head.prev = None
        else:
            self.tail = None
        self._size -= 1
        return val

    def delete_at_tail(self) -> int:
        if self.tail is None:
            raise IndexError("list is empty")
        val = self.tail.val
        self.tail = self.tail.prev
        if self.tail is not None:
            self.tail.next = None
        else:
            self.head = None
        self._size -= 1
        return val

    def search(self, val: int) -> int:
        curr = self.head
        i = 0
        while curr is not None:
            if curr.val == val:
                return i
            curr = curr.next
            i += 1
        return -1

    def reverse(self) -> None:
        curr = self.head
        while curr is not None:
            curr.prev, curr.next = curr.next, curr.prev
            curr = curr.prev
        self.head, self.tail = self.tail, self.head

    def __len__(self) -> int:
        return self._size

    def to_list(self) -> list[int]:
        result = []
        curr = self.head
        while curr is not None:
            result.append(curr.val)
            curr = curr.next
        return result
