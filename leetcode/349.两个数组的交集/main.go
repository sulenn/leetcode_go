package main

func intersection(nums1 []int, nums2 []int) []int {
	numDict := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		numDict[nums1[i]] = struct{}{}
	}
	result := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := numDict[nums2[i]]; ok {
			result = append(result, nums2[i])
			delete(numDict, nums2[i])
		}
	}
	return result
}
