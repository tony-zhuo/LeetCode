from solution import MyQueue


class TestMyQueue:
    def test_example(self):
        q = MyQueue()
        q.push(1)
        q.push(2)
        assert q.peek() == 1
        assert q.pop() == 1
        assert q.empty() is False

    def test_empty_check(self):
        q = MyQueue()
        assert q.empty() is True
        q.push(1)
        assert q.empty() is False
        q.pop()
        assert q.empty() is True

    def test_fifo_order(self):
        q = MyQueue()
        for i in range(1, 6):
            q.push(i)
        for i in range(1, 6):
            assert q.pop() == i

    def test_interleaved_push_and_pop(self):
        q = MyQueue()
        q.push(1)
        q.push(2)
        assert q.pop() == 1
        q.push(3)
        assert q.pop() == 2
        assert q.pop() == 3

    def test_peek_does_not_remove(self):
        q = MyQueue()
        q.push(42)
        assert q.peek() == 42
        assert q.peek() == 42
        assert q.empty() is False
