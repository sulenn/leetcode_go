package main

import "fmt"

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n == 1 {
			fmt.Println(2)
			continue
		}
		if n == 2 {
			fmt.Println(3)
			continue
		}
		if n == 3 {
			fmt.Println(5)
			continue
		}
		fmt.Println(calNum(n))
	}
}

func calNum(n int) int {
	nums1 := []int{2, 3, 5}
	nums2 := make([]int, 0)
	nums := []int{2, 3, 5}
	var num int
	i := 3
	for i < n {
		nums2 = make([]int, 0)
		for j := 0; j < len(nums1); j++ {
			for v := 0; v < 3; v++ {
				num = nums1[j]*10 + nums[v]
				if i == n-1 {
					return num
				}
				nums2 = append(nums2, num)
				i++
			}
		}
		nums1, nums2 = nums2, nums1
	}
	return -1
}
