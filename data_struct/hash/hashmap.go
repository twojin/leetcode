package hash

type MyHashMap struct {
	hashTable []*Backet
	keySpace  int
}

type Backet struct {
	Key  int
	Val  int
	Next *Backet
}

func (bkt *Backet) update(key, val int) *Backet {
	found := false
	pair := bkt
	for pair != nil {
		if pair.Key == key {
			pair.Val = val
			found = true
		}

		pair = pair.Next
	}

	if found {
		return bkt
	}

	bkt = &Backet{Key: key, Val: val, Next: bkt}
	return bkt
}

func (bkt *Backet) get(key int) int {
	pair := bkt
	for pair != nil {
		if pair.Key == key {
			return pair.Val
		}

		pair = pair.Next
	}

	return -1
}

func (bkt *Backet) remove(key int) *Backet {
	head := &Backet{Next: bkt}
	prev := head
	curr := prev.Next
	for curr != nil {
		if curr.Key == key {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}

	return head.Next
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	mhm := MyHashMap{}
	mhm.keySpace = 2069
	mhm.hashTable = make([]*Backet, mhm.keySpace)

	return mhm
}

func (this *MyHashMap) hash(key int) int {
	return key % this.keySpace
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hashIndex := this.hash(key)
	this.hashTable[hashIndex] = this.hashTable[hashIndex].update(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashIndex := this.hash(key)
	return this.hashTable[hashIndex].get(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashIndex := this.hash(key)
	this.hashTable[hashIndex] = this.hashTable[hashIndex].remove(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
