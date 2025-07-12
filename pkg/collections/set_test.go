package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	var s Set[string]

	s.Add("foo")
	assert.True(t, s.Contains("foo"))

	s.Remove("foo")
	assert.False(t, s.Contains("foo"))

	assert.NotPanics(t, func() { s.Remove("bar") })
}

func TestDiff(t *testing.T) {
	var a Set[string]
	var b Set[string]

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add("d")
	a.Add("e")
	a.Add("f")

	b.Add("a")
	b.Add("b")
	b.Add("c")

	c := a.Sub(b)

	assert.False(t, c.Contains("a"))
	assert.False(t, c.Contains("b"))
	assert.False(t, c.Contains("c"))
	assert.True(t, c.Contains("d"))
	assert.True(t, c.Contains("e"))
	assert.True(t, c.Contains("f"))
}

func TestIntersect(t *testing.T) {
	var a Set[string]
	var b Set[string]

	a.Add("a")
	a.Add("b")
	a.Add("c")

	b.Add("c")
	b.Add("d")
	b.Add("e")

	c := a.Intersect(b)

	assert.False(t, c.Contains("a"))
	assert.False(t, c.Contains("b"))
	assert.True(t, c.Contains("c"))
	assert.False(t, c.Contains("d"))
	assert.False(t, c.Contains("e"))
}

func TestUnion(t *testing.T) {
	var a Set[string]
	var b Set[string]

	a.Add("a")
	a.Add("b")
	a.Add("c")

	b.Add("d")
	b.Add("e")
	b.Add("f")

	c := a.Union(b)

	assert.True(t, c.Contains("a"))
	assert.True(t, c.Contains("b"))
	assert.True(t, c.Contains("c"))
	assert.True(t, c.Contains("d"))
	assert.True(t, c.Contains("e"))
	assert.True(t, c.Contains("f"))
}
