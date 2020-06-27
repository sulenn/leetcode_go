package main

//https://leetcode-cn.com/problems/pattern-matching-lcci/

//抄袭：https://leetcode-cn.com/problems/pattern-matching-lcci/solution/shuang-bai-golangti-jie-dai-ying-wen-zhu-shi-by-uk/
func patternMatching(pattern string, value string) bool {
	countA, countB := 0, 0
	for _, v := range pattern {
		if v == 'a' {
			countA++
		}
		// we don't do countB++ here to save the INC cycles when len(pattern) is quite long
	}
	countB = len(pattern) - countA

	if value == "" {
		// either no 'a' or no 'b' can lead to this result
		return countA == 0 || countB == 0
	}
	if pattern == "" {
		// value has to be "" otherwise it violates previous if statement
		return value == ""
	}

	lenV := len(value)

	// given: len(value) > 0 AND
	//   len(value) == countA*lenA + countB*lenB, find lenA and lenB in type of integer
	if countA == 0 {
		// countB can't be 0
		if lenV%countB == 0 {
			return match(value, pattern, 0, lenV/countB)
		} else {
			return false
		}
	}
	if countB == 0 {
		// countA can't be 0
		if lenV%countA == 0 {
			return match(value, pattern, lenV/countA, 0)
		} else {
			return false
		}
	}
	hopeB := 0
	// countA and countB both are non-zero, find lenA and lenB
	for lenA := 0; lenA*countA <= lenV; lenA++ {
		hopeB = lenV - lenA*countA
		if hopeB%countB == 0 {
			if match(value, pattern, lenA, hopeB/countB) {
				return true
			}
			// not answer, continue
		}
		// not answer, continue
	}
	return false
}

func match(value, pattern string, lenA, lenB int) bool {
	str, strA, strB := "", "", ""
	p := 0
	// iterate every pattern character to validate the corresponding
	// string slice from "value" matches the first appeared patterned string of this pattern.
	// be careful lenA or lenB can be zero here.
	for _, c := range pattern {
		if c == 'a' {
			if lenA == 0 {
				continue
			}
			str = value[p : p+lenA]
			if strA == "" {
				strA = str
			}
			if strA != str {
				return false
			}
			p += lenA
			continue
		}
		// else 'b'
		if lenB == 0 {
			continue
		}
		str = value[p : p+lenB]
		if strB == "" {
			strB = str
		}
		if strB != str {
			return false
		}
		p += lenB
	}
	return true
}
