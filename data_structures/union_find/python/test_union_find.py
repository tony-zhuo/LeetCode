from union_find import UnionFind


class TestUnionFind:
    def test_initial_state(self):
        uf = UnionFind(5)
        assert uf.count() == 5
        for i in range(5):
            assert uf.find(i) == i

    def test_union_and_connected(self):
        uf = UnionFind(5)
        assert uf.union(0, 1) is True
        assert uf.connected(0, 1) is True
        assert uf.count() == 4

    def test_already_connected(self):
        uf = UnionFind(5)
        uf.union(0, 1)
        assert uf.union(0, 1) is False
        assert uf.count() == 4

    def test_transitivity(self):
        uf = UnionFind(5)
        uf.union(0, 1)
        uf.union(1, 2)
        assert uf.connected(0, 2) is True
        assert uf.count() == 3

    def test_separate_components(self):
        uf = UnionFind(5)
        uf.union(0, 1)
        uf.union(2, 3)
        assert uf.connected(0, 1) is True
        assert uf.connected(2, 3) is True
        assert uf.connected(0, 2) is False
        assert uf.count() == 3

    def test_merge_all(self):
        uf = UnionFind(4)
        uf.union(0, 1)
        uf.union(2, 3)
        uf.union(0, 2)
        assert uf.count() == 1
        for i in range(4):
            for j in range(4):
                assert uf.connected(i, j) is True

    def test_path_compression(self):
        uf = UnionFind(5)
        uf.union(0, 1)
        uf.union(1, 2)
        uf.union(2, 3)
        uf.find(3)
        # After path compression, 3's parent should be the root
        assert uf._parent[3] == uf.find(0)
