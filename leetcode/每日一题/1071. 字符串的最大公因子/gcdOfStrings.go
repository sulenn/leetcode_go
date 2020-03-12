package main

//https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/

func gcdOfStrings(str1 string, str2 string) string {
	commonStr := ""
	j := 0
	for ;j < len(str1) && j <len(str2);j++ {
		if str1[j] == str2[j] {
			commonStr += string(str1[j])
		} else {
			break
		}
	}
	flag := false
	for ;len(commonStr) > 0;commonStr=commonStr[:len(commonStr)-1] {
		flag = false
		if len(str1) % len(commonStr) != 0 {continue} // 匹配 str1
		for i:=0; len(str1)-i >= len(commonStr); i += len(commonStr) {
			if str1[i:i+len(commonStr)] != commonStr {flag = true;break}
		}
		if flag {continue}
		if len(str2) % len(commonStr) != 0 {continue} // 匹配 str2
		for i:=0; len(str2)-i >= len(commonStr); i += len(commonStr) {
			if str2[i:i+len(commonStr)] != commonStr {flag = true;break}
		}
		if flag {continue}
		break
	}
	return commonStr
}
