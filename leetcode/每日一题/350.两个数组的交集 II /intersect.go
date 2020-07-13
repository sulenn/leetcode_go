package main

func intersect(nums1 []int, nums2 []int) []int {
	dic := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		dic[nums1[i]]++
	}
	result := make([]int, 0)
	for j := 0; j < len(nums2); j++ {
		if v, ok := dic[nums2[j]]; ok && v > 0 {
			dic[nums2[j]]--
			result = append(result, nums2[j])
		}
	}
	return result
}
