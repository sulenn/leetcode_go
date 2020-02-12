package main

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

//可以用 hash 和 桶思想 两种解法
//参考：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/tong-de-si-xiang-by-liweiwei1419/
func findRepeatNumber(nums []int) int {
	numMap := make(map[int]interface{})
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return num
		} else {
			numMap[num] = nil
		}
	}
	return -1
}

//func main() {
//
//}
