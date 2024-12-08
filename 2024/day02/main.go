package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func safe(report []int) bool {
	increasing := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		if increasing && (report[i] <= report[i-1] || report[i]-report[i-1] > 3) ||
			!increasing && (report[i] >= report[i-1] || report[i-1]-report[i] > 3) {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	part1Reports := 0
	part2Reports := 0
	for scanner.Scan() {
		var numbers []int
		line := scanner.Text()
		fields := strings.Fields(line)
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			numbers = append(numbers, num)
		}
		if safe(numbers) {
			part1Reports++
			part2Reports++
		} else {
			for i := range numbers {
				r := make([]int, len(numbers))
				copy(r, numbers)
				if safe(append(r[0:i], r[i+1:]...)) {
					part2Reports++
					break
				}
			}
		}
	}

	fmt.Println(part1Reports)
	fmt.Println(part2Reports)
}
