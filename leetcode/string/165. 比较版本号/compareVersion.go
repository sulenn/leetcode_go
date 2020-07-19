package main

import "strconv"

//https://leetcode-cn.com/problems/compare-version-numbers/

func compareVersion(version1 string, version2 string) int {
	str1IndexHead, str1IndexTail := 0, 0
	str2Indexhead, str2IndexTail := 0, 0
	for str1IndexHead < len(version1) && str2Indexhead < len(version2) {
		str1IndexTail = str1IndexHead + 1
		str2IndexTail = str2Indexhead + 1
		for str1IndexTail < len(version1) && version1[str1IndexTail] <= '9' && version1[str1IndexTail] >= '0' {
			str1IndexTail++
		}
		for str2IndexTail < len(version2) && version2[str2IndexTail] <= '9' && version2[str2IndexTail] >= '0' {
			str2IndexTail++
		}
		num1, _ := strconv.Atoi(version1[str1IndexHead:str1IndexTail])
		num2, _ := strconv.Atoi(version2[str2Indexhead:str2IndexTail])
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
		str1IndexHead = str1IndexTail + 1 // 跳过点 "."
		str2Indexhead = str2IndexTail + 1
	}
	for str1IndexHead < len(version1) {
		str1IndexTail = str1IndexHead + 1
		for str1IndexTail < len(version1) && version1[str1IndexTail] <= '9' && version1[str1IndexTail] >= '0' {
			str1IndexTail++
		}
		num1, _ := strconv.Atoi(version1[str1IndexHead:str1IndexTail])
		if num1 != 0 {
			return 1
		}
		str1IndexHead = str1IndexTail + 1
	}
	for str2Indexhead < len(version2) {
		str2IndexTail = str2Indexhead + 1
		for str2IndexTail < len(version2) && version2[str2IndexTail] <= '9' && version2[str2IndexTail] >= '0' {
			str2IndexTail++
		}
		num2, _ := strconv.Atoi(version2[str2Indexhead:str2IndexTail])
		if num2 != 0 {
			return -1
		}
		str2Indexhead = str2IndexTail + 1
	}
	return 0
}
