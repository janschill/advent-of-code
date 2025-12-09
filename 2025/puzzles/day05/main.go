package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/janschill/advent-of-code-2025/puzzles/helpers"
)

func main() {
	input := filepath.Join(puzzleDir(), "input.txt")
	lines := helpers.MustLines(input)
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	var idRanges [][]int
	checking := false
	for _, line := range lines {
		if line == "" {
			checking = true
			continue
		}
		if checking {
			id, _ := strconv.Atoi(line)
			for _, r := range idRanges {
				if id <= r[1] && id >= r[0] {
					sum += 1
					break
				}
			}
		} else {
			idRange := strings.Split(line, "-")
			lo, _ := strconv.Atoi(idRange[0])
			hi, _ := strconv.Atoi(idRange[1])
			idRanges = append(idRanges, []int{lo, hi})
		}

	}
	return sum
}

type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return make(IntSet)
}

func (s IntSet) Add(item int) {
	s[item] = struct{}{}
}

// set with lo and hi
// first iteration set them
// s: [3,5]
// s: [3,5,10,14]
// s: [3,5,10,14,16,20]
// s: [3,5,10,14,14]
func part2(lines []string) int {
	var ranges [][2]int

	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		lo, _ := strconv.Atoi(parts[0])
		hi, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, [2]int{lo, hi})
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})

	var merged [][2]int
	curLo, curHi := ranges[0][0], ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		nextLo := ranges[i][0]
		nextHi := ranges[i][1]
		if nextLo <= curHi+1 {
			curHi = max(curHi, nextHi)
		} else {
			merged = append(merged, [2]int{curLo, curHi})
			curLo, curHi = nextLo, nextHi
		}
	}
	merged = append(merged, [2]int{curLo, curHi})

	sum := 0
	for _, m := range merged {
		sum += m[1] - m[0] + 1
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
