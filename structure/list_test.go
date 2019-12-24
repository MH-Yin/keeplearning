package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	assert := assert.New(t)
	lru := NewLruList(0)
	assert.Nil(lru)

	lru = NewLruList(3)
	tests := []struct {
		value  string
		target string
	}{
		{
			value:  "aa",
			target: "aa",
		},
		{
			value:  "bb",
			target: "aabb",
		},
		{
			value:  "cc",
			target: "aabbcc",
		},
		{
			value:  "dd",
			target: "bbccdd",
		},
		{
			value:  "ee",
			target: "ccddee",
		},
	}

	for _, test := range tests {
		lru.Insert(test.value)
		assert.Equal(test.target, stringLru(lru))
	}
}

// helper
func stringLru(lr List) string {
	node := lr.GetFirstNode()
	var result string
	for !node.IsNil() {
		result += node.GetValue().(string)
		node = node.GetNextNode()
	}
	return result
}
