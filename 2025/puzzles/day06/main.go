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

func doMath(a int, operator rune, b int) int {
	switch operator {
	case '+':
		return a + b
	case '*':
		return a * b
	}
	return 0
}

func part1(lines []string) (sum int) {
	grid := helpers.ToGrid2(lines)
	rows, cols := len(grid), len(grid[0])
	for c := 0; c < cols; c++ {
		operator := grid[rows-1][c]
		switch operator {
		case "+":
			s := 0
			for r := 0; r < rows-1; r++ {
				n, _ := strconv.Atoi(grid[r][c])
				s += n
			}
			sum += s
		case "*":
			s := 1
			for r := 0; r < rows-1; r++ {
				n, _ := strconv.Atoi(string(grid[r][c]))
				s *= n
			}
			sum += s
		}
	}

	return
}

func part2(lines []string) int {
	// for _, line := range lines {
	// 	log.Default().Println(line)
	// }
	return 0
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
