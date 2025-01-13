package problems

type Trie struct {
	set   [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, w := range word {
		idx := w - 'a'
		if curr.set[idx] == nil {
			curr.set[idx] = &Trie{}
		}
		curr = curr.set[idx]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, w := range word {
		idx := w - 'a'
		if curr.set[idx] == nil {
			return false
		}
		curr = curr.set[idx]
	}

	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, w := range prefix {
		idx := w - 'a'
		if curr.set[idx] == nil {
			return false
		}
		curr = curr.set[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
