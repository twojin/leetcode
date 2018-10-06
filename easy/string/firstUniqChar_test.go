package string

import (
	"fmt"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	s1 := "aadadaad"
	s2 := "loveleetcode"
	fmt.Println(firstUniqChar(s1))
	fmt.Println(firstUniqChar(s2))
}
