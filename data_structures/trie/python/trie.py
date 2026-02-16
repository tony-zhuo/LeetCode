from __future__ import annotations


class TrieNode:
    def __init__(self):
        self.children: list[TrieNode | None] = [None] * 26
        self.is_end: bool = False


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            idx = ord(ch) - ord('a')
            if node.children[idx] is None:
                node.children[idx] = TrieNode()
            node = node.children[idx]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self._find_node(word)
        return node is not None and node.is_end

    def starts_with(self, prefix: str) -> bool:
        return self._find_node(prefix) is not None

    def delete(self, word: str) -> bool:
        found = [False]
        self._delete_helper(self.root, word, 0, found)
        return found[0]

    def _find_node(self, s: str) -> TrieNode | None:
        node = self.root
        for ch in s:
            idx = ord(ch) - ord('a')
            if node.children[idx] is None:
                return None
            node = node.children[idx]
        return node

    def _delete_helper(self, node: TrieNode | None, word: str, depth: int, found: list[bool]) -> bool:
        if node is None:
            return False
        if depth == len(word):
            if not node.is_end:
                return False
            found[0] = True
            node.is_end = False
            return not self._has_children(node)
        idx = ord(word[depth]) - ord('a')
        should_delete = self._delete_helper(node.children[idx], word, depth + 1, found)
        if should_delete:
            node.children[idx] = None
            return not node.is_end and not self._has_children(node)
        return False

    def _has_children(self, node: TrieNode) -> bool:
        return any(child is not None for child in node.children)
