package main

//https://leetcode-cn.com/problems/find-the-duplicate-number/

// 双指针（快慢指针）：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/qian-duan-ling-hun-hua-shi-tu-jie-kuai-man-zhi-z-3/
//func findDuplicate(nums []int) int {
//	slow, fast := 0, 0   // 快慢指针
//	slow = nums[slow]
//	fast = nums[nums[fast]]
//	for slow != fast {   // 快慢指针相交点
//		slow = nums[slow]
//		fast = nums[nums[fast]]
//	}
//	head := 0
//	for nums[head] != nums[slow] {   // 二次相交点为重复数字
//		head = nums[head]
//		slow = nums[slow]
//	}
//	return nums[head]
//}

//常规解法，二分法：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/er-fen-fa-si-lu-ji-dai-ma-python-by-liweiwei1419/
func findDuplicate(nums []int) int {
	pre, tail := 1, len(nums)-1
	for pre < tail {
		mid := pre + (tail-pre)/2
		count := 0
		for _, v := range nums {
			if v <= mid {count++}
		}
		// 根据抽屉原理，小于等于 4 的个数如果严格大于 4 个
		// 此时重复元素一定出现在 [1, 4] 区间里
		if count > mid {
			tail = mid  // 重复元素位于区间 [tail, mid]
		} else {
			// if 分析正确了以后，else 搜索的区间就是 if 的反面
			// [mid + 1, right]
			pre = mid + 1
		}
	}
	return pre
}