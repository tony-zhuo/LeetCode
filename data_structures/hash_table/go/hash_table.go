package datastructures

type entry struct {
	key   string
	value int
	next  *entry
}

type HashTable struct {
	buckets []*entry
	size    int
	cap     int
}

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		buckets: make([]*entry, capacity),
		size:    0,
		cap:     capacity,
	}
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % h.cap
}

func (h *HashTable) Put(key string, value int) {
	index := h.hash(key)
	curr := h.buckets[index]

	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}

	newEntry := &entry{
		key:   key,
		value: value,
		next:  h.buckets[index],
	}
	h.buckets[index] = newEntry
	h.size++
}

func (h *HashTable) Get(key string) (int, bool) {
	index := h.hash(key)
	curr := h.buckets[index]

	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
		curr = curr.next
	}

	return 0, false
}

func (h *HashTable) Delete(key string) bool {
	index := h.hash(key)
	curr := h.buckets[index]

	if curr == nil {
		return false
	}

	if curr.key == key {
		h.buckets[index] = curr.next
		h.size--
		return true
	}

	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			h.size--
			return true
		}
		curr = curr.next
	}

	return false
}

func (h *HashTable) Contains(key string) bool {
	_, found := h.Get(key)
	return found
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) Keys() []string {
	keys := make([]string, 0, h.size)
	for _, bucket := range h.buckets {
		curr := bucket
		for curr != nil {
			keys = append(keys, curr.key)
			curr = curr.next
		}
	}
	return keys
}
