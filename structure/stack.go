package structure

type Stack interface {
	Pop()
	Push(item int) bool
	Count() int
}

// arrayStack implemented using arrays
type arrayStack struct {
	items []int
	count int
	cap   int
}

func NewArrayStack(cap int) Stack {
	return &arrayStack{
		items: make([]int, cap),
		cap:   cap,
	}
}

func (s *arrayStack) Push(item int) bool {
	if s.cap == s.count {
		return false
	}
	s.items[s.count] = item
	s.count++
	return true
}

func (s *arrayStack) Pop() {
	if s.count == 0 {
		return
	}
	s.count--
}

func (s *arrayStack) Count() int {
	return s.count
}
