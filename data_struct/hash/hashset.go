package hash

type MyHashSet struct {
	buckets  []*LinkedList
	keyRange int
}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

/** Initialize your data structure here. */
func _Constructor() MyHashSet {
	mhs := MyHashSet{}
	mhs.keyRange = 769
	mhs.buckets = make([]*LinkedList, mhs.keyRange)
	return mhs
}

func (this *MyHashSet) hash(key int) int {
	return key % this.keyRange
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

	hashIndex := this.hash(key)
	bkt := this.buckets[hashIndex]
	if bkt == nil {
		this.buckets[hashIndex] = &LinkedList{Val: key}
	} else {
		this.buckets[hashIndex] = &LinkedList{Val: key, Next: bkt}
	}
}

func (this *MyHashSet) Remove(key int) {
	if !this.Contains(key) {
		return
	}

	hashIndex := this.hash(key)
	bkt := this.buckets[hashIndex]
	if bkt == nil {
		return
	}

	this.buckets[hashIndex] = bkt.remove(key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hashIndex := this.hash(key)
	bkt := this.buckets[hashIndex]
	if bkt == nil {
		return false
	}

	return bkt.contains(key)
}

func (bkt *LinkedList) contains(key int) bool {
	node := bkt
	for node != nil {
		if node.Val == key {
			return true
		}
	}

	return false
}

func (bkt *LinkedList) remove(key int) *LinkedList {
	head := &LinkedList{Next: bkt}
	prev := head
	curr := prev.Next
	for curr != nil {
		if curr.Val == key {
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}

	return head.Next
}
