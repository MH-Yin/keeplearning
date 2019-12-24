package algorithm

import (
	"github.com/magiconair/properties/assert"
	"keeplearning/structure"
	"testing"
)

func Test_isListPalindrome(t *testing.T) {
	tests := []struct {
		li     structure.List
		insert []int
		result bool
	}{
		{
			li:     structure.NewList(),
			insert: []int{},
			result: true,
		},
		{
			li:     structure.NewList(),
			insert: []int{1, 2, 3, 4, 3, 2, 1},
			result: true,
		},
		{
			li:     structure.NewList(),
			insert: []int{0, 1, 2, 3, 4, 5, 3, 2, 1, 0},
			result: false,
		},
		{
			li:     structure.NewList(),
			insert: []int{1,2,3,4,5},
			result: false,
		},
		{
			li:     structure.NewList(),
			insert: []int{1},
			result: true,
		},
	}

	for _, test :=  range tests {
		for _, v := range test.insert {
			test.li.Insert(v)
		}
		assert.Equal(t, test.result, isListPalindrome(test.li))
	}
}
