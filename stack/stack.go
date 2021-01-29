package stack

type Stack struct {
	capacity int
	top      int
	items    []interface{}
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		top:      -1,
		items:    make([]interface{}, capacity),
	}
}

func (s *Stack) Push(item interface{}) {
	if !isFull(s) {
		s.top++
		s.items[s.top] = item
	}
}

func (s *Stack) Pop() interface{} {
	if isEmpty(s) {
		return nil
	}
	item := s.items[s.top]
	s.top--
	return item
}

func isEmpty(s *Stack) bool {
	return s.top == -1
}

func isFull(s *Stack) bool {
	return s.capacity-1 == s.top
}
