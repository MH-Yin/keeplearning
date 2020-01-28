package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack(3)
	tests := []struct {
		target    []int
		pushItems []int
	}{
		{
			pushItems: []int{1},
			target:    []int{0: 1, 2: 0},
		},
		{
			pushItems: []int{1, 2},
			target:    []int{1, 1, 2},
		},
		{
			pushItems: []int{1},
			target:    []int{1, 1, 2},
		},
	}

	for _, test := range tests {
		for _, item := range test.pushItems {
			stack.Push(item)
		}
		assert.Equal(t, test.target, stack.(*arrayStack).items)
	}
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack(3)
	tests := []struct {
		target   int
		preItems []int
	}{
		{
		},
		{
			preItems: []int{1},
		},
		{
			preItems: []int{1, 2},
			target:   1,
		},
	}

	for _, test := range tests {
		for _, item := range test.preItems {
			stack.Push(item)
		}
		stack.Pop()
		assert.Equal(t, test.target, stack.Count())
	}
}
