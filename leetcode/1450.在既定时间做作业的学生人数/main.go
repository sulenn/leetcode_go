package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(busyStudent([]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7))

}
