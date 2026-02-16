import pytest
from queue import Queue, CircularQueue


class TestQueueFIFO:
    @pytest.mark.parametrize("enqueue,want", [
        ([1], [1]),
        ([1, 2, 3, 4, 5], [1, 2, 3, 4, 5]),
    ])
    def test_fifo_order(self, enqueue, want):
        q = Queue()
        for v in enqueue:
            q.enqueue(v)
        for w in want:
            assert q.dequeue() == w


class TestQueueEmpty:
    def test_dequeue_empty(self):
        with pytest.raises(IndexError):
            Queue().dequeue()

    def test_peek_empty(self):
        with pytest.raises(IndexError):
            Queue().peek()


class TestQueuePeek:
    def test_peek_returns_front(self):
        q = Queue()
        for v in [10, 20, 30]:
            q.enqueue(v)
        assert q.peek() == 10
        assert q.size() == 3


class TestQueueSizeAndEmpty:
    def test_new_queue_empty(self):
        q = Queue()
        assert q.is_empty() is True
        assert q.size() == 0

    def test_with_elements(self):
        q = Queue()
        q.enqueue(1)
        assert q.is_empty() is False
        assert q.size() == 1


class TestCircularQueueFIFO:
    @pytest.mark.parametrize("capacity,enqueue,want", [
        (5, [1], [1]),
        (3, [1, 2, 3], [1, 2, 3]),
    ])
    def test_fifo_order(self, capacity, enqueue, want):
        q = CircularQueue(capacity)
        for v in enqueue:
            q.enqueue(v)
        for w in want:
            assert q.dequeue() == w


class TestCircularQueueWrapAround:
    def test_wrap_around(self):
        q = CircularQueue(3)
        q.enqueue(1); q.enqueue(2); q.enqueue(3)
        assert q.is_full() is True
        assert q.dequeue() == 1
        assert q.dequeue() == 2
        q.enqueue(4); q.enqueue(5)
        assert q.is_full() is True
        for want in [3, 4, 5]:
            assert q.dequeue() == want
        assert q.is_empty() is True


class TestCircularQueueFull:
    def test_enqueue_full(self):
        q = CircularQueue(2)
        q.enqueue(1); q.enqueue(2)
        with pytest.raises(OverflowError):
            q.enqueue(3)


class TestCircularQueueEmpty:
    def test_dequeue_empty(self):
        with pytest.raises(IndexError):
            CircularQueue(3).dequeue()

    def test_peek_empty(self):
        with pytest.raises(IndexError):
            CircularQueue(3).peek()


class TestCircularQueuePeek:
    def test_peek(self):
        q = CircularQueue(3)
        q.enqueue(10); q.enqueue(20)
        assert q.peek() == 10
        assert q.size() == 2


class TestCircularQueueSizeAndFlags:
    def test_is_empty(self):
        assert CircularQueue(3).is_empty() is True

    def test_not_empty(self):
        q = CircularQueue(3)
        q.enqueue(1)
        assert q.is_empty() is False

    def test_not_full(self):
        q = CircularQueue(3)
        q.enqueue(1)
        assert q.is_full() is False

    def test_full(self):
        q = CircularQueue(2)
        q.enqueue(1); q.enqueue(2)
        assert q.is_full() is True

    def test_size(self):
        q = CircularQueue(5)
        q.enqueue(1); q.enqueue(2)
        assert q.size() == 2
