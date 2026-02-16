from collections import deque


class Queue:
    def __init__(self):
        self._items: deque[int] = deque()

    def enqueue(self, val: int) -> None:
        self._items.append(val)

    def dequeue(self) -> int:
        if self.is_empty():
            raise IndexError("queue is empty")
        return self._items.popleft()

    def peek(self) -> int:
        if self.is_empty():
            raise IndexError("queue is empty")
        return self._items[0]

    def is_empty(self) -> bool:
        return len(self._items) == 0

    def size(self) -> int:
        return len(self._items)


class CircularQueue:
    def __init__(self, capacity: int):
        self._items = [0] * capacity
        self._head = 0
        self._tail = 0
        self._size = 0
        self._cap = capacity

    def enqueue(self, val: int) -> None:
        if self.is_full():
            raise OverflowError("queue is full")
        self._items[self._tail] = val
        self._tail = (self._tail + 1) % self._cap
        self._size += 1

    def dequeue(self) -> int:
        if self.is_empty():
            raise IndexError("queue is empty")
        front = self._items[self._head]
        self._head = (self._head + 1) % self._cap
        self._size -= 1
        return front

    def peek(self) -> int:
        if self.is_empty():
            raise IndexError("queue is empty")
        return self._items[self._head]

    def is_empty(self) -> bool:
        return self._size == 0

    def is_full(self) -> bool:
        return self._size == self._cap

    def size(self) -> int:
        return self._size
