package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var money int
	for {
		_, err := fmt.Scan(&money)
		if err == io.EOF {
			break
		}
		var nums [6]int
		values := [6]int{100, 50, 20, 10, 5, 1}
		for i := 0; i < 6; i++ {
			nums[i] = money / values[i]
			money %= values[i]
		}
		fmt.Println(strconv.Itoa(nums[0]) + " " + strconv.Itoa(nums[1]) + " " + strconv.Itoa(nums[2]) + " " + strconv.Itoa(nums[3]) + " " + strconv.Itoa(nums[4]) + " " + strconv.Itoa(nums[5]))
	}
}
