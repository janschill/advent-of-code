package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/janschill/advent-of-code-2025/puzzles/helpers"
)

func main() {
	input := filepath.Join(puzzleDir(), "input.txt")
	lines := helpers.MustLines(input)
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

// Beams: notes current beams and their column index
// Look for S, note down column index
// For every row when '^' encoutered, add to splitSum and
// remove current index from beams, and add curIdx-1 and curIdx+1 to it,

type intSet map[int]struct{}

func (s intSet) Add(v int)    { s[v] = struct{}{} }
func (s intSet) Remove(v int) { delete(s, v) }
func (s intSet) Has(v int) bool {
	_, ok := s[v]
	return ok
}

func part1(lines []string) int {
	beamsColIdx := make(intSet)
	amountSplits := 0
	for _, line := range lines {
		for i, c := range line {
			if c == 'S' {
				beamsColIdx.Add(i)
			}
			if c == '^' && beamsColIdx.Has(i) {
				beamsColIdx.Remove(i)
				amountSplits += 1
				beamsColIdx.Add(i - 1)
				beamsColIdx.Add(i + 1)
			}
		}
	}
	return amountSplits
}

func part2(lines []string) (total int) {
	beamsColIdx := make(map[int]int)
	total = 1
	for _, line := range lines {
		for i, c := range line {
			if c == 'S' {
				beamsColIdx[i]++
			}
			if c == '^' && beamsColIdx[i] > 0 {
				count := beamsColIdx[i]
				beamsColIdx[i] = 0
				beamsColIdx[i-1] += count
				beamsColIdx[i+1] += count
				total += count
			}
		}
	}
	return
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
