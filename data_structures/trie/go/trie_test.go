package datastructures

import "testing"

func TestInsertAndSearch(t *testing.T) {
	t.Run("search word that exists", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if !trie.Search("hello") {
			t.Error("Search(\"hello\") = false, want true")
		}
	})

	t.Run("search word that does not exist", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if trie.Search("world") {
			t.Error("Search(\"world\") = true, want false")
		}
	})

	t.Run("search prefix should not match as exact word", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if trie.Search("hel") {
			t.Error("Search(\"hel\") = true, want false (prefix only, not a complete word)")
		}
	})

	t.Run("search empty string", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if trie.Search("") {
			t.Error("Search(\"\") = true, want false")
		}
	})

	t.Run("insert and search empty string", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("")

		if !trie.Search("") {
			t.Error("Search(\"\") = false, want true after inserting empty string")
		}
	})
}

func TestStartsWith(t *testing.T) {
	t.Run("prefix exists", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if !trie.StartsWith("hel") {
			t.Error("StartsWith(\"hel\") = false, want true")
		}
	})

	t.Run("full word as prefix", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if !trie.StartsWith("hello") {
			t.Error("StartsWith(\"hello\") = false, want true")
		}
	})

	t.Run("prefix does not exist", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if trie.StartsWith("world") {
			t.Error("StartsWith(\"world\") = true, want false")
		}
	})

	t.Run("empty prefix matches any inserted word", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if !trie.StartsWith("") {
			t.Error("StartsWith(\"\") = false, want true")
		}
	})
}

func TestSharedPrefixes(t *testing.T) {
	t.Run("insert multiple words with shared prefixes", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("apple")
		trie.Insert("app")
		trie.Insert("application")

		tests := []struct {
			word string
			want bool
		}{
			{"apple", true},
			{"app", true},
			{"application", true},
			{"ap", false},
			{"appl", false},
			{"applica", false},
		}

		for _, tt := range tests {
			t.Run("Search_"+tt.word, func(t *testing.T) {
				got := trie.Search(tt.word)
				if got != tt.want {
					t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
				}
			})
		}
	})

	t.Run("prefix matches for shared prefix words", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("apple")
		trie.Insert("app")
		trie.Insert("application")

		prefixes := []struct {
			prefix string
			want   bool
		}{
			{"a", true},
			{"ap", true},
			{"app", true},
			{"appl", true},
			{"apple", true},
			{"appli", true},
			{"applica", true},
			{"applicatio", true},
			{"application", true},
			{"applications", false},
			{"b", false},
		}

		for _, tt := range prefixes {
			t.Run("StartsWith_"+tt.prefix, func(t *testing.T) {
				got := trie.StartsWith(tt.prefix)
				if got != tt.want {
					t.Errorf("StartsWith(%q) = %v, want %v", tt.prefix, got, tt.want)
				}
			})
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word not in trie returns false", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if trie.Delete("world") {
			t.Error("Delete(\"world\") = true, want false")
		}
	})

	t.Run("delete word that is prefix of another word", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("app")
		trie.Insert("apple")

		if !trie.Delete("app") {
			t.Error("Delete(\"app\") = false, want true")
		}
		if trie.Search("app") {
			t.Error("Search(\"app\") = true after deletion, want false")
		}
		if !trie.Search("apple") {
			t.Error("Search(\"apple\") = false after deleting \"app\", want true")
		}
	})

	t.Run("delete word where another word is its prefix", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("app")
		trie.Insert("apple")

		if !trie.Delete("apple") {
			t.Error("Delete(\"apple\") = false, want true")
		}
		if trie.Search("apple") {
			t.Error("Search(\"apple\") = true after deletion, want false")
		}
		if !trie.Search("app") {
			t.Error("Search(\"app\") = false after deleting \"apple\", want true")
		}
	})

	t.Run("delete word with no shared prefix", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")
		trie.Insert("world")

		if !trie.Delete("hello") {
			t.Error("Delete(\"hello\") = false, want true")
		}
		if trie.Search("hello") {
			t.Error("Search(\"hello\") = true after deletion, want false")
		}
		if !trie.Search("world") {
			t.Error("Search(\"world\") = false after deleting \"hello\", want true")
		}
		if trie.StartsWith("hel") {
			t.Error("StartsWith(\"hel\") = true after deleting \"hello\", want false")
		}
	})

	t.Run("delete with multiple shared prefix words", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("apple")
		trie.Insert("app")
		trie.Insert("application")

		// Delete "apple" - nodes shared with "app" and "application" should remain
		if !trie.Delete("apple") {
			t.Error("Delete(\"apple\") = false, want true")
		}
		if trie.Search("apple") {
			t.Error("Search(\"apple\") = true after deletion, want false")
		}
		if !trie.Search("app") {
			t.Error("Search(\"app\") = false, want true")
		}
		if !trie.Search("application") {
			t.Error("Search(\"application\") = false, want true")
		}
	})

	t.Run("delete already deleted word returns false", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("hello")

		if !trie.Delete("hello") {
			t.Error("first Delete(\"hello\") = false, want true")
		}
		if trie.Delete("hello") {
			t.Error("second Delete(\"hello\") = true, want false")
		}
	})

	t.Run("delete partial word that was never inserted returns false", func(t *testing.T) {
		trie := NewTrie()
		trie.Insert("apple")

		if trie.Delete("app") {
			t.Error("Delete(\"app\") = true, want false (\"app\" was never inserted)")
		}
		if !trie.Search("apple") {
			t.Error("Search(\"apple\") = false, want true (should not be affected)")
		}
	})
}
