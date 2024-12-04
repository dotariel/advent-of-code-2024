package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 11, Diff(a, b))
}

func TestSimilarity(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, 31, Similarity(a, b))
}
