package main

import (
	"fmt"
	"io"
	"sort"
)

//https://www.nowcoder.com/question/next?pid=1088888&qid=36846&tid=31500058

func main() {
	for {
		nums := []int{}
		num := 0
		_, err := fmt.Scan(&num)
		if err == io.EOF {
			break
		}
		temp := 0
		for i := 0; i < num; i++ {
			_, err := fmt.Scan(&temp)
			if err == io.EOF {
				break
			}
			nums = append(nums, temp)
		}
		sort.Ints(nums)
		for i := 0; i < len(nums)-1; {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				i++
			}
		}
		for _, v := range nums {
			fmt.Println(v)
		}
	}
}
