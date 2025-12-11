package main

import (
	"testing"

	"github.com/janschill/advent-of-code-2025/puzzles/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		expected := 40
		actual := part1(helpers.MustLines("test-input.txt"))
		assert.Equal(t, expected, actual)
	})
}

func TestPart2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		expected := 0
		actual := part2(helpers.MustLines("test-input.txt"))
		assert.Equal(t, expected, actual)
	})
}
