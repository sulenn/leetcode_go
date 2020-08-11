package main

import (
	"errors"
	"fmt"
)

//https://leetcode-cn.com/problems/validate-stack-sequences/

type Stack struct {
	Capacity int
	Index    int
	Elements []int
}

func NewStack(capacity int) *Stack {
	return &Stack{Capacity: capacity, Index: -1, Elements: make([]int, capacity, capacity)}
}

func (s *Stack) Push(num int) error {
	if s.Index >= s.Capacity-1 {
		return errors.New("stack has filled")
	}
	s.Index++
	s.Elements[s.Index] = num
	return nil
}

func (s *Stack) Pop() (error, int) {
	if s.Index < 0 {
		return errors.New("stack is empty"), -1
	}
	num := s.Elements[s.Index]
	s.Index--
	return nil, num
}

func (s *Stack) Head() (error, int) {
	if s.Index < 0 {
		return errors.New("stack is empty"), -1
	}
	return nil, s.Elements[s.Index]
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := NewStack(len(pushed))
	i, j := 0, 0
	for i < len(pushed) && j < len(pushed) {
		stack.Push(pushed[i])
		for {
			err, num := stack.Head()
			if err != nil {
				break
			}
			if num == popped[j] {
				_, _ = stack.Pop()
				j++
			} else {
				break
			}
		}
		i++
	}
	return stack.Index == -1
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
