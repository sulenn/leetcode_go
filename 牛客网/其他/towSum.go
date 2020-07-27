package main

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	numDic := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		if index, ok := numDic[numbers[i]]; ok {
			return []int{index, i + 1}
		}
		numDic[target-numbers[i]] = i + 1
	}
	return []int{}
}
