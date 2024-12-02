package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayXX(t *testing.T) {
	assert := assert.New(t)
	input := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := Part1(input)

		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := Part2(input)

		assert.Equal(actual, expected)
	})
}
