class MinHeap:
    def __init__(self):
        self._items: list[int] = []

    def insert(self, val: int) -> None:
        self._items.append(val)
        self._sift_up(len(self._items) - 1)

    def extract_min(self) -> int:
        if self.is_empty():
            raise IndexError("heap is empty")
        min_val = self._items[0]
        last = self._items.pop()
        if self._items:
            self._items[0] = last
            self._sift_down(0)
        return min_val

    def peek(self) -> int:
        if self.is_empty():
            raise IndexError("heap is empty")
        return self._items[0]

    def size(self) -> int:
        return len(self._items)

    def is_empty(self) -> bool:
        return len(self._items) == 0

    def _sift_up(self, index: int) -> None:
        while index > 0:
            parent = (index - 1) // 2
            if self._items[index] < self._items[parent]:
                self._items[index], self._items[parent] = self._items[parent], self._items[index]
                index = parent
            else:
                break

    def _sift_down(self, index: int) -> None:
        size = len(self._items)
        while True:
            smallest = index
            left = 2 * index + 1
            right = 2 * index + 2
            if left < size and self._items[left] < self._items[smallest]:
                smallest = left
            if right < size and self._items[right] < self._items[smallest]:
                smallest = right
            if smallest == index:
                break
            self._items[index], self._items[smallest] = self._items[smallest], self._items[index]
            index = smallest


class MaxHeap:
    def __init__(self):
        self._items: list[int] = []

    def insert(self, val: int) -> None:
        self._items.append(val)
        self._sift_up(len(self._items) - 1)

    def extract_max(self) -> int:
        if self.is_empty():
            raise IndexError("heap is empty")
        max_val = self._items[0]
        last = self._items.pop()
        if self._items:
            self._items[0] = last
            self._sift_down(0)
        return max_val

    def peek(self) -> int:
        if self.is_empty():
            raise IndexError("heap is empty")
        return self._items[0]

    def size(self) -> int:
        return len(self._items)

    def is_empty(self) -> bool:
        return len(self._items) == 0

    def _sift_up(self, index: int) -> None:
        while index > 0:
            parent = (index - 1) // 2
            if self._items[index] > self._items[parent]:
                self._items[index], self._items[parent] = self._items[parent], self._items[index]
                index = parent
            else:
                break

    def _sift_down(self, index: int) -> None:
        size = len(self._items)
        while True:
            largest = index
            left = 2 * index + 1
            right = 2 * index + 2
            if left < size and self._items[left] > self._items[largest]:
                largest = left
            if right < size and self._items[right] > self._items[largest]:
                largest = right
            if largest == index:
                break
            self._items[index], self._items[largest] = self._items[largest], self._items[index]
            index = largest


def heapify(arr: list[int]) -> MinHeap:
    h = MinHeap()
    h._items = arr.copy()
    for i in range(len(h._items) // 2 - 1, -1, -1):
        h._sift_down(i)
    return h
