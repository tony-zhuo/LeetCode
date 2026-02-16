from __future__ import annotations


class _Entry:
    def __init__(self, key: str, value: int, next: _Entry | None = None):
        self.key = key
        self.value = value
        self.next = next


class HashTable:
    def __init__(self, capacity: int = 16):
        self._buckets: list[_Entry | None] = [None] * capacity
        self._size = 0
        self._cap = capacity

    def _hash(self, key: str) -> int:
        return sum(ord(c) for c in key) % self._cap

    def put(self, key: str, value: int) -> None:
        index = self._hash(key)
        curr = self._buckets[index]
        while curr is not None:
            if curr.key == key:
                curr.value = value
                return
            curr = curr.next
        self._buckets[index] = _Entry(key, value, self._buckets[index])
        self._size += 1

    def get(self, key: str) -> tuple[int, bool]:
        index = self._hash(key)
        curr = self._buckets[index]
        while curr is not None:
            if curr.key == key:
                return curr.value, True
            curr = curr.next
        return 0, False

    def delete(self, key: str) -> bool:
        index = self._hash(key)
        curr = self._buckets[index]
        if curr is None:
            return False
        if curr.key == key:
            self._buckets[index] = curr.next
            self._size -= 1
            return True
        while curr.next is not None:
            if curr.next.key == key:
                curr.next = curr.next.next
                self._size -= 1
                return True
            curr = curr.next
        return False

    def contains(self, key: str) -> bool:
        _, found = self.get(key)
        return found

    def size(self) -> int:
        return self._size

    def keys(self) -> list[str]:
        result: list[str] = []
        for bucket in self._buckets:
            curr = bucket
            while curr is not None:
                result.append(curr.key)
                curr = curr.next
        return result
