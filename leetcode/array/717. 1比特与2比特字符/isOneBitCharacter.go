package main

//https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/

//顺序遍历
func isOneBitCharacter(bits []int) bool {
	l := 0
	for ; l <= len(bits)-2; {  // 遍历到最后一个字符
		if bits[l] == 1 {
			l+=2
		} else {
			l++
		}
	}
	if l == len(bits)-1 {return true}  // 如果最后一个字符存在，则为 true
	return false
}
