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
	h := 0
	t := length-1
	mid := h + (t-h)/2
	midValue := mountainArr.get(mid)
	if midValue == target {
		return mid
	} else if midValue < target {
		if mid-1 >= 0 &&
	}
}

func findMountainLocation(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	h := 0
	t := length-1
	for h <= t {
		mid := h + (t-h)/2
		midValue := mountainArr.get(mid)
		if midValue == target {return mid}
		if
	}
}
