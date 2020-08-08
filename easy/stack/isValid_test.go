package stack

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	s := "([)"
	t.Log(isValid(s))
}
