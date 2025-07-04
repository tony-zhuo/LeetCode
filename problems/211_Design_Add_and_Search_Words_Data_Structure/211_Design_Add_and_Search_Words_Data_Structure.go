package problems

type WordDictionary struct {
	nodes map[rune]*WordDictionary
	end   bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		nodes: make(map[rune]*WordDictionary),
		end:   false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	for _, w := range word {
		if _, found := this.nodes[w]; !found {
			this.nodes[w] = &WordDictionary{
				nodes: make(map[rune]*WordDictionary),
				end:   false,
			}
		}
		this = this.nodes[w]
	}

	this.end = true
}

func (this *WordDictionary) Search(word string) bool {
	for i, w := range word {
		if _, found := this.nodes[w]; !found {
			if w != '.' {
				return false
			}
			for _, n := range this.nodes {
				if n.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		this = this.nodes[w]
	}
	return this.end
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
