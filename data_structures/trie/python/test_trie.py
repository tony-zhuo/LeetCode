from trie import Trie


class TestTrieInsertAndSearch:
    def test_insert_and_search(self):
        t = Trie()
        t.insert("apple")
        assert t.search("apple") is True
        assert t.search("app") is False

    def test_search_empty(self):
        assert Trie().search("a") is False

    def test_multiple_words(self):
        t = Trie()
        t.insert("apple"); t.insert("app"); t.insert("banana")
        assert t.search("apple") is True
        assert t.search("app") is True
        assert t.search("banana") is True
        assert t.search("ban") is False


class TestTrieStartsWith:
    def test_prefix(self):
        t = Trie()
        t.insert("apple")
        assert t.starts_with("app") is True
        assert t.starts_with("apple") is True
        assert t.starts_with("b") is False

    def test_empty(self):
        assert Trie().starts_with("a") is False


class TestTrieDelete:
    def test_delete_word(self):
        t = Trie()
        t.insert("apple")
        assert t.delete("apple") is True
        assert t.search("apple") is False

    def test_delete_not_found(self):
        t = Trie()
        t.insert("apple")
        assert t.delete("app") is False

    def test_delete_preserves_prefix(self):
        t = Trie()
        t.insert("apple"); t.insert("app")
        assert t.delete("apple") is True
        assert t.search("apple") is False
        assert t.search("app") is True

    def test_delete_preserves_longer_word(self):
        t = Trie()
        t.insert("apple"); t.insert("app")
        assert t.delete("app") is True
        assert t.search("app") is False
        assert t.search("apple") is True

    def test_delete_empty_trie(self):
        assert Trie().delete("a") is False
