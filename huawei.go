package main

func CalLocation(x,y int) int {
	if x < 0 {
		return -1
	}
	if x != -y {
		return -1
	}
	return (2*x+1)*(2*x+1)
}