package main

import ("fmt")

// https://leetcode-cn.com/problems/next-permutation/

// 参考：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func nextPermutation(nums []int)  {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}

	isMax := true
	for l := numsLen - 2; l >= 0; l-- {
		r := l + 1

		if nums[l] < nums[r] {
			isMax = false
			// 找到第一大于nusm[l]的数，并交换
			for i := numsLen - 1; i > l; i-- {
				if nums[i] > nums[l] {
					nums[i], nums[l] = nums[l], nums[i]
					break
				}
			}
			// 倒序排序nums[r:]
			rNums := nums[r:]
			rNumsLen := numsLen - r
			for i := 0; i < rNumsLen/2; i++ {
				j := rNumsLen - i - 1
				rNums[i], rNums[j] = rNums[j], rNums[i]
			}
			break
		}
	}

	if !isMax {
		return
	}
	// 若为最大的则翻转为最小的
	for i := 0; i < numsLen/2; i++ {
		j := numsLen - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	arr := []int {1, 3, 2}
	nextPermutation(arr)
	fmt.Println(arr)
}