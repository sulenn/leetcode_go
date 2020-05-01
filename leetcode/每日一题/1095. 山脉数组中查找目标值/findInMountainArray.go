package main

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	mountainIndex := findMountainLocation(mountainArr)    // 先找数组中最大的峰值，该峰值左右两边递减
	left := findTarget1(target, 0, mountainIndex, mountainArr)   // 先找峰值左边是否有target
	if left != -1 {return left}
	right := findTarget2(target, mountainIndex, length-1, mountainArr) // 再找峰值左边是否有target
	if right != -1 {return right}
	return -1
}

func findMountainLocation(mountainArr *MountainArray) int {   //寻找最大峰值
	length := mountainArr.length()
	h := 0
	t := length - 1
	for h < t {
		mid := h + (t-h)/2
		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			t = mid
		} else {
			h = mid+1
		}
	}
	return h
}

func findTarget1(target int, h int, t int, mountainArr *MountainArray) int {   // 从小到大
	for h <=t {
		mid := h + (t-h)/2
		midValue := mountainArr.get(mid)
		if midValue == target {return mid}
		if midValue > target {t = mid-1}
		if midValue < target {h = mid+1}
	}
	return -1
}

func findTarget2(target int, h int, t int, mountainArr *MountainArray) int {   // 从大到小
	for h <=t {
		mid := h + (t-h)/2
		midValue := mountainArr.get(mid)
		if midValue == target {return mid}
		if midValue > target {h = mid+1}
		if midValue < target {t = mid-1}
	}
	return -1
}
