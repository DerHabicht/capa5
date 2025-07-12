package collections

// Stack is a simple implementation of the stack data structure in Go
// using generics.
type Stack[T any] []T

// Push(t T) adds an item to the Stack.
func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

// Peek(t T) returns the item on the top of the Stack.
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

// Pop() returns the item on the top of the Stack and then removes it from
// the Stack.
func (s *Stack[T]) Pop() T {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

// List() returns the contents of the Stack as a slice in the order
// that the items would be returned if Pop() were to be successively called.
func (s *Stack[T]) List() []T {
	var t []T
	for i := len(*s) - 1; i >= 0; i-- {
		t = append(t, (*s)[i])
	}

	return t
}
