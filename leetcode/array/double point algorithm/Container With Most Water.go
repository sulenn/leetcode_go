package main

import "fmt"

func maxArea(height []int) int {
	frontPoint := 0
	backPoint := len(height) - 1
	maxArea := 0
	for frontPoint < backPoint {
		if height[frontPoint] > height[backPoint] {
			tempArea := height[backPoint] * (backPoint - frontPoint)
			if tempArea > maxArea {
				maxArea = tempArea
			}
			backPoint--
		} else {
			tempArea := height[frontPoint] * (backPoint - frontPoint)
			if tempArea > maxArea {
				maxArea = tempArea
			}
			frontPoint++
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
