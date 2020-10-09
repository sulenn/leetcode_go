package main

//https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/

//自己的解法，时间和空间复杂度为 O（n^2）
//前缀和+状态压缩，时间复杂度为O(n),空间复杂度为常数：https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/solution/c-qian-zhui-he-zhuang-tai-ya-suo-xiang-jie-by-yizh/
func findTheLongestSubstring(s string) int {
	arr := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = make([]int, 6) // 第一位为字母总数，第二到第六位为  'a'，'e'，'u'，'i'，'o' 各个字母的数量
	}
	max := 0
	for i := 0; i < len(s); i++ {
		if i > 0 { // 复制上一个状态
			arr[i][0] = arr[i-1][0]
			arr[i][1] = arr[i-1][1]
			arr[i][2] = arr[i-1][2]
			arr[i][3] = arr[i-1][3]
			arr[i][4] = arr[i-1][4]
			arr[i][5] = arr[i-1][5]
		}
		arr[i][0]++   // 总数+1
		switch s[i] { // 新字符
		case 'a':
			arr[i][1]++
		case 'e':
			arr[i][2]++
		case 'u':
			arr[i][3]++
		case 'i':
			arr[i][4]++
		case 'o':
			arr[i][5]++
		}
		if judge(arr[i]) { // 当前字符串满足条件
			max = arr[i][0]
			continue
		}
		num := getMaxLen(arr[:i], arr[i], max) // 当前字符串和前面的字符串匹配
		if num > max {
			max = num
		}
	}
	return max
}

func judge(arr []int) bool { // 判断当前字符串是否满足条件
	if arr[1]%2 == 0 && arr[2]%2 == 0 && arr[3]%2 == 0 && arr[4]%2 == 0 && arr[5]%2 == 0 {
		return true
	}
	return false
}

func getMaxLen(arr [][]int, latest []int, max int) int { // 获得当前字符串和前面字符串的差值
	for i := 0; i < len(arr); i++ {
		if latest[0]-arr[i][0] <= max {
			return 0
		}
		if (latest[1]-arr[i][1])%2 != 0 {
			continue
		}
		if (latest[2]-arr[i][2])%2 != 0 {
			continue
		}
		if (latest[3]-arr[i][3])%2 != 0 {
			continue
		}
		if (latest[4]-arr[i][4])%2 != 0 {
			continue
		}
		if (latest[5]-arr[i][5])%2 != 0 {
			continue
		}
		return latest[0] - arr[i][0]
	}
	return 0
}
