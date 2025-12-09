package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
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

// is true when num is made only of some
// sequence of digits repeated twice.
// 11, 22, 99, 1010, 1188511885, 222222
func isInvalid(num string) bool {
	if len(num)%2 != 0 {
		return false
	}
	mid := len(num) / 2
	left := num[:mid]
	right := num[mid:]

	if left == right {
		return true
	}

	return false
}

func part1(lines []string) int {
	sum := 0
	// 11-22
	idRanges := strings.Split(lines[0], ",")
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := 0; i <= (end - start); i++ {
			number := start + i
			if isInvalid(strconv.Itoa(number)) {
				sum += number
			}
		}
	}

	return sum
}

var divisorMap = make(map[int][]int)

func divisors(n int) []int {
	if cached, ok := divisorMap[n]; ok {
		return cached
	}
	var ds []int
	for d := 1; d*d <= n; d++ {
		if n%d != 0 {
			continue
		}
		ds = append(ds, d)
		if d*d != n {
			ds = append(ds, n/d)
		}
	}
	divisorMap[n] = ds

	return ds
}

func splitIntoParts(num string, d int) []string {
	size := len(num) / d
	parts := make([]string, d)
	for i := 0; i < d; i++ {
		start := i * size
		end := start + size
		parts[i] = num[start:end]
	}

	return parts
}

var invalidMap = make(map[int]bool)

func isInvalid2(num int) bool {
	if cached, ok := invalidMap[num]; ok {
		return cached
	}
	sNum := strconv.Itoa(num)
	length := len(sNum)
	for _, d := range divisors(length) {
		if d < 2 || length%d != 0 {
			continue
		}
		parts := splitIntoParts(sNum, d)
		first := parts[0]
		match := true
		for _, p := range parts[1:] {
			if p != first {
				match = false
				break
			}
		}
		if match {
			invalidMap[num] = true
			return true
		}
	}

	invalidMap[num] = false
	return false
}

func part2(lines []string) int {
	sum := 0
	idRanges := strings.Split(lines[0], ",")
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := 0; i <= (end - start); i++ {
			number := start + i
			if isInvalid2(number) {
				sum += number
			}
		}
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
