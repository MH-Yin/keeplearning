package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LRU(t *testing.T) {
	lru := newLruList(0)
	assert.Nil(t, lru)

	var cap = 3
	lru = newLruList(cap)
	tests := []struct {
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
			target: "bbccdd",
			size:   3,
		},
		{
			value:  "ee",
			target: "ccddee",
			size:   3,
		},
	}

	for _, test := range tests {
		lru.Insert(test.value)
		assert.Equal(t, test.target, stringTestList(lru.l))
		assert.Equal(t, test.size, lru.GetSize())
	}
}
