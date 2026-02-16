_DEFAULT_CAPACITY = 8


class DynamicArray:
    def __init__(self, capacity: int = _DEFAULT_CAPACITY):
        if capacity <= 0:
            capacity = _DEFAULT_CAPACITY
        self._data = [0] * capacity
        self._size = 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self._size:
            raise IndexError("index out of range")
        return self._data[index]

    def set(self, index: int, val: int) -> None:
        if index < 0 or index >= self._size:
            raise IndexError("index out of range")
        self._data[index] = val

    def append(self, val: int) -> None:
        if self._size == len(self._data):
            self._resize(len(self._data) * 2)
        self._data[self._size] = val
        self._size += 1

    def insert_at(self, index: int, val: int) -> None:
        if index < 0 or index > self._size:
            raise IndexError("index out of range")
        if self._size == len(self._data):
            self._resize(len(self._data) * 2)
        for i in range(self._size, index, -1):
            self._data[i] = self._data[i - 1]
        self._data[index] = val
        self._size += 1

    def delete_at(self, index: int) -> int:
        if index < 0 or index >= self._size:
            raise IndexError("index out of range")
        val = self._data[index]
        for i in range(index, self._size - 1):
            self._data[i] = self._data[i + 1]
        self._size -= 1
        if self._size > 0 and self._size == len(self._data) // 4:
            self._resize(len(self._data) // 2)
        return val

    def size(self) -> int:
        return self._size

    def capacity(self) -> int:
        return len(self._data)

    def is_empty(self) -> bool:
        return self._size == 0

    def contains(self, val: int) -> bool:
        for i in range(self._size):
            if self._data[i] == val:
                return True
        return False

    def index_of(self, val: int) -> int:
        for i in range(self._size):
            if self._data[i] == val:
                return i
        return -1

    def reverse(self) -> None:
        i, j = 0, self._size - 1
        while i < j:
            self._data[i], self._data[j] = self._data[j], self._data[i]
            i += 1
            j -= 1

    def to_list(self) -> list[int]:
        return self._data[:self._size].copy()

    def _resize(self, new_cap: int) -> None:
        new_data = [0] * new_cap
        for i in range(self._size):
            new_data[i] = self._data[i]
        self._data = new_data
