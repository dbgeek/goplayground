package stack

import (
	"errors"
)

type (
	stack []int
)

func (s *stack) pop() (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("Stack is empty")
	}
	lastValue := len(*s) - 1
	value, stack := (*s)[lastValue], (*s)[:lastValue]
	*s = stack
	return value, nil
}

func (s *stack) push(value int) {
	*s = append(*s, value)
}
