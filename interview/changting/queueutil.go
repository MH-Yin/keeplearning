package changting

// helper
// circular array implementation queue
type queue struct {
	items []int
	head  int
	tail  int
	size  int
}

func newQueue(cap int, link []int) *queue {
	q := &queue{
		items: make([]int, cap),
		head:  0,
		tail:  0,
		size:  0,
	}

	if link != nil {
		for _, v := range link {
			q.Enqueue(v)
		}
	}
	return q
}

func (q *queue) Enqueue(item int) {
	lens := len(q.items)
	if q.size == lens {
		return
	}
	q.items[q.tail] = item
	q.tail++
	q.size++
	if q.tail == lens {
		q.tail = 0
	}
}

func (q *queue) Dequeue() int {
	if q.size == 0 {
		return 0
	}
	item := q.items[q.head]
	q.head++
	q.size--
	if q.head == len(q.items) {
		q.head = 0
	}
	return item
}
