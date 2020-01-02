package algorithm

import (
	"fmt"
	"testing"

	"github.com/MH-Yin/keeplearning/structure"
	"github.com/stretchr/testify/assert"
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
		insert: []int{0, 1, 2, 3, 5, 5, 3, 2, 1, 0},
		result: true,
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
		assert.Equal(t, test.result, isSinglyListPalindrome(li), fmt.Sprintf("%v", test))
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

func Test_reverseSiglyList(t *testing.T) {
	tests := []struct {
		insert []int
		result []int
	}{
		{
			insert: []int{1},
			result: []int{1},
		},
		{
			insert: []int{0, 1, 2, 3},
			result: []int{3, 2, 1, 0},
		},
	}

	for _, test := range tests {
		li := structure.NewSingleList()
		for _, v := range test.insert {
			li.Insert(v)
		}

		reverseSiglyList(li)
		assert.Equal(t, test.result, intArrayTestList(li))
	}
}

// helper
func intArrayTestList(lr interface{}) []int {
	var result []int
	switch l := lr.(type) {
	case *structure.SinglyList:
		node := l.FirstNode
		for node != nil {
			result = append(result, node.Value.(int))
			node = node.NextNode
		}
	}
	return result
}
