package stack

// Stack implements a simple stack structure
type Stack[T any] struct {
	data []T
}

// New creates a new Stack object
func New[T any]() Stack[T] {
	return Stack[T]{data: make([]T, 0)}
}

// Len returns the length (depth) of the stack
func (s Stack[T]) Len() int {
	return len(s.data)
}

// Pop returns and removes the last element of the stack
func (s *Stack[T]) Pop() T {
	item := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return item
}

// Push insert the item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Top returns the last element of the stack, but does not remove it
func (s Stack[T]) Top() T {
	return s.data[len(s.data) - 1]
}
