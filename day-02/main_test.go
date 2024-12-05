package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsSafe(t *testing.T) {
	testCases := []struct {
		report []int
		safe   bool
	}{
		{report: []int{7, 6, 4, 2, 1}, safe: true},
		{report: []int{1, 2, 7, 8, 9}, safe: false},
		{report: []int{9, 7, 6, 2, 1}, safe: false},
		{report: []int{1, 3, 2, 4, 5}, safe: false},
		{report: []int{8, 6, 4, 4, 1}, safe: false},
		{report: []int{1, 3, 6, 7, 9}, safe: true},
	}

	for _, tc := range testCases {
		assert.Equal(t, safe(tc.report), tc.safe)
	}
}
