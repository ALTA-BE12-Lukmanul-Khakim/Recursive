package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 0, Fibonacci(0))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 144, Fibonacci(12))
	})
}
