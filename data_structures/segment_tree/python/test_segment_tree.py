import math
from segment_tree import SegmentTree, RangeMinSegmentTree


class TestSegmentTree:
    def test_range_sum(self):
        st = SegmentTree([1, 3, 5, 7, 9, 11])
        assert st.query(0, 5) == 36
        assert st.query(1, 3) == 15
        assert st.query(0, 0) == 1

    def test_update(self):
        st = SegmentTree([1, 3, 5, 7, 9, 11])
        st.update(1, 10)
        assert st.query(0, 5) == 43
        assert st.query(1, 1) == 10

    def test_single_element(self):
        st = SegmentTree([42])
        assert st.query(0, 0) == 42
        st.update(0, 10)
        assert st.query(0, 0) == 10

    def test_multiple_updates(self):
        st = SegmentTree([1, 2, 3, 4, 5])
        st.update(0, 10)
        st.update(4, 20)
        assert st.query(0, 4) == 39


class TestRangeMinSegmentTree:
    def test_range_min(self):
        st = RangeMinSegmentTree([5, 2, 8, 1, 4, 3])
        assert st.query(0, 5) == 1
        assert st.query(0, 2) == 2
        assert st.query(3, 5) == 1

    def test_update(self):
        st = RangeMinSegmentTree([5, 2, 8, 1, 4, 3])
        st.update(3, 10)
        assert st.query(3, 5) == 3
        assert st.query(0, 5) == 2

    def test_single_element(self):
        st = RangeMinSegmentTree([42])
        assert st.query(0, 0) == 42

    def test_out_of_range_returns_inf(self):
        st = RangeMinSegmentTree([1, 2, 3])
        # query completely out of range returns inf
        assert st.query(3, 5) == math.inf
