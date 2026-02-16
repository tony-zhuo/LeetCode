import pytest
from stack import Stack


class TestStackPushPopLIFO:
    @pytest.mark.parametrize("push_vals,want_pops", [
        ([1], [1]),
        ([1, 2, 3, 4, 5], [5, 4, 3, 2, 1]),
        ([-3, 0, 7], [7, 0, -3]),
    ])
    def test_lifo_order(self, push_vals, want_pops):
        s = Stack()
        for v in push_vals:
            s.push(v)
        for want in want_pops:
            assert s.pop() == want


class TestStackPopEmpty:
    def test_pop_new_empty_stack(self):
        with pytest.raises(IndexError):
            Stack().pop()

    def test_pop_after_all_removed(self):
        s = Stack()
        s.push(10)
        s.push(20)
        s.pop()
        s.pop()
        with pytest.raises(IndexError):
            s.pop()


class TestStackPeek:
    def test_peek_returns_top(self):
        s = Stack()
        for v in [1, 2, 3]:
            s.push(v)
        assert s.peek() == 3
        assert s.size() == 3

    def test_peek_empty(self):
        with pytest.raises(IndexError):
            Stack().peek()


class TestStackSize:
    def test_empty(self):
        assert Stack().size() == 0

    def test_after_pushes(self):
        s = Stack()
        for v in [1, 2, 3]:
            s.push(v)
        assert s.size() == 3

    def test_mixed_push_pop(self):
        s = Stack()
        s.push(1); s.push(2); s.push(3)
        s.pop(); s.push(4); s.pop(); s.pop()
        assert s.size() == 1


class TestStackIsEmpty:
    def test_new_stack(self):
        assert Stack().is_empty() is True

    def test_with_elements(self):
        s = Stack()
        s.push(1)
        assert s.is_empty() is False

    def test_after_removing_all(self):
        s = Stack()
        s.push(1); s.push(2); s.pop(); s.pop()
        assert s.is_empty() is True
