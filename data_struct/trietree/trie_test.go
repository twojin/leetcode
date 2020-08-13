package trietree

import (
	"testing"
)

func Test_trie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	s1 := trie.Search("apple")
	t.Log(s1)
	s2 := trie.Search("app")
	t.Log(s2)
	s3 := trie.StartsWith("app")
	t.Log(s3)
	trie.Insert("app")
	s4 := trie.Search("app")
	t.Log(s4)
}
