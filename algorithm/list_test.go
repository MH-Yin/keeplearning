package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

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

func Test_checkRingInSinglyList(t *testing.T) {
	tests := []struct {
		singlyList *structure.SinglyList
		result     bool
	}{
		{
			singlyList: initSingleNodesWithRing(false,20),
			result:     false,
		},
		{
			singlyList: initSingleNodesWithRing(true,20),
			result:     true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, checkRingInSinglyList(test.singlyList))
	}
}

func initSingleNodes(i int) (singleNodes []*structure.SinglyNode) {
	for ; i > 0; i-- {
		singleNodes = append(singleNodes, &structure.SinglyNode{Value: i})
	}
	return
}

func initSingleNodesWithRing(ring bool, length int) *structure.SinglyList {
	nodes := initSingleNodes(length)
	l := structure.NewSingleList()
	l.FirstNode = nodes[0]
	nodeK := l.FirstNode
	var mark = -1
	if ring {
		rand.Seed(time.Now().UnixNano())
		mark = rand.Intn(length)
	}
	for i, node := range nodes {
		if i == 0 {
			continue
		}
		nodeK.NextNode = node
		nodeK = node
		if i == mark {
			node.NextNode = nodes[0]
			break
		}
	}
	return l
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
