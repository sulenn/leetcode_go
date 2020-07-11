package main

//https://leetcode-cn.com/contest/weekly-contest-169/problems/find-n-unique-integers-sum-up-to-zero/

//![image.png](https://ws1.sinaimg.cn/large/006alGmrgy1gakreos0ctj30wy0gljsd.jpg)

func sumZero(n int) []int {
	result := []int{}
	if n%2 == 0 { // 偶数
		for i := n / 2; i > 0; i-- {
			result = append(result, -i, i)
		}
	} else { //  奇数
		for i := n / 2; i > 0; i-- {
			result = append(result, -i, i)
		}
		result = append(result, 0)
	}
	return result
}

func main() {

}
