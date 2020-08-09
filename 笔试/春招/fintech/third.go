package main

import (
	"fmt"
	"io"
	"math"
)

func main() {
	var n int
	var k int
	for {
		_, err := fmt.Scan(&n, &k)
		if err == io.EOF {
			break
		}
		arr := make([]int, 10001)
		var num int
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			arr[num]++
		}
		result := math.MaxInt64
		for i := 1; i < 10001; i++ {
			if arr[i] != 0 {
				value := way1(i, arr, k)
				if value < result {
					result = value
				}
			}
		}
		fmt.Println(result)
	}
}

func way1(index int, arr []int, k int) int {
	left := left(index, arr, k)
	right := right(index, arr, k)
	if left < right {
		return left
	}
	return right
}

func left(index int, arr []int, k int) int {
	count := arr[index]
	result := 0
	for i := 1; i < index; i++ {
		if arr[i] != 0 {
			if arr[i] >= k-count {
				result += (k - count) * (index - i)
				count = k
				break
			} else {
				result += arr[i] * (index - i)
				count += arr[i]
			}
		}
	}
	if count != k {
		for i := 10000; i > index; i-- {
			if arr[i] != 0 {
				if arr[i] >= k-count {
					result += (k - count) * (i - index)
					count = k
					break
				} else {
					result += arr[i] * (i - index)
					count += arr[i]
				}
			}
		}
	}
	return result
}

func right(index int, arr []int, k int) int {
	count := arr[index]
	result := 0
	for i := 10000; i > index; i-- {
		if arr[i] != 0 {
			if arr[i] >= k-count {
				result += (k - count) * (i - index)
				count = k
				break
			} else {
				result += arr[i] * (i - index)
				count += arr[i]
			}
		}
	}
	if count != k {
		for i := 1; i < index; i++ {
			if arr[i] != 0 {
				if arr[i] >= k-count {
					result += (k - count) * (index - i)
					count = k
					break
				} else {
					result += arr[i] * (index - i)
					count += arr[i]
				}
			}
		}
	}
	return result
}
