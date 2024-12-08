package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	enabled := true

	// Combined pattern to match all types at once
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches in order of appearance
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fullMatch := match[1] // The complete matched pattern

			switch {
			case fullMatch == "do()":
				enabled = true
			case fullMatch == "don't()":
				enabled = false
			case enabled && fullMatch[:3] == "mul":
				// Extract and multiply numbers only if enabled
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				sum += num1 * num2
			}
		}
	}

	fmt.Printf("Total sum: %d\n", sum)
}
