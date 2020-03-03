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
	size  int
}

func NewArrayStack(cap int) Stack {
	return &arrayStack{
		items: make([]int, cap),
		size:  cap,
	}
}

func (s *arrayStack) Push(item int) bool {
	if s.size == s.count {
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
