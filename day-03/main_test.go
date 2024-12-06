package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSum(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	assert.Equal(t, GetSum(input), 161)
}

func TestTraverse(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "mul(1,1)", expected: 1},
		{input: "mul(1,1)---mul(2,2)", expected: 5},
		{input: "mul(5,5)don't()mul(2,2)", expected: 25},
		{input: "mul(5,5)don't()mul(2,2)do()mul(4,4)", expected: 41},
		{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", expected: 48},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, GetSumWithConditions(tc.input))
	}
}
