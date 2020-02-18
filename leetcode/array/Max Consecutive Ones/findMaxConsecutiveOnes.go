package main

func findMaxConsecutiveOnes(nums []int) int {
	num := 0
	maximum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			num += 1
		} else {
			if maximum < num {
				maximum = num
			}
			num = 0
		}
	}
	if maximum < num {
		maximum = num
	}
	return maximum
}

func main() {
	
}
