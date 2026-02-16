package datastructures

// TrieNode represents a single node in the trie.
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// Trie is a prefix tree data structure for efficient string operations.
type Trie struct {
	root *TrieNode
}

// NewTrie creates and returns a new Trie with an initialized root node.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search returns true if the exact word exists in the trie.
func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if any word in the trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

// Delete removes a word from the trie. Returns false if the word is not found.
func (t *Trie) Delete(word string) bool {
	found := false
	t.deleteHelper(t.root, word, 0, &found)
	return found
}

// findNode traverses the trie following the given string and returns the final node,
// or nil if the path does not exist.
func (t *Trie) findNode(s string) *TrieNode {
	node := t.root
	for _, ch := range s {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

// deleteHelper recursively deletes a word from the trie.
// Returns true if the parent should remove the reference to the current node.
func (t *Trie) deleteHelper(node *TrieNode, word string, depth int, found *bool) bool {
	if node == nil {
		return false
	}

	// Reached the end of the word
	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		*found = true
		node.isEnd = false
		return !t.hasChildren(node)
	}

	idx := word[depth] - 'a'
	shouldDeleteChild := t.deleteHelper(node.children[idx], word, depth+1, found)

	if shouldDeleteChild {
		node.children[idx] = nil
		// Remove current node if it has no other children and is not end of another word
		return !node.isEnd && !t.hasChildren(node)
	}

	return false
}

// hasChildren returns true if the node has any non-nil children.
func (t *Trie) hasChildren(node *TrieNode) bool {
	for _, child := range node.children {
		if child != nil {
			return true
		}
	}
	return false
}
