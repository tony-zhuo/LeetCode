package datastructures

import (
	"sort"
	"testing"
)

func TestHashTable_PutAndGet(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     int
		wantValue int
		wantFound bool
	}{
		{
			name:      "put and get single entry",
			key:       "apple",
			value:     10,
			wantValue: 10,
			wantFound: true,
		},
		{
			name:      "put and get empty string key",
			key:       "",
			value:     0,
			wantValue: 0,
			wantFound: true,
		},
		{
			name:      "put and get long key",
			key:       "a very long key string for testing",
			value:     999,
			wantValue: 999,
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			ht.Put(tt.key, tt.value)

			got, found := ht.Get(tt.key)
			if found != tt.wantFound {
				t.Errorf("Get() found = %v, want %v", found, tt.wantFound)
			}
			if got != tt.wantValue {
				t.Errorf("Get() value = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

func TestHashTable_OverwriteExistingKey(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		firstValue   int
		secondValue  int
		wantValue    int
		wantSize     int
	}{
		{
			name:        "overwrite updates value",
			key:         "apple",
			firstValue:  10,
			secondValue: 20,
			wantValue:   20,
			wantSize:    1,
		},
		{
			name:        "overwrite with same value",
			key:         "banana",
			firstValue:  5,
			secondValue: 5,
			wantValue:   5,
			wantSize:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			ht.Put(tt.key, tt.firstValue)
			ht.Put(tt.key, tt.secondValue)

			got, found := ht.Get(tt.key)
			if !found {
				t.Fatal("Get() found = false, want true")
			}
			if got != tt.wantValue {
				t.Errorf("Get() value = %v, want %v", got, tt.wantValue)
			}
			if ht.Size() != tt.wantSize {
				t.Errorf("Size() = %v, want %v", ht.Size(), tt.wantSize)
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	tests := []struct {
		name       string
		putKeys    []string
		putValues  []int
		deleteKey  string
		wantDelete bool
		wantSize   int
	}{
		{
			name:       "delete existing key",
			putKeys:    []string{"apple"},
			putValues:  []int{10},
			deleteKey:  "apple",
			wantDelete: true,
			wantSize:   0,
		},
		{
			name:       "delete non-existing key",
			putKeys:    []string{"apple"},
			putValues:  []int{10},
			deleteKey:  "banana",
			wantDelete: false,
			wantSize:   1,
		},
		{
			name:       "delete from empty table",
			putKeys:    []string{},
			putValues:  []int{},
			deleteKey:  "apple",
			wantDelete: false,
			wantSize:   0,
		},
		{
			name:       "delete one of multiple keys",
			putKeys:    []string{"apple", "banana", "cherry"},
			putValues:  []int{10, 20, 30},
			deleteKey:  "banana",
			wantDelete: true,
			wantSize:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			for i, key := range tt.putKeys {
				ht.Put(key, tt.putValues[i])
			}

			deleted := ht.Delete(tt.deleteKey)
			if deleted != tt.wantDelete {
				t.Errorf("Delete() = %v, want %v", deleted, tt.wantDelete)
			}
			if ht.Size() != tt.wantSize {
				t.Errorf("Size() = %v, want %v", ht.Size(), tt.wantSize)
			}
			if deleted && ht.Contains(tt.deleteKey) {
				t.Errorf("Contains(%q) = true after deletion, want false", tt.deleteKey)
			}
		})
	}
}

func TestHashTable_GetNonExisting(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		wantValue int
		wantFound bool
	}{
		{
			name:      "get from empty table",
			key:       "apple",
			wantValue: 0,
			wantFound: false,
		},
		{
			name:      "get non-existing key from non-empty table",
			key:       "missing",
			wantValue: 0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			if tt.name == "get non-existing key from non-empty table" {
				ht.Put("exists", 42)
			}

			got, found := ht.Get(tt.key)
			if found != tt.wantFound {
				t.Errorf("Get() found = %v, want %v", found, tt.wantFound)
			}
			if got != tt.wantValue {
				t.Errorf("Get() value = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

func TestHashTable_Contains(t *testing.T) {
	tests := []struct {
		name     string
		putKeys  []string
		checkKey string
		want     bool
	}{
		{
			name:     "contains existing key",
			putKeys:  []string{"apple", "banana"},
			checkKey: "apple",
			want:     true,
		},
		{
			name:     "does not contain missing key",
			putKeys:  []string{"apple", "banana"},
			checkKey: "cherry",
			want:     false,
		},
		{
			name:     "empty table contains nothing",
			putKeys:  []string{},
			checkKey: "apple",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			for _, key := range tt.putKeys {
				ht.Put(key, 1)
			}

			if got := ht.Contains(tt.checkKey); got != tt.want {
				t.Errorf("Contains(%q) = %v, want %v", tt.checkKey, got, tt.want)
			}
		})
	}
}

func TestHashTable_Keys(t *testing.T) {
	tests := []struct {
		name     string
		putKeys  []string
		wantKeys []string
	}{
		{
			name:     "empty table returns empty keys",
			putKeys:  []string{},
			wantKeys: []string{},
		},
		{
			name:     "single key",
			putKeys:  []string{"apple"},
			wantKeys: []string{"apple"},
		},
		{
			name:     "multiple keys",
			putKeys:  []string{"apple", "banana", "cherry"},
			wantKeys: []string{"apple", "banana", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			for _, key := range tt.putKeys {
				ht.Put(key, 1)
			}

			got := ht.Keys()
			sort.Strings(got)
			sort.Strings(tt.wantKeys)

			if len(got) != len(tt.wantKeys) {
				t.Fatalf("Keys() length = %d, want %d", len(got), len(tt.wantKeys))
			}
			for i := range got {
				if got[i] != tt.wantKeys[i] {
					t.Errorf("Keys()[%d] = %q, want %q", i, got[i], tt.wantKeys[i])
				}
			}
		})
	}
}

func TestHashTable_SizeTracking(t *testing.T) {
	tests := []struct {
		name     string
		ops      []struct {
			action string // "put" or "delete"
			key    string
			value  int
		}
		wantSize int
	}{
		{
			name: "size after multiple puts",
			ops: []struct {
				action string
				key    string
				value  int
			}{
				{"put", "a", 1},
				{"put", "b", 2},
				{"put", "c", 3},
			},
			wantSize: 3,
		},
		{
			name: "size after put and delete",
			ops: []struct {
				action string
				key    string
				value  int
			}{
				{"put", "a", 1},
				{"put", "b", 2},
				{"put", "c", 3},
				{"delete", "b", 0},
			},
			wantSize: 2,
		},
		{
			name: "size after overwrite does not increase",
			ops: []struct {
				action string
				key    string
				value  int
			}{
				{"put", "a", 1},
				{"put", "a", 2},
				{"put", "a", 3},
			},
			wantSize: 1,
		},
		{
			name: "size after delete non-existing key stays same",
			ops: []struct {
				action string
				key    string
				value  int
			}{
				{"put", "a", 1},
				{"delete", "z", 0},
			},
			wantSize: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := NewHashTable(16)
			for _, op := range tt.ops {
				switch op.action {
				case "put":
					ht.Put(op.key, op.value)
				case "delete":
					ht.Delete(op.key)
				}
			}

			if ht.Size() != tt.wantSize {
				t.Errorf("Size() = %v, want %v", ht.Size(), tt.wantSize)
			}
		})
	}
}

func TestHashTable_CollisionHandling(t *testing.T) {
	// Use a very small capacity to force collisions
	ht := NewHashTable(2)

	keys := []struct {
		key   string
		value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"ab", 5},
		{"ba", 6},
	}

	t.Run("put and get with collisions", func(t *testing.T) {
		for _, kv := range keys {
			ht.Put(kv.key, kv.value)
		}

		for _, kv := range keys {
			got, found := ht.Get(kv.key)
			if !found {
				t.Errorf("Get(%q) found = false, want true", kv.key)
			}
			if got != kv.value {
				t.Errorf("Get(%q) = %v, want %v", kv.key, got, kv.value)
			}
		}

		if ht.Size() != len(keys) {
			t.Errorf("Size() = %v, want %v", ht.Size(), len(keys))
		}
	})

	t.Run("delete with collisions", func(t *testing.T) {
		// Delete a key that shares a bucket with others
		ht.Delete("c")
		if ht.Contains("c") {
			t.Error("Contains(\"c\") = true after deletion, want false")
		}

		// Verify other keys are still accessible
		for _, kv := range keys {
			if kv.key == "c" {
				continue
			}
			got, found := ht.Get(kv.key)
			if !found {
				t.Errorf("Get(%q) found = false after deleting \"c\", want true", kv.key)
			}
			if got != kv.value {
				t.Errorf("Get(%q) = %v after deleting \"c\", want %v", kv.key, got, kv.value)
			}
		}
	})

	t.Run("overwrite with collisions", func(t *testing.T) {
		sizeBefore := ht.Size()
		ht.Put("a", 100)

		got, found := ht.Get("a")
		if !found {
			t.Fatal("Get(\"a\") found = false, want true")
		}
		if got != 100 {
			t.Errorf("Get(\"a\") = %v, want 100", got)
		}
		if ht.Size() != sizeBefore {
			t.Errorf("Size() = %v after overwrite, want %v", ht.Size(), sizeBefore)
		}
	})
}
