package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		arr1 := make([][2]int, n)
		var x, y int
		for i := 0; i < n; i++ {
			fmt.Scan(&x, &y)
			arr1[i] = [2]int{x, y}
		}
		arr2 := make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&x, &y)
			arr2[i] = [2]int{x, y}
		}
		min := 99999999.0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				num := solution(arr1[i], arr2[j])
				if num < min {
					min = num
				}
			}
		}
		fmt.Printf("%.3f\n", min)
	}
}

func solution(p1 [2]int, p2 [2]int) float64 {
	result := math.Sqrt(float64((p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])))
	return result
}

//2
//4
//0 0
//0 1
//1 0
//1 1
//2 2
//2 3
//3 2
//3 3
//4
//0 0
//0 0
//0 0
//0 0
//0 0
//0 0
//0 0
//0 0
