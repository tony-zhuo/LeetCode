from hash_table import HashTable


class TestPutAndGet:
    def test_single(self):
        ht = HashTable()
        ht.put("a", 1)
        val, ok = ht.get("a")
        assert ok is True
        assert val == 1

    def test_multiple(self):
        ht = HashTable()
        ht.put("a", 1); ht.put("b", 2); ht.put("c", 3)
        assert ht.get("b") == (2, True)

    def test_overwrite(self):
        ht = HashTable()
        ht.put("a", 1)
        ht.put("a", 99)
        assert ht.get("a") == (99, True)
        assert ht.size() == 1

    def test_not_found(self):
        ht = HashTable()
        val, ok = ht.get("missing")
        assert ok is False
        assert val == 0


class TestDelete:
    def test_delete_existing(self):
        ht = HashTable()
        ht.put("a", 1); ht.put("b", 2)
        assert ht.delete("a") is True
        assert ht.contains("a") is False
        assert ht.size() == 1

    def test_delete_not_found(self):
        ht = HashTable()
        assert ht.delete("missing") is False

    def test_delete_empty(self):
        ht = HashTable()
        assert ht.delete("x") is False


class TestContains:
    def test_contains(self):
        ht = HashTable()
        ht.put("x", 42)
        assert ht.contains("x") is True
        assert ht.contains("y") is False


class TestKeys:
    def test_keys(self):
        ht = HashTable()
        ht.put("a", 1); ht.put("b", 2); ht.put("c", 3)
        assert sorted(ht.keys()) == ["a", "b", "c"]

    def test_keys_empty(self):
        assert HashTable().keys() == []


class TestSize:
    def test_empty(self):
        assert HashTable().size() == 0

    def test_after_operations(self):
        ht = HashTable()
        ht.put("a", 1); ht.put("b", 2)
        assert ht.size() == 2
        ht.delete("a")
        assert ht.size() == 1


class TestCollisions:
    def test_small_capacity(self):
        ht = HashTable(capacity=2)
        ht.put("a", 1); ht.put("b", 2); ht.put("c", 3)
        assert ht.get("a") == (1, True)
        assert ht.get("b") == (2, True)
        assert ht.get("c") == (3, True)
        assert ht.size() == 3
