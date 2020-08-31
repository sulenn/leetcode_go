package main

import (
	"container/list"
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	length := len(rooms)
	flag := make([]bool, length)
	li := list.New()
	for i := 0; i < len(rooms[0]); i++ {
		li.PushBack(rooms[0][i])
	}
	flag[0] = true
	for li.Len() != 0 {
		room := li.Front()
		value := room.Value.(int)
		li.Remove(room)
		if value >= length {
			return false
		}
		if flag[value] == false {
			flag[value] = true
			for j := 0; j < len(rooms[value]); j++ {
				li.PushBack(rooms[value][j])
			}
		}
	}
	for i := 0; i < len(flag); i++ {
		if !flag[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
