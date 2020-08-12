package hash

import (
	"testing"
)

func Test_hashmap(t *testing.T) {
	hashMap := Constructor()
	hashMap.Put(2, 11)
	t.Log(hashMap.Get(2))
	t.Log(hashMap.Get(2))
	hashMap.Put(2071, 2071)
	t.Log(hashMap.Get(2071))
	t.Log(hashMap.Get(2))
}
