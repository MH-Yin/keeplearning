package algorithm

import (
	"github.com/magiconair/properties/assert"
	"keeplearning/structure"
	"testing"
)

var tests = []struct {
	insert []int
	result bool
}{
	{
		insert: []int{},
		result: true,
	},
	{
		insert: []int{1, 2, 3, 4, 3, 2, 1},
		result: true,
	},
	{
		insert: []int{0, 1, 2, 3, 4, 5, 3, 2, 1, 0},
		result: false,
	},
	{
		insert: []int{1, 2, 3, 4, 5},
		result: false,
	},
	{
		insert: []int{1},
		result: true,
	},
}

func Test_isSinglyListPalindrome(t *testing.T) {
	for _, test := range tests {
		li := structure.NewSingleList()
		for _, v := range test.insert {
			li.Insert(v)
		}
		assert.Equal(t, test.result, isSinglyListPalindrome(li))
	}
}

func Test_isDoublyListPalindrome(t *testing.T) {
	for _, test := range tests {
		li := structure.NewList()
		for _, v := range test.insert {
			li.Insert(v)
		}
		assert.Equal(t, test.result, isDoublyListPalindrome(li))
	}
}
