package main

//https://leetcode-cn.com/problems/long-pressed-name/

func isLongPressedName(name string, typed string) bool {
	p1 := 0
	p2 := 0
	for p1 < len(name) && p2 < len(typed) {
		curP1Byte := name[p1]
		numP1 := 1
		p1++
		for p1 < len(name) {
			if name[p1] == curP1Byte {
				numP1++
				p1++
			} else {
				break
			}
		}
		curP2Byte := typed[p2]
		numP2 := 1
		p2++
		for p2 < len(typed) {
			if typed[p2] == curP2Byte {
				numP2++
				p2++
			} else {
				break
			}
		}
		if curP1Byte != curP2Byte || numP1 > numP2 {
			return false
		}
	}
	if p1 < len(name) || p2 < len(typed) {
		return false
	}
	return true
}
