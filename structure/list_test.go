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
		assert.Equal(t, test.size, l.Size)
	}
}

func Test_SingleListInsert(t *testing.T) {
	l := NewSingleList()
	for _, test := range testsInsert {
		l.Insert(test.value)
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.Size)
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
		removeNode func(l *List) *Node
		target     string
		size       int
	}{
		{
			removeNode: func(li *List) *Node { return li.FirstNode },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(li *List) *Node { return nil },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(li *List) *Node { return li.LastNode },
			target:     "bbccdd",
			size:       3,
		},
		{
			removeNode: func(li *List) *Node { return li.FirstNode.NextNode },
			target:     "bbdd",
			size:       2,
		},
	}

	for _, test := range tests {
		l.Remove(test.removeNode(l))
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.Size)
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
		removeNode func(li *SinglyList) *SinglyNode
		target     string
		size       int
	}{
		{
			removeNode: func(li *SinglyList) *SinglyNode { return li.FirstNode },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(li *SinglyList) *SinglyNode { return nil },
			target:     "bbccddee",
			size:       4,
		},
		{
			removeNode: func(li *SinglyList) *SinglyNode { return li.LastNode },
			target:     "bbccdd",
			size:       3,
		},
		{
			removeNode: func(li *SinglyList) *SinglyNode { return li.FirstNode.NextNode },
			target:     "bbdd",
			size:       2,
		},
	}

	for _, test := range tests {
		l.Remove(test.removeNode(l))
		assert.Equal(t, test.target, stringTestList(l))
		assert.Equal(t, test.size, l.Size)
	}
}

// helper
func stringTestList(lr interface{}) string {
	var result string
	switch l := lr.(type) {
	case *List:
		node := l.FirstNode
		for node != nil {
			result += node.Value.(string)
			node = node.NextNode
		}
	case *SinglyList:
		node := l.FirstNode
		for node != nil {
			result += node.Value.(string)
			node = node.NextNode
		}
	}
	return result
}
