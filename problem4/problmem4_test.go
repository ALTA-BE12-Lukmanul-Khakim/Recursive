package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSquence(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 6, MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 7, MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	})
}
