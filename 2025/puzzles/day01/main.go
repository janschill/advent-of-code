package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/janschill/advent-of-code-2025/puzzles/helpers"
)

func main() {
	input := filepath.Join(puzzleDir(), "input.txt")
	lines := helpers.MustLines(input)
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

// Read lines
// Start at 50
// Subtract when starts with L
// Add when starts with R
// Only operate on base 100
// Count occurences of 0 after operation
func part1(lines []string) int {
	sum := 0
	pos := 50
	for _, line := range lines {
		dir := line[0]
		delta, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == 'L' {
			delta = -delta
		}

		pos = (pos + delta) % 100
		if pos < 0 {
			pos += 100
		}
		if pos == 0 {
			sum += 1
		}
	}

	return sum
}

func part2(lines []string) int {
	zeroHits := 0
	pos := 50
	absPos := 50

	for _, line := range lines {
		dir := line[0]
		delta, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir == 'L' {
			delta = -delta
		}

		newAbs := absPos + delta
		zeroHits += crossings(absPos, newAbs)

		pos = (pos + delta) % 100
		if pos < 0 {
			pos += 100
		}
		absPos = newAbs
	}

	return zeroHits
}

func crossings(oldPos, newPos int) int {
	switch {
	case newPos > oldPos:
		// multiples in (oldPos, newPos]
		return floorDiv(newPos, 100) - floorDiv(oldPos, 100)
	case newPos < oldPos:
		// multiples in [newPos, oldPos)
		return floorDiv(oldPos-1, 100) - floorDiv(newPos-1, 100)
	default:
		return 0
	}
}

func floorDiv(x, y int) int {
	if x >= 0 {
		return x / y
	}
	return -(((-x) + y - 1) / y)
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
