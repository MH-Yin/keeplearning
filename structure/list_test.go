package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsInsert = []struct {
	value  string
	target string
	size   int
}{
	{
		value:  "aa",
		target: "aa",
		size:   1,
	},
	{
		value:  "bb",
		target: "aabb",
		size:   2,
	},
	{
		value:  "cc",
		target: "aabbcc",
		size:   3,
	},
	{
		value:  "dd",
		target: "aabbccdd",
		size:   4,
	},
	{
		value:  "ee",
		target: "aabbccddee",
		size:   5,
	},
}

func Test_ListInsert(t *testing.T) {
	l := NewList()
	for _, test := range testsInsert {
		l.Insert(test.value)
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.GetSize())
	}
}

func Test_SingleListInsert(t *testing.T) {
	l := NewSingleList()
	for _, test := range testsInsert {
		l.Insert(test.value)
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.GetSize())
	}
}

func Test_ListRemove(t *testing.T) {
	l := NewList()
	l.Insert("aa")
	l.Insert("bb")
	l.Insert("cc")
	l.Insert("dd")
	l.Insert("ee")

	tests := []struct {
		removeNode func(l List) Node
		target     string
		size       int
	}{
		{
			removeNode: func(l List) Node { return l.GetFirstNode() },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(l List) Node { return nil },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(l List) Node { return l.GetLastNode() },
			target:     "bbccdd",
			size:       3,
		},
		{
			removeNode: func(l List) Node { return l.GetFirstNode().GetNextNode() },
			target:     "bbdd",
			size:       2,
		},
	}

	for _, test := range tests {
		l.Remove(test.removeNode(l))
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.GetSize())
	}
}

func Test_SingleListRemove(t *testing.T) {
	l := NewSingleList()
	l.Insert("aa")
	l.Insert("bb")
	l.Insert("cc")
	l.Insert("dd")
	l.Insert("ee")

	tests := []struct {
		removeNode func(l SinglyList) SinglyNode
		target     string
		size       int
	}{
		{
			removeNode: func(l SinglyList) SinglyNode { return l.GetFirstNode() },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(l SinglyList) SinglyNode { return nil },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(l SinglyList) SinglyNode { return l.GetLastNode() },
			target:     "bbccdd",
			size:       3,
		},
		{
			removeNode: func(l SinglyList) SinglyNode { return l.GetFirstNode().GetNextNode() },
			target:     "bbdd",
			size:       2,
		},
	}

	for _, test := range tests {
		l.Remove(test.removeNode(l))
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.GetSize())
	}
}

// helper
func stringTestList(lr interface{}) string {
	var result string
	switch l := lr.(type) {
	case List:
		node := l.GetFirstNode()
		for !node.IsNil() {
			result += node.GetValue().(string)
			node = node.GetNextNode()
		}
	case SinglyList:
		node := l.GetFirstNode()
		for !node.IsNil() {
			result += node.GetValue().(string)
			node = node.GetNextNode()
		}
	}
	return result
}
