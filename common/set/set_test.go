package set_test

import (
	"testing"

	"github.com/henningstorck/advent-of-code/common/geometry"
	"github.com/henningstorck/advent-of-code/common/set"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := set.Set[string]{}
	data = data.Add("a")
	data = data.Add("a")
	data = data.Add("b")
	assert.Equal(t, []string{"a", "b"}, data.Values())
}

func TestRemove(t *testing.T) {
	data := set.Set[string]{}
	data = data.Add("a")
	data = data.Add("a")
	data = data.Add("b")
	data = data.Remove("a")
	assert.Equal(t, []string{"b"}, data.Values())
}

func TestContains(t *testing.T) {
	data := set.Set[string]{}
	data = data.Add("a")
	assert.True(t, data.Contains("a"))
	assert.False(t, data.Contains("b"))
}

func TestContainsComplexType(t *testing.T) {
	data := set.Set[geometry.Point2D]{}
	data = data.Add(geometry.NewPoint2D(2, 3))
	assert.True(t, data.Contains(geometry.NewPoint2D(2, 3)))
	assert.False(t, data.Contains(geometry.NewPoint2D(2, 4)))
}
