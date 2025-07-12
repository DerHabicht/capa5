package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var s Stack[string]

	s.Push("a")
	assert.Equal(t, "a", s[0]) // Make sure the item is added
	assert.Equal(t, "a", s.Peek()) // Test the Peek() method
	assert.Equal(t, "a", s[0]) // Make sure Peek() doesn't pop

	s.Push("b")
	assert.Equal(t, "b", s[1]) // Make sure the item is added to the end of the slice

	s.Push("c")
	assert.Equal(t, "c", s.Peek()) // Peek() should be good by now, so we use it to verify push

	v := s.Pop()
	assert.Equal(t, "c", v)
	assert.Equal(t, "b", s.Peek())
	
	s.Push("d")
	s.Push("e")

	l := s.List()
	assert.Equal(t, []string{"e", "d", "b", "a"}, l)
}
