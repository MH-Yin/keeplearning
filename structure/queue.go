package structure

import (
	"errors"
)

var (
	emptyQueue = errors.New("queue is empty")
	fullQueue  = errors.New("queue is full")
)

type Queue interface {
	Enqueue(item int) error
	Dequeue() (int, error)
}

// circular array implementation queue
type queue struct {
	items []int
	head  int
	tail  int
	size  int
}

func NewQueue(cap int) Queue {
	return &queue{
		items: make([]int, cap),
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *queue) Enqueue(item int) error {
	lens := len(q.items)
	if q.size == lens {
		return fullQueue
	}
	q.items[q.tail] = item
	q.tail++
	q.size++
	if q.tail == lens {
		q.tail = 0
	}
	return nil
}

func (q *queue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, emptyQueue
	}
	item := q.items[q.head]
	q.head++
	q.size--
	if q.head == len(q.items) {
		q.head = 0
	}
	return item, nil
}
