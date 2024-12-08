package main

import (
    "testing"
)

func TestCalculateTotalDistance(t *testing.T) {
    leftList := []int{3, 4, 2, 1, 3, 3}
    rightList := []int{4, 3, 5, 3, 9, 3}

    totalDistance := calculateTotalDistance(leftList, rightList)
    expectedDistance := 11

    if totalDistance != expectedDistance {
        t.Errorf("expected %d, got %d", expectedDistance, totalDistance)
    }
}
