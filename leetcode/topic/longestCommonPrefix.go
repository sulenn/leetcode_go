package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := strs[0]
	max := strs[0]
	//for i := 1; i < len(strs) ;i++ {
	//	if min > strs[i] {
	//		min = strs[i]
	//	}
	//	if max < strs[i] {
	//		max = strs[i]
	//	}
	//}
	for _, elem:= range strs {
		if min > elem {
			min = elem
		}
		if max < elem {
			max = elem
		}
	}
	for j := 0; j < len(min); j++ {
		if min[j] != max[j] {
			return min[0:j]
		}
	}
	return min
}

func main() {
	//strArray := [] string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix([] string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([] string{"dog","racecar","car"}))
}
