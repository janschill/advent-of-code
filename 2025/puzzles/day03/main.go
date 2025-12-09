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

func part1(lines []string) int {
	for _, line := range lines {
		log.Default().Println(line)
	}
	return 0
}

func part2(lines []string) int {
	for _, line := range lines {
		log.Default().Println(line)
	}
	return 0
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
