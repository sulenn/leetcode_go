package main

//https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	dict := make(map[byte]struct{})
	for i := 0; i < len(J); i++ {
		dict[J[i]] = struct{}{}
	}
	count := 0
	for i := 0; i < len(S); i++ {
		if _, ok := dict[S[i]]; ok {
			count++
		}
	}
	return count
}
