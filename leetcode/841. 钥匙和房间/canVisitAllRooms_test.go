package main

import (
	"testing"
)

func Test_canVisitAllRooms(t *testing.T) {

	t.Log(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	t.Log(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
