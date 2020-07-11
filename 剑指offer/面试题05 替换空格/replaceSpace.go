package main

//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

//正规剑指 offer 解法，参考：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/solution/kong-ge-ti-huan-de-liang-chong-jie-fa-by-zouwx2cs/
//
//解题思路
//先扫一趟统计有多少个空格，然后多开辟 空格数 * 2 的存储空间。两个指针，一个指针把原字符串从后往前遍历，挨个往后面拷贝，碰到空格就拷贝"%20"
//
//时间复杂度
//O(n)， 设字符串长度、需要替换的次数分别为m、n
//
//空间复杂度
//O(1)

//这里的简单解法：从后往前，碰见空格就用 %20 替换
func replaceSpace(s string) string {
	newStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			newStr = string(s[i]) + newStr
		} else {
			newStr = "%20" + newStr
		}
	}
	return newStr
}
