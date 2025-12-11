package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	dx, dy int
}

func checkXmas(grid []string, row, col, dx, dy, rows, cols int) bool {
	if !isInBounds(row+dx*3, col+dy*3, rows, cols) {
		return false
	}

	return grid[row][col] == 'X' &&
		grid[row+dx][col+dy] == 'M' &&
		grid[row+dx*2][col+dy*2] == 'A' &&
		grid[row+dx*3][col+dy*3] == 'S'

}

func isInBounds(row, col, rows, cols int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func main() {
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	directions := []Direction{
		{0, 1},
		{1, 0},
		{1, 1},
		{-1, 1},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{1, -1},
	}

	count := 0
	rows := len(grid)
	if rows == 0 {
		fmt.Println("Empty input")
		return
	}
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				if checkXmas(grid, row, col, dir.dx, dir.dy, rows, cols) {
					count++
				}
			}
		}
	}

	fmt.Printf("Count of XMAS is %d\n", count)
}
