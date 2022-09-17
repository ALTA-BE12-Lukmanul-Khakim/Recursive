package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrismaSegiEmpat(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "17 19 \n23 29 \n31 37 \n156", PrismaSegiEmpat(2, 3, 13))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, "2 3 5 7 11 \n13 17 19 23 29 \n129", PrismaSegiEmpat(5, 2, 1))
	})
}
