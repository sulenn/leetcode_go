package minSteps

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-175/problems/minimum-number-of-steps-to-make-two-strings-anagram/

//借助 hashmap 结题。先保存 s 字符串中所有的字符和数量。然后遍历字符串 t，如果有陌生的或者相同但是数量累计超过 s 字符串的 result++
func minSteps(s string, t string) int {
	charMap := make(map[byte]int)
	length := len(s)
	for i:=0; i<length; i++ {    // 保存 s 字符串中所有的字符和数量
		if _,ok := charMap[s[i]]; ok {
			charMap[s[i]]++
		} else {
			charMap[s[i]] = 1
		}
	}
	result := 0
	for i:=0; i<length; i++ {  // 遍历字符串 t，如果有陌生的或者相同但是数量累计超过 s 字符串的 result++
		if count ,ok := charMap[t[i]]; ok && count>0 {
			charMap[t[i]]--
		} else {
				result++
		}
	}
	return result
}

func main() {
	fmt.Println(minSteps("bab", "aba"))
	fmt.Println(minSteps("leetcode", "practice"))
	fmt.Println(minSteps("anagram", "mangaar"))
}
