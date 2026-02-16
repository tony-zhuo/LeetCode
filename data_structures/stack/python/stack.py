class Stack:
    def __init__(self):
        self._items: list[int] = []

    def push(self, val: int) -> None:
        self._items.append(val)

    def pop(self) -> int:
        if self.is_empty():
            raise IndexError("stack is empty")
        return self._items.pop()

    def peek(self) -> int:
        if self.is_empty():
            raise IndexError("stack is empty")
        return self._items[-1]

    def is_empty(self) -> bool:
        return len(self._items) == 0

    def size(self) -> int:
        return len(self._items)
