import pytest
from bst import BST


class TestBSTInsert:
    def test_insert_and_inorder(self):
        b = BST()
        for v in [5, 3, 7, 1, 4]:
            b.insert(v)
        assert b.inorder_traversal() == [1, 3, 4, 5, 7]
        assert b.size() == 5

    def test_duplicate_ignored(self):
        b = BST()
        b.insert(5); b.insert(5)
        assert b.size() == 1


class TestBSTSearch:
    def test_found(self):
        b = BST()
        for v in [5, 3, 7]:
            b.insert(v)
        assert b.search(3) is True

    def test_not_found(self):
        b = BST()
        b.insert(5)
        assert b.search(99) is False

    def test_empty(self):
        assert BST().search(1) is False


class TestBSTDelete:
    def test_delete_leaf(self):
        b = BST()
        for v in [5, 3, 7]:
            b.insert(v)
        assert b.delete(3) is True
        assert b.inorder_traversal() == [5, 7]
        assert b.size() == 2

    def test_delete_one_child(self):
        b = BST()
        for v in [5, 3, 7, 6]:
            b.insert(v)
        assert b.delete(7) is True
        assert b.inorder_traversal() == [3, 5, 6]

    def test_delete_two_children(self):
        b = BST()
        for v in [5, 3, 7, 6, 8]:
            b.insert(v)
        assert b.delete(7) is True
        assert b.inorder_traversal() == [3, 5, 6, 8]

    def test_delete_root(self):
        b = BST()
        for v in [5, 3, 7]:
            b.insert(v)
        assert b.delete(5) is True
        assert b.search(5) is False
        assert b.size() == 2

    def test_delete_not_found(self):
        b = BST()
        b.insert(5)
        assert b.delete(99) is False
        assert b.size() == 1


class TestBSTMinMax:
    def test_min_max(self):
        b = BST()
        for v in [5, 3, 7, 1, 9]:
            b.insert(v)
        assert b.min() == 1
        assert b.max() == 9

    def test_empty_tree(self):
        b = BST()
        with pytest.raises(ValueError):
            b.min()
        with pytest.raises(ValueError):
            b.max()
