package functional_test

import (
	"fmt"
	"testing"

	"github.com/henningstorck/advent-of-code/common/functional"
	"github.com/stretchr/testify/assert"
)

type Agent struct {
	FirstName string
	LastName  string
}

func TestMap(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		fn       func(int) int
	}{
		{
			[]int{},
			[]int{},
			func(value int) int { return value * value },
		},
		{
			[]int{2, 3, 5, 7, 11},
			[]int{4, 9, 25, 49, 121},
			func(value int) int { return value * value },
		},
	}

	for _, test := range tests {
		result := functional.Map(test.input, test.fn)
		assert.Equal(t, test.expected, result)
	}
}

func TestMapDifferentTypes(t *testing.T) {
	input := []string{"hello", "universe"}
	expected := []int{5, 8}
	fn := func(value string) int { return len(value) }
	result := functional.Map(input, fn)
	assert.Equal(t, expected, result)
}

func TestMapComplexTypes(t *testing.T) {
	input := []Agent{
		{FirstName: "Dana", LastName: "Scully"},
		{FirstName: "Fox", LastName: "Mulder"},
	}

	expected := []string{"Dana Scully", "Fox Mulder"}

	fn := func(agent Agent) string {
		return fmt.Sprintf("%s %s", agent.FirstName, agent.LastName)
	}

	result := functional.Map(input, fn)
	assert.Equal(t, expected, result)
}
