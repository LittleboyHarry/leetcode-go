package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode-go/common/test-utils/singlylist"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	assert.Equal(t,
		[]int{7, 0, 8},
		singlylist.ToSlice(addTwoNumbers(
			singlylist.Make(2, 4, 3),
			singlylist.Make(5, 6, 4))),
	)

	assert.Equal(t,
		[]int{0},
		singlylist.ToSlice(addTwoNumbers(
			singlylist.Make(0),
			singlylist.Make(0))),
	)

	assert.Equal(t,
		[]int{8, 9, 9, 9, 0, 0, 0, 1},
		singlylist.ToSlice(addTwoNumbers(
			singlylist.Make(9, 9, 9, 9, 9, 9, 9),
			singlylist.Make(9, 9, 9, 9))),
	)
}
