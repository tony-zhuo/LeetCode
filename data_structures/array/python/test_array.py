import pytest
from array import DynamicArray


class TestAppendAndGet:
    @pytest.mark.parametrize("values,index,want", [
        ([10], 0, 10),
        ([1, 2, 3, 4, 5], 0, 1),
        ([1, 2, 3, 4, 5], 4, 5),
        ([10, 20, 30], 1, 20),
    ])
    def test_append_and_get(self, values, index, want):
        a = DynamicArray()
        for v in values:
            a.append(v)
        assert a.get(index) == want


class TestGetOutOfRange:
    @pytest.mark.parametrize("size,index", [
        (3, -1), (3, 3), (3, 10), (0, 0),
    ])
    def test_out_of_range(self, size, index):
        a = DynamicArray()
        for i in range(size):
            a.append(i)
        with pytest.raises(IndexError):
            a.get(index)


class TestSet:
    @pytest.mark.parametrize("initial,index,val,want", [
        ([1, 2, 3], 0, 99, [99, 2, 3]),
        ([1, 2, 3], 2, 99, [1, 2, 99]),
    ])
    def test_set(self, initial, index, val, want):
        a = DynamicArray()
        for v in initial:
            a.append(v)
        a.set(index, val)
        assert a.to_list() == want


class TestInsertAt:
    @pytest.mark.parametrize("initial,index,val,want", [
        ([2, 3, 4], 0, 1, [1, 2, 3, 4]),
        ([1, 3, 4], 1, 2, [1, 2, 3, 4]),
        ([1, 2, 3], 3, 4, [1, 2, 3, 4]),
        ([], 0, 1, [1]),
    ])
    def test_insert_at(self, initial, index, val, want):
        a = DynamicArray()
        for v in initial:
            a.append(v)
        a.insert_at(index, val)
        assert a.to_list() == want

    def test_insert_out_of_range(self):
        a = DynamicArray()
        a.append(1); a.append(2)
        with pytest.raises(IndexError):
            a.insert_at(-1, 0)
        with pytest.raises(IndexError):
            a.insert_at(3, 0)


class TestDeleteAt:
    @pytest.mark.parametrize("initial,index,want_val,want_arr", [
        ([1, 2, 3], 0, 1, [2, 3]),
        ([1, 2, 3], 1, 2, [1, 3]),
        ([1, 2, 3], 2, 3, [1, 2]),
    ])
    def test_delete_at(self, initial, index, want_val, want_arr):
        a = DynamicArray()
        for v in initial:
            a.append(v)
        assert a.delete_at(index) == want_val
        assert a.to_list() == want_arr


class TestResize:
    def test_grow(self):
        a = DynamicArray(capacity=2)
        assert a.capacity() == 2
        a.append(1); a.append(2); a.append(3)
        assert a.capacity() == 4
        assert a.to_list() == [1, 2, 3]


class TestContains:
    @pytest.mark.parametrize("values,search,want", [
        ([1, 2, 3], 1, True),
        ([1, 2, 3], 3, True),
        ([1, 2, 3], 4, False),
        ([], 1, False),
    ])
    def test_contains(self, values, search, want):
        a = DynamicArray()
        for v in values:
            a.append(v)
        assert a.contains(search) == want


class TestIndexOf:
    @pytest.mark.parametrize("values,search,want", [
        ([10, 20, 30], 10, 0),
        ([10, 20, 30], 30, 2),
        ([10, 20, 30], 99, -1),
        ([1, 2, 1, 3], 1, 0),
    ])
    def test_index_of(self, values, search, want):
        a = DynamicArray()
        for v in values:
            a.append(v)
        assert a.index_of(search) == want


class TestReverse:
    @pytest.mark.parametrize("values,want", [
        ([1, 2, 3], [3, 2, 1]),
        ([1, 2, 3, 4], [4, 3, 2, 1]),
        ([1], [1]),
        ([], []),
    ])
    def test_reverse(self, values, want):
        a = DynamicArray()
        for v in values:
            a.append(v)
        a.reverse()
        assert a.to_list() == want


class TestSizeAndIsEmpty:
    def test_new_array(self):
        a = DynamicArray()
        assert a.is_empty() is True
        assert a.size() == 0

    def test_with_elements(self):
        a = DynamicArray()
        a.append(1); a.append(2)
        assert a.is_empty() is False
        assert a.size() == 2
