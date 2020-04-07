package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	caps := 3
	tests := []struct {
		item   int
		target []int
	}{
		{
			item:   1,
			target: []int{1, 0, 0},
		},
		{
			item:   2,
			target: []int{1, 2, 0},
		},
		{
			item:   3,
			target: []int{1, 2, 3},
		},
		{
			item:   4,
			target: []int{1, 2, 3},
		},
	}
	q := NewQueue(caps)
	for _, test := range tests {
		if err := q.Enqueue(test.item); err != nil {
			assert.Equal(t, errFullQueue, err)
		} else {
			assert.Equal(t, test.target, q.(*queue).items)
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue(3)
	for i := 1; i < 4; i++ {
		q.Enqueue(i)
	}

	for i := 1; i < 5; i++ {
		item, err := q.Dequeue()
		if err != nil {
			assert.Equal(t, errEmptyQueue, err)
		} else {
			assert.Equal(t, i, item)
		}
	}
}
