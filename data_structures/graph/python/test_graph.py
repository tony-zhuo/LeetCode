from graph import Graph


class TestGraphUndirected:
    def _build_graph(self) -> Graph:
        g = Graph(directed=False)
        g.add_edge(1, 2)
        g.add_edge(1, 3)
        g.add_edge(2, 4)
        g.add_edge(3, 4)
        return g

    def test_has_edge(self):
        g = self._build_graph()
        assert g.has_edge(1, 2) is True
        assert g.has_edge(2, 1) is True
        assert g.has_edge(1, 4) is False

    def test_neighbors(self):
        g = self._build_graph()
        assert sorted(g.neighbors(1)) == [2, 3]

    def test_vertices(self):
        g = self._build_graph()
        assert g.vertices() == [1, 2, 3, 4]

    def test_remove_edge(self):
        g = self._build_graph()
        g.remove_edge(1, 2)
        assert g.has_edge(1, 2) is False
        assert g.has_edge(2, 1) is False

    def test_add_vertex(self):
        g = Graph()
        g.add_vertex(5)
        assert 5 in g.vertices()
        assert g.neighbors(5) == []


class TestGraphDirected:
    def test_directed_edge(self):
        g = Graph(directed=True)
        g.add_edge(1, 2)
        assert g.has_edge(1, 2) is True
        assert g.has_edge(2, 1) is False


class TestBFS:
    def test_bfs(self):
        g = Graph()
        g.add_edge(1, 2); g.add_edge(1, 3); g.add_edge(2, 4); g.add_edge(3, 4)
        order = g.bfs(1)
        assert order[0] == 1
        assert set(order) == {1, 2, 3, 4}

    def test_bfs_nonexistent(self):
        g = Graph()
        assert g.bfs(99) == []

    def test_bfs_single(self):
        g = Graph()
        g.add_vertex(1)
        assert g.bfs(1) == [1]


class TestDFS:
    def test_dfs(self):
        g = Graph()
        g.add_edge(1, 2); g.add_edge(1, 3); g.add_edge(2, 4); g.add_edge(3, 4)
        order = g.dfs(1)
        assert order[0] == 1
        assert set(order) == {1, 2, 3, 4}

    def test_dfs_nonexistent(self):
        g = Graph()
        assert g.dfs(99) == []


class TestShortestPath:
    def test_shortest(self):
        g = Graph()
        g.add_edge(1, 2); g.add_edge(2, 3); g.add_edge(1, 3)
        path, dist = g.shortest_path(1, 3)
        assert path == [1, 3]
        assert dist == 1

    def test_same_node(self):
        g = Graph()
        g.add_vertex(1)
        path, dist = g.shortest_path(1, 1)
        assert path == [1]
        assert dist == 0

    def test_no_path(self):
        g = Graph()
        g.add_vertex(1); g.add_vertex(2)
        path, dist = g.shortest_path(1, 2)
        assert path is None
        assert dist == -1

    def test_nonexistent(self):
        g = Graph()
        path, dist = g.shortest_path(1, 2)
        assert path is None
        assert dist == -1
