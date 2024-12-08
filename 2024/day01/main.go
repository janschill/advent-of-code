package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// os.Open(filename): Opens the file for reading.
// defer file.Close(): Ensures the file is closed when the function exits, even if an error occurs.
func readInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
		return nil, nil, err
	}
	defer file.Close()

	// bufio.NewScanner(file): Creates a new scanner for the file.
	// scanner.Scan(): Advances the scanner to the next line.
	// scanner.Text(): Returns the current line as a string.
	scanner := bufio.NewScanner(file)
	var leftList, rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		// strings.Fields(line): Splits the line into fields (words) based on whitespace.
		numbers := strings.Fields(line)
		// strconv.Atoi(s): Converts the string s to an integer.
		leftNum, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("failed to read leftnum %s", err)
		}
		rightNum, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("failed to read rightnum %s", err)
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}

func calculateSimilarityScore(leftList, rightList []int) int {
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	return similarityScore
}

func calculateTotalDistance(leftList, rightList []int) int {
	// sort.Ints(slice): Sorts the slice of integers in ascending order.
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	return totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	leftList, rightList, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("failed to read input %s", err)
	}

	totalDistance := calculateTotalDistance(leftList, rightList)
	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
