package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeX(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 2, PrimeX(1))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 29, PrimeX(10))
	})
}
