package main

import (
	"fmt"
	"io"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {break}
		arr := make([]int, 2*n)
		for i:=0; i<n; i++ {
			fmt.Scan(&arr[i*2], &arr[i*2+1])
		}
		fmt.Println(solution2(arr))
	}
}

func solution2(nums []int) int {
	arr:=make([]int,len(nums)/2)
	for i:=0; i<len(arr); i++{
		arr[i]=i
	}
	for i:=0;i<len(nums);i+=2{
		temp1 := nums[i]/2
		temp2 := nums[i+1]/2
		join(arr,temp1,temp2)
	}
	count:=0
	for i:=0;i<len(arr);i++{
		if searchElement(arr,i)==i {
			count++
		}
	}
	return len(arr)-count

}
func searchElement(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}
func join(arr []int,i,j int)  {
	element1 := searchElement(arr,i)
	element2 := searchElement(arr,j)
	if element1 != element2{
		arr[element1] = element2
		arr[j] = element2
	}
}
