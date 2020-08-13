package trietree

type Trie struct {
	isEnd bool
	links [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

func (t *Trie) containsKey(char byte) bool {
	return t.links[char-'a'] != nil
}

func (t *Trie) get(char byte) *Trie {
	return t.links[char-'a']
}
func (t *Trie) put(char byte, node *Trie) {
	t.links[char-'a'] = node
}

func (t *Trie) setEnd() {
	t.isEnd = true
}

func (t *Trie) searchPrefix(word string) *Trie {
	node := t
	for _, v := range word {
		ch := byte(v)
		if !node.containsKey(ch) {
			return nil
		}
		node = node.get(ch)
	}

	return node
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		ch := byte(v)
		if node.get(ch) == nil {
			node.put(ch, &Trie{})
		}
		node = node.get(ch)
	}
	node.setEnd()
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
