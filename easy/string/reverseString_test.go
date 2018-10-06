package string

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	s1 := "hello"
	s2 := "amanaP :lanac a ,nalp a ,nam A"

	rs1 := reverseString(s1)
	rs2 := reverseString(s2)
	fmt.Println(rs1)
	fmt.Println(rs2)
}
