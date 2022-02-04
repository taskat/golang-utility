package stack

//Stack implements a stack structure
type Stack struct {
	data []interface{}
}

//Creates a new Stack object
func Create() Stack {
	return Stack{data: make([]interface{}, 0)}
}

//Len returns the length (depth) of the stack
func (s Stack) Len() int {
	return len(s.data)
}

//Pop returns and removes the last element of the stack
func (s *Stack) Pop() interface{} {
	item := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return item
}

//Push insert the item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

//Top returns the last element of the stack, but does not remove it
func (s Stack) Top() interface{} {
	return s.data[len(s.data) - 1]
}
