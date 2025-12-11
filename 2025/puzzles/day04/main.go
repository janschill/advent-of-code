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
	grid := helpers.ToGrid(lines)
	rows, cols := len(grid), len(grid[0])
	sum := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			rollCount := 0
			cur := grid[r][c]
			if cur == '@' {
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr, nc := r+dr, c+dc
						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						neighbor := grid[nr][nc]
						if neighbor == '@' {
							rollCount += 1
						}
						_ = neighbor
					}
				}
				if rollCount < 4 {
					sum += 1
					// fmt.Printf("(%d,%d)", r, c)
					// fmt.Println()
				}
			}
		}
	}
	return sum
}

func part2(lines []string) int {
	grid := helpers.ToGrid(lines)
	rows, cols := len(grid), len(grid[0])
	sum := 0
	var changed bool

	for {
		var changedCells [][]int
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				rollCount := 0
				cur := grid[r][c]
				if cur == '@' {
					for dr := -1; dr <= 1; dr++ {
						for dc := -1; dc <= 1; dc++ {
							if dr == 0 && dc == 0 {
								continue
							}
							nr, nc := r+dr, c+dc
							if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
								continue
							}
							neighbor := grid[nr][nc]
							if neighbor == '@' {
								rollCount += 1
							}
							_ = neighbor
						}
					}
					if rollCount < 4 {
						sum += 1
						changedCells = append(changedCells, []int{r, c})
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
		for i := 0; i < len(changedCells); i++ {
			x, y := changedCells[i][0], changedCells[i][1]
			grid[x][y] = '.'
		}
		changed = false
	}

	return sum
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
