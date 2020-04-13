package main

//https://leetcode-cn.com/problems/intersection-lcci/

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {

}

func line(x1, y1, x2, y2 float64) (float64, float64, float64) {
	if x2 - x1 == 0 {
		return 0, 1, -float64(x1)
	}
	if y2 - y1 == 0 {
		return 1, 0, float64(y1)
	}
	k := (y2-y1)/(x2-x1)
	c := y1-k*x1
	return 1, k, c
}

func main() {
	
}
