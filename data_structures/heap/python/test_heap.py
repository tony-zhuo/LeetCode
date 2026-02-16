import pytest
from heap import MinHeap, MaxHeap, heapify


class TestMinHeap:
    def test_insert_and_extract(self):
        h = MinHeap()
        for v in [5, 3, 8, 1, 2]:
            h.insert(v)
        result = []
        while not h.is_empty():
            result.append(h.extract_min())
        assert result == [1, 2, 3, 5, 8]

    def test_peek(self):
        h = MinHeap()
        h.insert(5); h.insert(3); h.insert(8)
        assert h.peek() == 3
        assert h.size() == 3

    def test_empty_errors(self):
        h = MinHeap()
        with pytest.raises(IndexError):
            h.extract_min()
        with pytest.raises(IndexError):
            h.peek()

    def test_size_and_empty(self):
        h = MinHeap()
        assert h.is_empty() is True
        assert h.size() == 0
        h.insert(1)
        assert h.is_empty() is False
        assert h.size() == 1


class TestMaxHeap:
    def test_insert_and_extract(self):
        h = MaxHeap()
        for v in [5, 3, 8, 1, 2]:
            h.insert(v)
        result = []
        while not h.is_empty():
            result.append(h.extract_max())
        assert result == [8, 5, 3, 2, 1]

    def test_peek(self):
        h = MaxHeap()
        h.insert(5); h.insert(3); h.insert(8)
        assert h.peek() == 8
        assert h.size() == 3

    def test_empty_errors(self):
        h = MaxHeap()
        with pytest.raises(IndexError):
            h.extract_max()
        with pytest.raises(IndexError):
            h.peek()


class TestHeapify:
    def test_heapify(self):
        h = heapify([5, 3, 8, 1, 2])
        result = []
        while not h.is_empty():
            result.append(h.extract_min())
        assert result == [1, 2, 3, 5, 8]

    def test_heapify_sorted(self):
        h = heapify([1, 2, 3, 4, 5])
        assert h.peek() == 1

    def test_heapify_empty(self):
        h = heapify([])
        assert h.is_empty() is True
