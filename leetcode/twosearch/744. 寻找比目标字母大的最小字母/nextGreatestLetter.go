package main

//https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/description/

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters)-1; i++ { // 调整序列
		if letters[i+1] < letters[i] {
			letters = append(letters[i+1:], letters[:i+1]...)
		}
	}
	head := 0
	tail := len(letters) - 1
	for head < tail { // 二分
		mid := head + (tail-head)/2
		if letters[mid] <= target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	if letters[head] > target {
		return letters[head]
	}
	return letters[0]
}
