package first

func Game24Points(arr []int) bool {
	// write code here
	if len(arr) < 4 {
		return false
	}
	return backtracking(arr, 24)
}

func backtracking(arr []int, target int) bool {
	if len(arr) == 1 {
		if arr[0] == target {
			return true
		} else {
			return false
		}
	}
	num1 := arr[0]
	num2 := arr[1]
	arr = arr[1:]

	arr[0] = num1 + num2
	newArr1 := make([]int, len(arr))
	copy(newArr1, arr)
	if backtracking(newArr1, target) {
		return true
	}

	arr[0] = num1 - num2
	newArr2 := make([]int, len(arr))
	copy(newArr2, arr)
	if backtracking(newArr2, target) {
		return true
	}

	if num2 != 0 {
		arr[0] = num1 / num2
		newArr3 := make([]int, len(arr))
		copy(newArr3, arr)
		if backtracking(newArr3, target) {
			return true
		}
	}

	arr[0] = num1 * num2
	newArr4 := make([]int, len(arr))
	copy(newArr4, arr)
	if backtracking(newArr4, target) {
		return true
	}

	return false
}
