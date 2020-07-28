package main

/**
 * 反转字符串，输入Douyu，输出uyuoD
 * @param str string字符串 任意字符串
 * @return string字符串
 */

func reverse(str string) string {
	// write code here
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
