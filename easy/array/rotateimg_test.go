package array

import (
	"fmt"
	"testing"
)

func Test_rotateImg(t *testing.T) {
	matrix1 := [][]int{{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3}}

	rotateImg(matrix1)
	for i := 0; i < len(matrix1); i++ {
		fmt.Println(matrix1[i])
	}

	matrix2 := [][]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}

	rotateImg(matrix2)
	fmt.Println(matrix2)
	for i := 0; i < len(matrix2); i++ {
		fmt.Println(matrix2[i])
	}
}
