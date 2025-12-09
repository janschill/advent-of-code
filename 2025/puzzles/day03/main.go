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

// find largest digit with index N-1
// keep index and value
// use first digit
// find second digits that is largest to the right of first (loop from i + 1)
func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		maxFirst, idxFirst := -1, -1
		maxSecond := -1
		for i := 0; i < len(line)-1; i++ {
			d, _ := strconv.Atoi(string(line[i]))
			if d > maxFirst {
				maxFirst = d
				idxFirst = i
			}
		}
		for i := idxFirst + 1; i < len(line); i++ {
			d, _ := strconv.Atoi(string(line[i]))
			if d > maxSecond {
				maxSecond = d
			}
		}

		// fmt.Println(maxFirst)
		// fmt.Println(maxSecond)
		joltage, _ := strconv.Atoi(strconv.Itoa(maxFirst) + strconv.Itoa(maxSecond))
		// fmt.Println(joltage)
		// fmt.Println("---")
		sum += joltage
	}

	return sum
}

func findMax(s string, sIdx, eIdx int) (m, idx int) {
	for i := sIdx; i < eIdx; i++ {
		d, _ := strconv.Atoi(string(s[i]))
		if d > m {
			m = d
			idx = i
		}
	}

	return
}

// find largest digit in position 0 to N-12
// find largest digit in position prevIdx to N-11
func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		joltage := ""
		v, idx := 0, -1

		for i := 12; i > 0; i-- {
			v, idx = findMax(line, idx+1, len(line)-i+1)
			joltage += strconv.Itoa(v)
		}

		j, _ := strconv.Atoi(joltage)
		sum += j
	}

	return sum
}

// 811111111111
// 811111111119
// ---
// 423423423427
// 434234234278
// ---
// 888191111211
// 888911112111

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
