package structure

import (
	"errors"
)

var (
	errEmptyQueue = errors.New("queue is empty")
	errFullQueue  = errors.New("queue is full")
)

// Queue 队列接口
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

// NewQueue 返回队列接口
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
		return errFullQueue
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
		return 0, errEmptyQueue
	}
	item := q.items[q.head]
	q.head++
	q.size--
	if q.head == len(q.items) {
		q.head = 0
	}
	return item, nil
}
